package translator

import (
	"fmt"
	"time"
)

// AllServiceUrls
// Returns all the services that are currently available.
func (a *translator) AllServiceUrls() []string {
	return defaultServiceUrls
}

// Latency
// Test the delay of each service node
func (a *translator) Latency(serviceUrls []string) {
	var fastestTime time.Duration
	var fastestUrl string
	for _, url := range serviceUrls {
		var total time.Duration
		var flag bool
		a.host = url
		for i := 0; i < 3; i++ {
			now := time.Now()
			text, err := a.translate("hello world!", "auto", "zh-ch", false)
			if err != nil || text == "" {
				fmt.Println(a.host, "failed")
				flag = false
				break
			} else {
				total += time.Since(now)
				flag = true
			}
		}
		if flag {
			avg := total / 3
			fmt.Println("Host: ", a.host, "Time average: ", avg)
			if fastestTime == 0 || avg < fastestTime {
				fastestTime = avg
				fastestUrl = url
			}
		}
	}
	fmt.Println("fastest:", fastestUrl, fastestTime)
}
