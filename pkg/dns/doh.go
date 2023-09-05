package dns

import (
	"bytes"
	"github.com/miekg/dns"
	"io"
	"net"
	"net/http"
)

func DohResolver(domain, server string) ([]net.IP, error) {
	m := new(dns.Msg)
	m.SetQuestion(domain+".", dns.TypeA)
	m.RecursionDesired = true
	msg, err := m.Pack()
	if err != nil {
		return []net.IP{}, err
	}
	request, err := http.NewRequest("POST", server, bytes.NewReader(msg))
	request.Header.Set("Accept", "application/dns-message")
	request.Header.Set("Content-Type", "application/dns-message")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return []net.IP{}, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	m2 := new(dns.Msg)
	m2.Unpack(body)
	ips := make([]net.IP, 0)
	for _, answer := range m2.Answer {
		switch answer.(type) {
		case *dns.A:
			ips = append(ips, answer.(*dns.A).A)
		}
	}
	return ips, nil
}
