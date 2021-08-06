package main

import (
	"aghsdu/aghapi"
	"aghsdu/config"
	"aghsdu/processor"
	"aghsdu/sc"
	"fmt"
	"os"
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
		if !aghc.Active || !dnsc.Active {
			fmt.Println(aghc.Name, "or", dnsc.Name, "containers are unreachable")
			os.Exit(1)
		} else {
			if aghc.IPChange || dnsc.IPChange {
				fmt.Println("-- IP Change detected")
				aghc.Print()
				dnsc.Print()

				gsuc, oldinfo := aapi.GetDnsInfo()
				if gsuc {
					psuc, newinfo := processor.GenerateNew(oldinfo, aghc.Addresses, dnsc.Addresses, conf.ContainerOnly, conf.Verbose)
					if psuc {
						aapi.SetDnsInfo(newinfo)
					}
				}
			} else {
				conf.VPrint("-- IP diden't change")
				dnsc.Print()
			}

		}
		conf.VPrint("------- Iteration End -------")
		conf.LoopWait()
	}
}
