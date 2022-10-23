package translator

import (
	"fmt"
	"testing"
)

// TestTranslator_Translate calls translate.translate.
func TestTranslator_Translate(t *testing.T) {
	origin := "你好，世界！"
	c := Config{
		Proxy:       "socks5://10.10.40.10:30001",
		UserAgent:   []string{"Custom Agent"},
		ServiceUrls: []string{"translate.google.com.hk"},
	}
	trans := New(c)
	result, err := trans.Translate(origin, "auto", "en")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result.Text)
}
