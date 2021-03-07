package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type config struct {
	DNSContainer string
	AghContainer string
	AghHost      string
	AghUser      string
	AghPass      string
	TimerLoop    int
	Verbose      bool
}

func Get() config {
	res := config{
		strings.TrimSpace(os.Getenv("DNS_CONTAINER")),
		strings.TrimSpace(os.Getenv("AGH_CONTAINER")),
		"http",
		strings.TrimSpace(os.Getenv("AGH_USER")),
		strings.TrimSpace(os.Getenv("AGH_PASSWORD")),
		10,
		false,
	}
	secval, secerr := strconv.ParseBool(os.Getenv("AGH_SECURE"))
	if secerr == nil && secval == true {
		res.AghHost += "s"
	}
	res.AghHost += "://" + res.AghContainer + ":" + strings.TrimSpace(os.Getenv("AGH_PORT")) + "/control/"
	loopval, looperr := strconv.Atoi(os.Getenv("TIMER_LOOP"))
	if looperr == nil && loopval > 0 {
		res.TimerLoop = loopval
	}
	verbval, verberr := strconv.ParseBool(os.Getenv("VERBOSE"))
	if verberr == nil {
		res.Verbose = verbval
	}
	if res.Verbose {
		fmt.Println("--- Configuration ---")
		fmt.Println("- DNS Container:", res.DNSContainer)
		fmt.Println("- Agh Container:", res.AghContainer)
		fmt.Println("- Agh Host:", res.AghHost)
		fmt.Println("- Loop every", res.TimerLoop, "seconds")
		fmt.Println("---------------------")
	}
	return res
}

func (c config) VPrint(text string) {
	if c.Verbose {
		fmt.Println(text)
	}
}

func (c config) LoopWait() {
	s := c.TimerLoop
	for {
		if s <= 0 {
			break
		} else {
			time.Sleep(1 * time.Second)
			s--
		}
	}
}
