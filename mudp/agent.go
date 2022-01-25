package main

import (
	"context"
	"net"
	"sync"
	"time"

	capnp "capnproto.org/go/capnp/v3"
	"github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p-core/record"
)

const (
	multicastAddr = "224.0.1.241:3037"
)

type Mudp struct {
	h             host.Host
	mc            multicaster
	disc          discovery.Discoverer
	mustFind      map[string]chan peer.AddrInfo
	mustAdvertise map[string]chan time.Duration
	mu            sync.Mutex
}

func NewMudp(disc discovery.Discoverer) (mudp Mudp, err error) {
	mc, err := NewMulticaster(multicastAddr)
	if err != nil {
		return
	}
	mc.Listen(mudp.multicastHandler)

	return Mudp{mc: mc, disc: disc, mustFind: make(map[string]chan peer.AddrInfo)}, nil
}

func (mudp *Mudp) multicastHandler(addr *net.UDPAddr, n int, buffer []byte) {
	msg, err := capnp.UnmarshalPacked(buffer[:n])
	if err != nil {
		return
	}

	root, err := ReadRootMudpPacket(msg)
	if err != nil {
		return
	}
	switch root.Which() {
	case MudpPacket_Which_request:
		request, err := root.Request()
		if err != nil {
			return
		}
		mudp.handleMudpRequest(request)
	case MudpPacket_Which_response:
		response, err := root.Response()
		if err != nil {
			return
		}
		mudp.handleMudpResponse(response)
	default:
	}
}

func (mudp *Mudp) handleMudpRequest(request MudpRequest) {
	mudp.mu.Lock()
	defer mudp.mu.Unlock()

	ns, err := request.Namespace()
	if err != nil {
		return
	}

	if _, ok := mudp.mustAdvertise[ns]; !ok {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	finder, err := mudp.disc.FindPeers(ctx, ns)
	if err != nil {
		return
	}
	peers := make([]peer.AddrInfo, 0)
	for peer := range finder {
		peers = append(peers, peer)
	}

	response, err := mudp.buildResponse(ns, peers)
	if err == nil {
		mudp.mc.Multicast(response)
	}
}

func (mudp *Mudp) buildResponse(ns string, peers []peer.AddrInfo) ([]byte, error) {
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	root, err := NewRootMudpPacket(seg)
	if err != nil {
		panic(err)
	}
	response, err := root.NewResponse()

	response.SetNamespace(ns)
	envelope, err := response.NewEnvelope(int32(len(peers) + 1))

	if cab, ok := peerstore.GetCertifiedAddrBook(mudp.h.Peerstore()); ok {
		peers = append(peers, *host.InfoFromHost(mudp.h)) // Add itself
		i := 0
		for _, info := range peers {
			env := cab.GetPeerRecord(info.ID)
			rec, err := env.Marshal()
			if err == nil {
				envelope.Set(i, rec)
				i++
			}
		}
	}
	return root.Message().MarshalPacked()
}

func (mudp *Mudp) handleMudpResponse(response MudpResponse) {
	var (
		finder chan peer.AddrInfo
		ok     bool
	)

	mudp.mu.Lock()
	defer mudp.mu.Unlock()

	ns, err := response.Namespace()
	if err != nil {
		return
	}

	if finder, ok = mudp.mustFind[ns]; !ok {
		return
	}

	var envelope capnp.DataList
	if envelope, err = response.Envelope(); err != nil {
		return
	}

	for i := 0; i < envelope.Len(); i++ {
		var rec peer.PeerRecord

		rawEnvelope, err := envelope.At(0)
		if err != nil {
			continue
		}
		if _, err = record.ConsumeTypedEnvelope(rawEnvelope, &rec); err != nil {
			continue
		}

		finder <- peer.AddrInfo{ID: rec.PeerID, Addrs: rec.Addrs}

		// TODO: handle when there is no consumer
	}
}

func (mudp *Mudp) Advertise(ctx context.Context, ns string, opt ...discovery.Option) (time.Duration, error) {
	mudp.mu.Lock()
	defer mudp.mu.Unlock()

	ttl := getTtl(opt)
	if ttlChan, ok := mudp.mustAdvertise[ns]; ok {
		ttlChan <- ttl
	} else {
		resetTtl := make(chan time.Duration)
		mudp.mustAdvertise[ns] = resetTtl
		go mudp.trackNamespace(ns, resetTtl, ttl)
	}
	return ttl, nil
}

func getTtl(opt []discovery.Option) time.Duration {
	panic("unimplemented")
}

func (mudp *Mudp) trackNamespace(ns string, resetTtl chan time.Duration, ttl time.Duration) {
	timer := time.NewTimer(ttl)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			mudp.mu.Lock()
			defer mudp.mu.Unlock()

			select { // check again TTL after acquiring lock
			case ttl := <-resetTtl:
				timer.Reset(ttl)
			default:
				delete(mudp.mustFind, ns)
				return
			}
		case ttl := <-resetTtl:
			timer.Reset(ttl)
		}
	}
}

func (mudp *Mudp) FindPeers(ctx context.Context, ns string, opt ...discovery.Option) (<-chan peer.AddrInfo, error) {
}
