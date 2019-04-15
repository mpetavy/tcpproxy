package main

import (
	"flag"
	"github.com/google/tcpproxy"
	"github.com/mpetavy/common"
	"time"
)

var (
	source *string
	dest   *string

	proxy tcpproxy.Proxy
)

func init() {
	source = flag.String("s", "", "server socket host address")
	dest = flag.String("d", "", "destination socket host address")
}

func start() error {
	proxy.AddRoute(*source, tcpproxy.To(*dest))

	var err error

	go func(err *error) {
		*err = proxy.Run()
	}(&err)

	time.Sleep(time.Second)

	return nil
}

func stop() error {
	return proxy.Close()
}

func main() {
	defer common.Cleanup()

	common.New(&common.App{"tcpproxy", "1.0.0", "2018", "tcpproxy", "mpetavy", common.APACHE, "https://github.com/mpetavy/hl7send", true, nil,start, stop, nil, 0}, []string{"s", "d"})
	common.Run()
}
