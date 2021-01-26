package main

import (
	"flag"
	"fmt"
	"github.com/saqfish/go-xinput"
	"os"
	"os/signal"
	"strings"
)

func main() {
	var keys []string
	var shift bool
	var ctrl bool
	var alt bool

	deviceOpt := flag.Int("dev", -1, "Device number")
	delimiterOpt := flag.String("d", "", "character between keys")
	keyCountOpt := flag.Int("c", 10, "key count")
	reverseOpt := flag.Bool("r", false, "reverse key list")

	flag.Parse()

	if *deviceOpt == -1 {
		fmt.Println("-d [id] required!")
		os.Exit(1)
	}

	dp := xinput.XOpenDisplay(nil)
	defer xinput.XCloseDisplay(dp)

	for _, dv := range xinput.GetXDeviceInfos(dp) {
		if dv.Id != uint64(*deviceOpt) {
			continue
		}

		eventMap, err := xinput.NewEventMap(dp, dv)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		stopChan := make(chan os.Signal, 1)
		signal.Notify(stopChan, os.Interrupt)

	loop:
		for {
			select {
			case event := <-eventMap.Events():
				var k string
				if _, ok := shiftkeys[int(event.Field)]; ok {
					shift = !shift
				}
				if _, ok := ctrlkeys[int(event.Field)]; ok {
					ctrl = !ctrl
				}
				if _, ok := altkeys[int(event.Field)]; ok {
					alt = !alt
				}
				if kv, ok := keymap[int(event.Field)]; ok {
					var i int
					var prefix string

					if shift {
						i = 1
					}
					if ctrl {
						prefix += "Ctrl-"
					}

					if alt {
						prefix += "Alt-"
					}

					if event.Type == xinput.KeyReleaseEvent {
						k = fmt.Sprintf("%s%s%s", prefix, kv[i], *delimiterOpt)
					}

					count := *keyCountOpt * 2

					if len(keys) > count {
						if *reverseOpt {
							keys = keys[:count]
						} else {
							keys = keys[1:]
						}
					}
					if *reverseOpt {
						keys = append([]string{k}, keys...)
					} else {
						keys = append(keys, k)
					}
					s := strings.Join(keys, "")

					fmt.Printf("%s\n", s)
				}
			case <-stopChan:
				break loop
			}
		}

		_ = eventMap.Close()
	}
}
