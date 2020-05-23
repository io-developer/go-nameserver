package dns

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/miekg/dns"
)

// Strategy
const (
	StrategyCheck        = "check"
	StrategyCheckAndStop = "check_and_stop"
	StrategySkipOrCheck  = "skip_or_check"
)

// Record ..
type Record struct {
	Strategy string `json:"strategy,omitempty"`
	IP       string `json:"ip,omitempty"`
	Answer   string `json:"answer"`
}

// RecordMapLoadJSON ..
func RecordMapLoadJSON(jsonPath string) (map[string][]Record, error) {
	bytes, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	var recMap map[string][]Record
	json.Unmarshal(bytes, &recMap)

	return recMap, nil
}

// RecordHandleAnswer ..
func RecordHandleAnswer(m *dns.Msg, q dns.Question, record Record, answersCount int) (bool, bool) {
	handled := false
	stopHandle := false
	isChecked := true
	skip := answersCount > 0 && record.Strategy == StrategySkipOrCheck
	switch record.Strategy {
	case StrategyCheck, StrategyCheckAndStop, StrategySkipOrCheck:
		isChecked = !skip && UtilPingHost(record.IP, 1)
	}
	if isChecked {
		stopHandle = record.Strategy == StrategyCheckAndStop
		params := map[string]interface{}{
			"domain": q.Name,
			"ip":     record.IP,
		}
		rr, err := dns.NewRR(UtilRenderTpl(record.Answer, params))
		if err == nil {
			m.Answer = append(m.Answer, rr)
			handled = true
		}
	}
	return handled, stopHandle
}
