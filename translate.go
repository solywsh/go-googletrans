package translator

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"strings"
	"time"
)

// Config basic config.
type Config struct {
	ServiceUrls []string
	UserAgent   []string
	Proxy       string
}

// translated result object.
type translated struct {
	Src    string // source language
	Dest   string // destination language
	Origin string // original text
	Text   string // translated text
}

type sentences struct {
	Sentences []sentence `json:"sentences"`
}

type sentence struct {
	Trans   string `json:"trans"`
	Orig    string `json:"orig"`
	Backend int    `json:"backend"`
}

type translator struct {
	host   string
	client *resty.Client
	ta     *tokenAcquirer
}

func randomChoose(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func New(config ...Config) *translator {
	rand.Seed(time.Now().Unix())
	var c Config
	if len(config) > 0 {
		c = config[0]
	}
	// set default value
	if len(c.ServiceUrls) == 0 {
		c.ServiceUrls = defaultServiceUrls
	}
	if len(c.UserAgent) == 0 {
		c.UserAgent = []string{defaultUserAgent}
	}
	host := randomChoose(c.ServiceUrls)
	userAgent := randomChoose(c.UserAgent)
	proxy := c.Proxy
	client := resty.New().SetHeaders(map[string]string{
		"User-Agent": userAgent,
	})
	// set proxy
	if strings.HasPrefix(proxy, "http") || strings.HasPrefix(proxy, "socks") {
		client.SetProxy(proxy)
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	ta := Token(host, client)
	return &translator{
		host:   host,
		client: client,
		ta:     ta,
	}
}

// Translate given content.
// Set src to `auto` and system will attempt to identify the source language automatically.
func (a *translator) Translate(origin, src, dest string) (*translated, error) {
	// check src & dest
	src = strings.ToLower(src)
	dest = strings.ToLower(dest)
	if _, ok := languages[src]; !ok {
		return nil, fmt.Errorf("src language code error")
	}
	if val, ok := languages[dest]; !ok || val == "auto" {
		return nil, fmt.Errorf("dest language code error")
	}
	text, err := a.translate(origin, src, dest, false)
	if err != nil || text == "" {
		text, err = a.translate(origin, src, dest, true)
	}
	if err != nil {
		return nil, err
	}
	result := &translated{
		Src:    src,
		Dest:   dest,
		Origin: origin,
		Text:   text,
	}
	return result, nil
}

func (a *translator) translate(origin, src, dest string, defaultUrl bool) (string, error) {
	var host string
	if defaultUrl {
		host = defaultServiceUrl
	} else {
		host = a.host
	}
	tk, err := a.ta.do(origin)
	if err != nil {
		return "", err
	}
	resp, err := a.client.R().SetQueryParams(map[string]string{
		"client": "gtx",
		"sl":     src,
		"tl":     dest,
		"hl":     dest,
		"tk":     tk,
		"q":      origin,
		"dt":     "t",
		// "dt": "bd",
		"dj":     "1",
		"source": "popup",
	}).Get(fmt.Sprintf("https://%s/translate_a/single", host))
	if err != nil {
		return "", err
	}
	if resp.StatusCode() == 200 {
		var sentences sentences
		err = json.Unmarshal(resp.Body(), &sentences)
		if err != nil {
			return "", err
		}
		translated := ""
		// parse trans
		for _, s := range sentences.Sentences {
			translated += s.Trans
		}
		return translated, nil
	} else {
		return "", fmt.Errorf("request error")
	}
}
