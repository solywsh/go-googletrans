package translator

import (
	"fmt"
	"time"
)

// AllServiceUrls
// Returns all the services that are currently available.
func AllServiceUrls() []string {
	return serviceUrls
}

// Latency
// Test the delay of each service node
func (t *Translator) Latency(serviceUrls []string) {
	var fastestTime time.Duration
	var fastestUrl string
	for _, url := range serviceUrls {
		var total time.Duration
		var flag bool
		t.serviceUrl = url
		t.ta = Token(url, t.client)

		for i := 0; i < 3; i++ {
			now := time.Now()
			text, err := t.translate("hello world!", "auto", "zh-ch")
			if err != nil || text == "" {
				fmt.Println(t.serviceUrl, "failed")
				flag = false
				break
			} else {
				total += time.Since(now)
				flag = true
			}
		}
		if flag {
			avg := total / 3
			fmt.Println("Host: ", t.serviceUrl, "Time average: ", avg)
			if fastestTime == 0 || avg < fastestTime {
				fastestTime = avg
				fastestUrl = url
			}
		}
	}
	fmt.Println("fastest:", fastestUrl, fastestTime)
}
