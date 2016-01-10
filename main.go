package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/jkiet/go-pie/comet"
	"github.com/jkiet/go-pie/driver"
	"github.com/jkiet/go-pie/listener"
	"net"
	"net/http"
	"os"
	"strings"
)

func usage() {
	s := strings.Split(os.Args[0], "/")
	appName := s[len(s)-1]

	fmt.Printf(`Raspbery Pi lights controller
usage:		
%[1]s listen <ip>:<port>
%[1]s comet  <url>

examples:
%[1]s listen 0.0.0.0:8888
%[1]s comet  http://192.168.1.121:8080/lamps

`, appName)
}

func main() {
	if len(os.Args) != 3 {
		usage()
		return
	}

	section := driver.NewSection()
	err := section.Init()
	if err != nil {
		fmt.Printf("Can't init, err: %v", err)
		return
	}

	switch strings.ToLower(os.Args[1]) {
	case "listen":
		bind := strings.Split(os.Args[2], ":")
		if len(bind) != 2 {
			usage()
			return
		}
		listenerSvc := listener.NewListenerSvc(section)
		wsContainer := restful.NewContainer()
		listenerSvc.Register(wsContainer)

		listenAddress := net.JoinHostPort(bind[0], bind[1])
		httpd := &http.Server{Addr: listenAddress, Handler: wsContainer}
		httpd.ListenAndServe()
		return
	case "comet":
		cometClient := comet.NewCometClient(section)
		cometClient.Listen(os.Args[2])
		return
	}
	usage()
}
