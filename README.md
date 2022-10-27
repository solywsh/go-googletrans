English | [简体中文](./README_CN.md)

# Googletrans

[![Sourcegraph](https://sourcegraph.com/github.com/solywsh/go-googletrans/-/badge.svg)](https://sourcegraph.com/github.com/Conight/go-googletrans?badge)[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Conight/go-googletrans/blob/master/LICENSE)

> Transformation of [Conight/go-googletrans](https://github.com/Conight/go-googletrans) project, support socks5/http proxy and replace it with resty client

## Download from Github

```shell script
GO111MODULE=on go get github.com/solywsh/go-googletrans
```

## Quick Start Example

### Simple translate

```go
package main

import (
	"fmt"
	"github.com/solywsh/go-googletrans"
)

func main() {
	t := translator.New()
	result, err := t.Translate("你好，世界！", "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)
}
```

### Using proxy

```go
c := translator.Config{
    Proxy: "http://PROXY_HOST:PROXY_PORT",
    // Proxy: "socks5://PROXY_HOST:PROXY_PORT"
}
t := translate.New(c)
```

### Test the response time of each service

```go
c := Config{
		Proxy: "socks5://127.0.0.1:7890",
}
trans := New(c)
trans.Latency(trans.AllServiceUrls())

// Output
Host:  translate.google.ac Time average:  896.617266ms
Host:  translate.google.ad Time average:  787.691933ms
Host:  translate.google.ae Time average:  749.772166ms
Host:  translate.google.al Time average:  712.4626ms
Host:  translate.google.am Time average:  796.3925ms
Host:  translate.google.as Time average:  750.585733ms
Host:  translate.google.at Time average:  847.926733ms
...
Host:  translate.google.to Time average:  648.1111ms
Host:  translate.google.tt Time average:  583.769633ms
Host:  translate.google.us Time average:  540.138666ms
Host:  translate.google.vg Time average:  575.703566ms
Host:  translate.google.vu Time average:  673.641833ms
Host:  translate.google.ws Time average:  563.181933ms
fastest: translate.google.com.ec 466.401633ms
```

### Using custom service urls or user agent

```go
c := translator.Config{
    UserAgent: []string{"Custom Agent"},
    ServiceUrls: []string{"translate.google.com.hk"},
}
t := translate.New(c)
```

See [Examples](./examples) for more examples.

## Special thanks

* [Conight/go-googletrans](https://github.com/Conight/go-googletrans)
* [py-googletrans](https://github.com/ssut/py-googletrans)

## License

This SDK is distributed under the [The MIT License](https://opensource.org/licenses/MIT), see [LICENSE](./LICENSE) for more information.
