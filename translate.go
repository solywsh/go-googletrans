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

// Translated result object.
type Translated struct {
	Src    string // source language
	Dest   string // destination language
	Origin string // original text
	Text   string // Translated text
}

type sentences struct {
	Sentences []sentence `json:"sentences"`
}

type sentence struct {
	Trans   string `json:"trans"`
	Orig    string `json:"orig"`
	Backend int    `json:"backend"`
}

type Translator struct {
	randomEverytime bool     // random service url every time
	randomServices  []string // random service url list
	serviceUrl      string
	client          *resty.Client
	ta              *TokenAcquirer
}

func randomChoose(slice []string) string {
	rand.Seed(time.Now().Unix())
	return slice[rand.Intn(len(slice))]
}

type Option func(*Translator)

func WithServiceUrl(serviceUrl string) Option {
	return func(t *Translator) {
		t.ta = Token(serviceUrl, t.client)
		t.serviceUrl = serviceUrl
	}
}

func WithRandomServiceUrl() Option {
	return func(t *Translator) {
		t.serviceUrl = randomChoose(serviceUrls)
		t.ta = Token(t.serviceUrl, t.client)
	}
}

func WithUserAgent(userAgent string) Option {
	return func(t *Translator) {
		t.client.SetHeaders(map[string]string{
			"User-Agent": userAgent,
		})
	}
}

func WithProxy(proxy string) Option {
	return func(translator *Translator) {
		translator.client.SetProxy(proxy)
		translator.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
}

func WithRetryTimes(times int) Option {
	return func(translator *Translator) {
		translator.client.SetRetryCount(times)
	}
}

func WithRandomServiceUrlEveryTime(serviceUrlList []string) Option {
	return func(t *Translator) {
		t.randomEverytime = true
		t.randomServices = serviceUrlList
	}
}

func New(options ...Option) *Translator {
	client := resty.New().
		SetHeaders(map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
		}).SetRetryCount(3)
	t := &Translator{
		serviceUrl: "translate.google.com",
		client:     client,
		ta:         Token("translate.google.com", client),
	}
	for _, option := range options {
		option(t)
	}
	return t
}

// Translate given content.
// Set src to `auto` and system will attempt to identify the source language automatically.
func (t *Translator) Translate(origin, src, dest string) (*Translated, error) {
	src = strings.ToLower(src)
	dest = strings.ToLower(dest)
	if _, ok := languages[src]; !ok {
		return nil, fmt.Errorf("src language code error")
	}
	if val, ok := languages[dest]; !ok || val == "auto" {
		return nil, fmt.Errorf("dest language code error")
	}
	if t.randomEverytime {
		t.serviceUrl = randomChoose(t.randomServices)
		t.ta = Token(t.serviceUrl, t.client)
	}
	text, err := t.translate(origin, src, dest)
	if err != nil {
		return nil, err
	}
	result := &Translated{
		Src:    src,
		Dest:   dest,
		Origin: origin,
		Text:   text,
	}
	return result, nil
}

func (t *Translator) translate(origin, src, dest string) (string, error) {
	tk, err := t.ta.do(origin)
	if err != nil {
		return "", err
	}
	resp, err := t.client.R().SetQueryParams(map[string]string{
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
	}).Get(fmt.Sprintf("https://%s/translate_a/single", t.serviceUrl))
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
