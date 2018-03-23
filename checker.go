package main

import (
	"crypto/tls"
	"log"
	"net"
	"time"
)

func checkerRun(hosts []host) {
	for {
		var conn net.Conn
		var err error
		var ips []string

		for _, v := range hosts {
			ips, err = net.LookupHost(v.host)
			if err != nil {
				log.Printf("[ERROR] Lookup Host %s error: %v\n", v.host, err)
				continue
			}
			for _, ip := range ips {
				host := ip + ":" + v.port

				if *use_tls {
					dialer := &net.Dialer{Timeout: 5 * time.Second}
					conn, err = tls.DialWithDialer(dialer, "tcp", host, &tls.Config{InsecureSkipVerify: false, ServerName: v.host})
				} else {
					conn, err = net.DialTimeout("tcp", host, 5*time.Second)
				}
				if err != nil {
					log.Printf("[ERROR] dial to %s:%s usage tls %t failed error: %v\n", v.host, v.port, *use_tls, err)
					continue
				}

				conn.Close()

				log.Printf("[OK] dial to %s:%s usage tls %t remote addr %s\n", v.host, v.port, *use_tls, conn.RemoteAddr())
			}
		}

		time.Sleep(10 * time.Second)
	}
}
