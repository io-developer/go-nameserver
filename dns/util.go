package dns

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os/exec"
	"strings"
)

// UtilPingHost ..
func UtilPingHost(host string, maxSeconds int) bool {
	cmd := fmt.Sprintf("ping %s -c 1 -w %d > /dev/null && echo reachable || echo 0", host, maxSeconds)
	output, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		log.Print(err)
		return false
	}
	isReachable := strings.HasPrefix(string(output), "reachable")
	log.Printf("Ping %s: %t", host, isReachable)
	return isReachable
}

// UtilRenderTpl ..
func UtilRenderTpl(tpl string, data map[string]interface{}) string {
	t := template.Must(template.New("").Parse(tpl))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		return ""
	}
	return buf.String()
}
