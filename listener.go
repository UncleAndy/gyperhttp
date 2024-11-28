package gyperhttp

import (
	"net"

	"github.com/cloudwego/netpoll"
)

func Listen(proto, addr string) (net.Listener, error) {
	return netpoll.CreateListener(proto, addr)
}
