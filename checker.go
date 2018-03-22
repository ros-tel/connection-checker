package main

import (
	"crypto/tls"
	"log"
	"net"
	"strings"
	"time"
)

func checkerRun(hosts []string) {
	for {
		var conn net.Conn
		var err error
		var ips []string

		for _, v := range hosts {
			h := strings.Split(v, ":")
			ips, err = net.LookupHost(h[0])
			if err != nil {
				log.Printf("[ERROR] Lookup Host %s error: %v\n", h[0], err)
				continue
			}
			host := ips[0] + ":" + h[1]

			if *use_tls {
				dialer := &net.Dialer{Timeout: 5 * time.Second}
				conn, err = tls.DialWithDialer(dialer, "tcp", host, &tls.Config{InsecureSkipVerify: false})
			} else {
				conn, err = net.DialTimeout("tcp", host, 5*time.Second)
			}
			if err != nil {
				log.Printf("[ERROR] dial to %s usage tls %t failed error: %v\n", host, *use_tls, err)
				continue
			}

			conn.Close()

			log.Printf("[OK] dial to %s usage tls %t remote addr %s\n", host, *use_tls, conn.RemoteAddr())
		}

		time.Sleep(10 * time.Second)
	}
}
