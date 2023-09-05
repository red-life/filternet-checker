package dns

import (
	"net"
	"regexp"
	"strings"
)

type Resolve func(domain, server string) ([]net.IP, error)

var (
	serverAddressRegex = regexp.MustCompile(`(?m)^(https://|tls://)?((\*)|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)|((\*\.)?([a-zA-Z0-9-]+\.){0,5}[a-zA-Z0-9-][a-zA-Z0-9-]+\.[a-zA-Z]{2,63}?))(:[\d]+)?`)
	domainAddressRegex = regexp.MustCompile(`(?m)^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]{0,1}\.([a-zA-Z]{1,6}|[a-zA-Z0-9-]{1,30}\.[a-zA-Z]{2,3})$`)
)

func NewResolver(dohResolver, douResolver Resolve) *Resolver {
	return &Resolver{
		dohResolver:        dohResolver,
		douResolver:        douResolver,
		serverAddressRegex: serverAddressRegex,
		domainRegex:        domainAddressRegex,
	}
}

type Resolver struct {
	dohResolver, douResolver        Resolve
	serverAddressRegex, domainRegex *regexp.Regexp
}

func (r Resolver) Resolve(domain, server string) ([]net.IP, error) {
	if !r.domainRegex.Match([]byte(domain)) {
		return []net.IP{}, InvalidDomainErr
	}
	if !r.serverAddressRegex.Match([]byte(server)) {
		return []net.IP{}, InvalidServerErr
	}
	resolver := r.getResolver(server)
	return resolver(domain, server)
}

func (r Resolver) getResolver(server string) Resolve {
	switch {
	case strings.HasPrefix(server, "https://"):
		return r.dohResolver
	default:
		return r.douResolver

	}
}
