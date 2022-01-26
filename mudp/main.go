package main

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/event"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	mx "github.com/wetware/matrix/pkg"
)

const (
	ns = "casm/mudp"
)

type MockDisc struct {
	h host.Host
}

func (md *MockDisc) FindPeers(ctx context.Context, ns string, opts ...discovery.Option) (<-chan peer.AddrInfo, error) {
	finder := make(chan peer.AddrInfo, 1)
	finder <- peer.AddrInfo{ID: md.h.ID(), Addrs: md.h.Addrs()}
	close(finder)
	return finder, nil

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sim := mx.New(ctx)
	h1 := sim.MustHost(ctx)
	h2 := sim.MustHost(ctx)
	waitReady(h1)
	waitReady(h2)

	a1, err := NewMudp(h1, &MockDisc{h: h1})
	if err != nil {
		panic(err)
	}
	defer a1.Close()

	a2, err := NewMudp(h2, &MockDisc{h: h2})
	if err != nil {
		panic(err)
	}
	defer a2.Close()

	a1.Advertise(ctx, ns, discovery.TTL(time.Second))
	a2.Advertise(ctx, ns, discovery.TTL(time.Second))
	println("Called Advertise")

	finder, err := a1.FindPeers(ctx, ns, discovery.TTL(time.Second))
	if err != nil {
		panic(err)
	}
	println("Called FindPeers")

	for info := range finder {
		println("Received peer", info.ID[:5])
	}
}

func newHost() host.Host {
	h, err := libp2p.New(context.Background())
	if err != nil {
		panic(err)
	}
	defer h.Close()

	waitReady(h)
	return h
}

func waitReady(h host.Host) {
	sub, err := h.EventBus().Subscribe(new(event.EvtLocalAddressesUpdated))
	if err != nil {
		panic(err)
	}
	defer sub.Close()

	<-sub.Out()
}
