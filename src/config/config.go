package config

import (
	"fmt"
	"time"

	EnvWrapper "github.com/kwitsch/go-env_wrapper"
)

type config struct {
	DNSContainer  string
	AghContainer  string
	AghHost       string
	AghUser       string
	AghPass       string
	TimerLoop     int
	ContainerOnly bool
	Verbose       bool
}

func Get() config {
	env := EnvWrapper.Default()
	res := config{
		env.GetString("DNS_CONTAINER"),
		env.GetString("AGH_CONTAINER"),
		"http",
		env.GetString("AGH_USER"),
		env.GetString("AGH_PASSWORD"),
		env.GetIntDef("TIMER_LOOP", 10),
		env.GetBool("CONTAINER_ONLY"),
		env.GetBool("VERBOSE"),
	}

	if env.GetBool("AGH_SECURE") {
		res.AghHost += "s"
	}
	res.AghHost += "://" + res.AghContainer + ":" + env.GetStringDef("AGH_PORT", "80") + "/control/"

	if res.Verbose {
		fmt.Println("--- Configuration ---")
		fmt.Println("- DNS Container:", res.DNSContainer)
		fmt.Println("- Agh Container:", res.AghContainer)
		fmt.Println("- Agh Host:", res.AghHost)
		fmt.Println("- Container only:", res.ContainerOnly)
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
