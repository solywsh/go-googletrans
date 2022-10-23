package translator

import (
	"fmt"
	"testing"
)

// TestTranslator_Translate calls translate.translate.
func TestTranslator_Translate(t *testing.T) {
	origin := "你好，世界！"
	c := Config{
		Proxy:       "socks5://127.0.0.1:7890",
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
