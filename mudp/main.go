package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	capnp "capnproto.org/go/capnp/v3"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/event"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p-core/record"
)

const (
	srvAddr         = "224.0.0.1:9999"
	maxDatagramSize = 8192
	ns              = "casm/mudp"
)

var (
	mdnsGroupIPv4 = net.IPv4(224, 0, 0, 251)
	mdnsGroupPort = 5353

	mdnsWildcardAddrIPv4 = &net.UDPAddr{
		IP:   net.ParseIP("224.0.0.0"),
		Port: mdnsGroupPort,
	}
	h host.Host
	err error
)

func main() {
	h, err = libp2p.New(context.Background())
	if err != nil {
		panic(err)
	}
	defer h.Close()

	waitReady(h)
	
	go sendResponses()
	readResponses()
}

func waitReady(h host.Host) {
	sub, err := h.EventBus().Subscribe(new(event.EvtLocalAddressesUpdated))
	if err != nil {
		panic(err)
	}
	defer sub.Close()

	<-sub.Out()
}

func sendResponses() {
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create a new Book struct.  Every message must have a root struct.
	mudp, err := NewRootMudpPacket(seg)
	if err != nil {
		panic(err)
	}
	response, err := mudp.NewResponse()
	response.SetNamespace(ns)
	var peerRecord *record.Envelope
	
	if cab, ok := peerstore.GetCertifiedAddrBook(h.Peerstore()); ok {
		peerRecord, err = cab.GetPeerRecord(h.ID()), nil
		if err != nil {
			panic(err)
		}
	}
	
	envelope, err := response.NewEnvelope(1)
	rec, err := peerRecord.Marshal()
	if err != nil {
		panic(err)
	}
	envelope.Set(0, rec)


	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%v:%v", mdnsGroupIPv4.String(), mdnsGroupPort))
	if err != nil {
		panic(err)
	}

	c, err := net.ListenUDP("udp", addr)
	enc, err := mudp.Message().MarshalPacked()
	if err!=nil{
		panic(err)
	}
	var n int
	for {	
		n, err = c.WriteTo(enc, addr)
		println("Write", n, "to", addr.String())
		time.Sleep(1 * time.Second)
	}
}

func readResponses() {
	udpConn, _ := net.ListenUDP("udp4", mdnsWildcardAddrIPv4)
	defer udpConn.Close()
	
	for {
		readResponse(udpConn)
	}
}

func readResponse(con net.PacketConn) {
	b :=make([]byte, 5_000)
	n,addr, err := con.ReadFrom(b)
	println("Read", n, "from", addr.String())
	if err!=nil{
		panic(err)
	}
	
	msg, err := capnp.UnmarshalPacked(b[:n])
	if err != nil {
		return
	}
	
	mudp, err := ReadRootMudpPacket(msg)
	if err != nil {
		return
	}
	switch mudp.Which() {
	case MudpPacket_Which_request:
		// it's a request packet, do nothing
	case MudpPacket_Which_response:
		// it's a response packet
		response, err := mudp.Response()
		ns, err := response.Namespace()
		if err != nil {
			panic(err)
		}
		log.Println("namespace:", ns)

		// get peer record
		var envelope capnp.DataList
		if envelope, err = response.Envelope(); err != nil {
			return
		}

		var invalidEnvelope *record.Envelope
		var peerRecord peer.PeerRecord

		rawEnvelope, err := envelope.At(0)
		if err!=nil{
			panic(err)
		}
		if invalidEnvelope, err = record.ConsumeTypedEnvelope(rawEnvelope, &peerRecord); err != nil {
			panic(fmt.Sprintf("Inavlid envelope %v\n", invalidEnvelope))
		}

		log.Println("PeerRecord", peerRecord.PeerID)
	default:
	}
}