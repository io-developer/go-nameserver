package main

import (
	"flag"

	"github.com/io-developer/go-nameserver/dns"
)

func main() {
	addr := flag.String("listen", "0.0.0.0:53", "Listen address:port")
	upstreamAddr := flag.String("upstream", "8.8.8.8:53", "Foreign DNS server address:port")
	recordsPath := flag.String("records", "records.json", "JSON path")
	verbose := flag.Bool("verbose", false, "Verbose logging")
	flag.Parse()

	dns.IsVerbose = *verbose
	dns.UpstreamAddr = *upstreamAddr
	dns.ServerLoadRecords(*recordsPath)
	dns.ServerStart(*addr)
}
