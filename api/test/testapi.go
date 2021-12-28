package test

import "git.jetbrains.space/yabbi/dsp/components/extrarpc"

type Test struct {
	rpcBidder *extrarpc.Client
}

func New(rpcBidder *extrarpc.Client) *Test {
	return &Test{rpcBidder: rpcBidder}
}
