package cmd

import (
	"github.com/fatih/color"
	"github.com/red-life/filternet/pkg/dns"
	"net"
)

var (
	FilternetIPs = []net.IP{
		net.ParseIP("10.10.34.34"),
		net.ParseIP("10.10.34.35"),
		net.ParseIP("10.10.34.36"),
	}
)

type Resolver struct {
	Name    string   `json:"name"`
	Servers []string `json:"servers"`
}

type Config struct {
	Resolvers []Resolver `json:"resolvers"`
}

func NewChecker(domains []string, config Config, noColor bool, resolver *dns.Resolver) *Checker {
	return &Checker{
		domains:  domains,
		config:   config,
		noColor:  noColor,
		resolver: resolver,
	}
}

type Checker struct {
	domains  []string
	config   Config
	noColor  bool
	resolver *dns.Resolver
}

func (c *Checker) Run() {
	color.NoColor = c.noColor
	c.run()
}

func (c *Checker) run() {
	for i, resolver := range c.config.Resolvers {
		color.Magenta("[%d] %s", i+1, resolver.Name)
		for _, domain := range c.domains {
			isDomainCensored := false
			filteredInResolvers := 0
			for _, server := range resolver.Servers {
				result, err := c.resolver.Resolve(domain, server)
				if err != nil {
					color.Red(err.Error())
					isDomainCensored = true
					filteredInResolvers++
					continue
				}
				if len(result) == 0 {
					color.Red("%s: Domain %s not found", server, domain)
					continue
				}
				if len(result) == 1 && isIPInFilternet(result[0]) { // Filternet always answers with only one IP
					isDomainCensored = true
					filteredInResolvers++
				}
			}
			if isDomainCensored {
				color.Red("[-] %s is filtered! (%d/%d)", domain, filteredInResolvers, len(resolver.Servers))
			} else {
				color.Green("[+] %s is not filtered!", domain)
			}

		}
	}
}

func isIPInFilternet(ip net.IP) bool {
	for _, filterNetIP := range FilternetIPs {
		if ip.Equal(filterNetIP) {
			return true
		}
	}
	return false
}
