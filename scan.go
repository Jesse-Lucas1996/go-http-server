package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var ScannedIpAddresses []string

func ipscanner() {
	http.HandleFunc("/scan", func(w http.ResponseWriter, r *http.Request) {
		ipAddress := r.FormValue("ip-address")
		port := "25565"

		// Respond with a 204 No Content status
		w.WriteHeader(http.StatusNoContent)

		go func() {
			ipComponents := strings.Split(ipAddress, ".")

			var wg sync.WaitGroup

			for i := 0; i < 255; i++ {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()

					ipComponents[3] = strconv.Itoa(i)
					newIpAddress := strings.Join(ipComponents, ".")

					_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", newIpAddress, port), time.Second*1)
					if err != nil {
						log.Printf("%s: Port is closed or host is unreachable\n", newIpAddress)
					} else {
						ScannedIpAddresses = append(ScannedIpAddresses, newIpAddress)
						jsonData, err := json.Marshal(ScannedIpAddresses)
						check(err)
						err = os.WriteFile("./ip-found.json", jsonData, 0644)
						check(err)
						log.Printf("%s: Port is open\n", newIpAddress)
					}
				}(i)
			}

			wg.Wait()
		}()
	})
}
