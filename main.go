package main

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/red-life/filternet/cmd"
	"github.com/red-life/filternet/pkg/dns"
	"os"
	"strings"
)

var (
	configFile, domains string
	noColor             bool
)

func main() {
	var config cmd.Config
	flag.StringVar(&configFile, "config", "default.json", "Config file path which contains resolvers")
	flag.StringVar(&domains, "domains", "twitter.com,instagram.com,facebook.com", "Domains to check. Separated with ,")
	flag.BoolVar(&noColor, "no-color", false, "Disable color output")
	flag.Parse()
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		panic("Config file doesn't exist.")
	}
	domainsSplitted := strings.Split(domains, ",")
	content, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}
	resolver := dns.NewResolver(dns.DohResolver, dns.DouResolver)
	checker := cmd.NewChecker(domainsSplitted, config, noColor, resolver)
	checker.Run()
}
