package dns

import (
	"errors"
	"log"

	"github.com/miekg/dns"
)

// Main props
var (
	RecordsMap   map[string][]Record
	UpstreamAddr string
	IsVerbose    bool
)

// ServerLoadRecords ..
func ServerLoadRecords(jsonPath string) {
	records, err := RecordMapLoadJSON(jsonPath)
	if err != nil {
		log.Fatal("Failed to load and parse json records: ", err)
	}
	RecordsMap = records
}

// ServerStart ..
func ServerStart(addr string) {
	if RecordsMap == nil {
		log.Fatal("Failed to start server: records not set or not loaded")
	}

	dns.HandleFunc(".", handleDNSRequest)

	server := &dns.Server{Addr: addr, Net: "udp"}
	log.Println("Listening ", addr)

	verbosef(
		"\n  IsVerbose: %t\n  UpstreamAddr: %s\n  Records: %v\n",
		IsVerbose,
		UpstreamAddr,
		RecordsMap,
	)

	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	verbosef("Handling request: \n%v\n", r)

	resp := new(dns.Msg)
	resp.SetReply(r)
	resp.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		if handleQuery(resp) {
			w.WriteMsg(resp)
			return
		}
	}

	upstreamResp, err := handleUpstream(r)

	if err != nil {
		verbose("Upstream response: ", upstreamResp)
		verbose("Upstream error: ", err)
	}

	if upstreamResp != nil && err == nil {
		resp = upstreamResp
	}

	verbose("Response: ", resp)

	w.WriteMsg(resp)
}

func handleQuery(m *dns.Msg) bool {
	queryHandled := false
	for _, q := range m.Question {
		verbosef("Handling query %s, type: %d", q.Name, q.Qtype)

		if records, exists := RecordsMap[q.Name]; exists {
			answersCount := 0
			for _, record := range records {
				handled, stopHandle := RecordHandleAnswer(m, q, record, answersCount)

				verbosef(
					"Trying record: %v\n  (handled: %t, stop: %t)",
					record,
					handled,
					stopHandle,
				)

				if handled {
					queryHandled = true
					answersCount++
				}
				if stopHandle {
					break
				}
			}
		}
	}
	return queryHandled
}

func handleUpstream(req *dns.Msg) (*dns.Msg, error) {
	if UpstreamAddr == "" {
		verbose("Upstream not defined. Skipping..")
		return nil, nil
	}

	verbose("Handling upstream ", UpstreamAddr)

	c := new(dns.Client)
	resp, _, err := c.Exchange(req, UpstreamAddr)
	if err == nil && resp.Rcode != dns.RcodeSuccess {
		return resp, errors.New("Response code is not success")
	}
	return resp, err
}

func verbose(args ...interface{}) {
	if IsVerbose {
		log.Println(args...)
	}
}

func verbosef(format string, args ...interface{}) {
	if IsVerbose {
		log.Printf(format+"\n", args...)
	}
}
