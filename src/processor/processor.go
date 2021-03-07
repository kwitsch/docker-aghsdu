package processor

import (
	"encoding/json"
	"fmt"
	"net"
)

type DnsInfo struct {
	Upstream_dns       []string
	Upstream_dns_file  string
	Bootstrap_dns      []string
	Protection_enabled bool
	Ratelimit          int
	Blocking_mode      string
	Blocking_ipv4      string
	Blocking_ipv6      string
	Edns_cs_enabled    bool
	Dnssec_enabled     bool
	Disable_ipv6       bool
	Upstream_mode      string
	Cache_size         int
	Cache_ttl_min      int
	Cache_ttl_max      int
}

func GenerateNew(old string, agh, dns []net.IP, verbose bool) (bool, string) {
	var dnsconf DnsInfo
	json.Unmarshal([]byte(old), &dnsconf)

	if verbose {
		fmt.Println("Old DNS:")
		PrintArr(dnsconf.Upstream_dns)
	}
	dnsar := RemoveAll(dnsconf.Upstream_dns, agh)
	dnsar = AddAll(dnsar, dns)

	if verbose {
		fmt.Println("New DNS:")
		PrintArr(dnsar)
	}
	dnsconf.Upstream_dns = dnsar

	b, err := json.Marshal(dnsconf)
	if err != nil {
		return false, ""
	}
	return true, string(b)
}

func RemoveAll(cur []string, agh []net.IP) []string {
	res := []string{}
	for _, i := range cur {
		if !IContains(agh, i) {
			res = append(res, i)
		}
	}
	return res
}

func AddAll(cur []string, dns []net.IP) []string {
	for _, i := range dns {
		if !SContains(cur, i.String()) {
			cur = append(cur, i.String())
		}
	}
	return cur
}

func IContains(ia []net.IP, ci string) bool {
	for _, i := range ia {
		if i.String() == ci {
			return true
		}
	}
	return false
}

func SContains(ia []string, ci string) bool {
	for _, i := range ia {
		if i == ci {
			return true
		}
	}
	return false
}

func PrintArr(strar []string) {
	for _, s := range strar {
		fmt.Println(" ", s)
	}
}
