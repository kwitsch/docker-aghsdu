package main

import (
	"aghsdu/aghapi"
	"aghsdu/config"
	"aghsdu/processor"
	"aghsdu/sc"
	"fmt"
)

func main() {
	conf := config.Get()
	aghc := sc.New("AGH", conf.AghContainer, conf.Verbose)
	dnsc := sc.New("DNS", conf.DNSContainer, conf.Verbose)
	aapi := aghapi.New(conf.AghHost, conf.AghUser, conf.AghPass, conf.Verbose)
	fmt.Println("---------------------------------")
	fmt.Println("- AdGuardHome Swarm Dns Updater -")
	fmt.Println("---------------------------------")
	for {
		aghc.Lookup()
		dnsc.Lookup()
		if !aghc.Active && !dnsc.Active {
			fmt.Println(aghc.Name, "&", dnsc.Name, "containers are unreachable")
		} else if !aghc.Active {
			fmt.Println(aghc.Name, "container is unreachable")
		} else if !dnsc.Active {
			fmt.Println(dnsc.Name, "container is unreachable")
		} else {
			if aghc.IPChange || dnsc.IPChange {
				fmt.Println("-- IP Change detected")
				aghc.Print()
				dnsc.Print()

				gsuc, oldinfo := aapi.GetDnsInfo()
				if gsuc {
					psuc, newinfo := processor.GenerateNew(oldinfo, aghc.Addresses, dnsc.Addresses, conf.Verbose)
					if psuc {
						aapi.SetDnsInfo(newinfo)
					}
				}
			} else {
				conf.VPrint("-- IP diden't change")
			}

		}
		conf.VPrint("------- Iteration End -------")
		conf.LoopWait()
	}
}
