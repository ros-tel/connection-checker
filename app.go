package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
)

type (
	host struct {
		host string
		port string
	}
)

var (
	hosts []host

	param_hosts = flag.String("hosts", "", "Usage: -hosts=google.com:80 or -hosts=87.250.250.242:80,google.com:80")
	use_tls     = flag.Bool("tls", false, "Usage: -tls")
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	flag.Parse()

	if *param_hosts == "" {
		flag.PrintDefaults()
		log.Fatal("Required parameter \"hosts\" not set")
	}

	ph := strings.Split(*param_hosts, ",")
	for _, hv := range ph {
		h := strings.Split(hv, ":")

		// No port specified. Use 80 as the default port
		if len(h) == 1 {
			hosts = append(hosts, host{host: h[0], port: "80"})
			continue
		}

		port, err := strconv.Atoi(h[1])
		if err != nil || port < 1 || 65535 < port {
			log.Printf("Bad port %s\n", hv)
			continue
		}

		hosts = append(hosts, host{host: h[0], port: h[1]})
	}

	if len(hosts) == 0 {
		log.Fatal("No host is specified for testing")
	}

	checkerRun(hosts)
}
