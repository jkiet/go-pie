package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/jkiet/go-pie/comet"
	"github.com/jkiet/go-pie/config"
	"github.com/jkiet/go-pie/driver"
	"github.com/jkiet/go-pie/listener"
	"io/ioutil"
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
%[1]s listen <ip>:<port> [<config filename>]
%[1]s comet  <url> [<config filename>]

examples:
%[1]s listen 0.0.0.0:8888 config.yml
%[1]s comet  http://192.168.1.121:8080/lamps

`, appName)
}

var defaultCfgYml = `
section: 0
layout: [4, 17, 27, 22, 18, 23, 24, 25]
`

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	var err = error(nil)
	cfgYml := []byte(defaultCfgYml)
	if len(os.Args) == 4 {
		cfgYml, err = ioutil.ReadFile(os.Args[3])
		if err != nil {
			fmt.Printf("Can't read file %v, err: %v", os.Args[3], err)
			os.Exit(1)
		}
	}
	cfg := config.NewConfig()
	err = cfg.Read(cfgYml)
	if err != nil {
		fmt.Printf("Can't read config; err: %v ; yaml: %v", err, cfgYml)
		os.Exit(1)
	}

	section := driver.NewSection()
	err = section.Init(cfg)
	if err != nil {
		fmt.Printf("Can't init, err: %v", err)
		os.Exit(1)
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
