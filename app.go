package main

import (
	"flag"
	"log"
	"strings"
)

var (
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

	hosts := strings.Split(*param_hosts, ",")

	checkerRun(hosts)
}
