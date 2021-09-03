package sc

import (
	"context"
	"fmt"
	"net"
	"time"
)

type Container struct {
	Name          string
	Hostname      string
	Active        bool
	IPChange      bool
	Addresses     []net.IP
	LastAddresses []net.IP
	Verbose       bool
}

func New(name, hostname string, verbose bool) Container {
	res := Container{
		name,
		hostname,
		false,
		false,
		[]net.IP{},
		[]net.IP{},
		verbose,
	}
	return res
}

func (c *Container) Lookup() {
	c.Active = false
	c.Addresses = []net.IP{}
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "127.0.0.11:53")
		},
	}
	ips, err := r.LookupIP(context.Background(), "ip4", c.Hostname)
	if err == nil ||
		len(ips) > 0 ||
		(len(ips) == 1 && ips[0].String() != "0.0.0.0") {
		c.Active = true
		c.Addresses = ips
	}
	c.IPChange = !Equal(c.LastAddresses, c.Addresses)
	if c.IPChange {
		c.LastAddresses = c.Addresses
	}
}

func Equal(last, current []net.IP) bool {
	if len(last) != len(current) {
		return false
	} else {
		for _, i := range current {
			if !Contains(last, i) {
				return false
			}
		}
		return true
	}
}

func Contains(ia []net.IP, ci net.IP) bool {
	for _, i := range ia {
		if i.String() == ci.String() {
			return true
		}
	}
	return false
}

func (c Container) Print() {
	if c.Verbose {
		fmt.Println("-", c.Name)
		fmt.Println("| - Hostname:", c.Hostname)
		fmt.Println("| - Active:", c.Active)
		if c.Active {
			fmt.Println("| - IPs:")
			for _, ip := range c.Addresses {
				fmt.Println("| | ", ip.String())
			}
		}
	}
}
