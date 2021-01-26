package main

import (
	"flag"
	"fmt"
	"github.com/saqfish/go-xinput"
	"os"
	"os/signal"
	"strings"
)

type Key struct {
	l int
	s string
}

func main() {
	var keys []Key
	var shift bool

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
				var k Key
				var tk string

				if _, ok := shiftkeys[int(event.Field)]; ok {
					shift = !shift
				}

				if event.Type == xinput.KeyPressEvent {
					var i int

					if shift {
						i = 1
					}

					if kv, ok := keymap[int(event.Field)]; ok {

						tk = fmt.Sprintf("%s%s", kv[i], *delimiterOpt)

					} else {
						pkeys(keys)
					}

					k = Key{len(tk), tk}

					count := *keyCountOpt

					if len(keys) > count {
						if *reverseOpt {
							keys = keys[:count]
						} else {
							keys = keys[1:]
						}
					}
					if *reverseOpt {
						keys = append([]Key{k}, keys...)
					} else {
						keys = append(keys, k)
					}
					pkeys(keys)

				}
			case <-stopChan:
				break loop
			}
		}

		_ = eventMap.Close()
	}
}

func pkeys(keys []Key) {
	var tks []string
	for _, atk := range keys {
		tks = append(tks, atk.s)
	}
	s := strings.Join(tks, "")
	fmt.Printf("%s\n", s)

}
