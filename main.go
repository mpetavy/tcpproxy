package main

import (
	"flag"
	"fmt"
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
	common.Init(false, "1.0.0", "", "", "2018", "tcpproxy", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, start, stop, nil, 0)

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
	defer common.Done()

	common.Run([]string{"s", "d"})
}
