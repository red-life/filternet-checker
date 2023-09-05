package dns

import (
	"github.com/miekg/dns"
	"net"
)

func DouResolver(domain, server string) ([]net.IP, error) {
	m := new(dns.Msg)
	m.SetQuestion(domain+".", dns.TypeA)
	m.RecursionDesired = true
	client := dns.Client{}
	resp, _, err := client.Exchange(m, server)
	if err != nil {
		return []net.IP{}, err
	}
	ips := make([]net.IP, 0)
	for _, answer := range resp.Answer {
		switch answer.(type) {
		case *dns.A:
			ips = append(ips, answer.(*dns.A).A)
		}
	}
	return ips, nil
}
