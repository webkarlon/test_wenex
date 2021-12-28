package test

import (
	"fmt"
	"git.jetbrains.space/yabbi/dsp/bidder/api"
	"net/http"
)

func (t *Test) TestQuery(w http.ResponseWriter, r *http.Request) {
	var s string
	var call api.Call

	if err := t.rpcBidder.Call("TestServer.Query", call, &s); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)
}
