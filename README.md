Language: English | [简体中文](./README_CN.md)
# Googletrans

[![Sourcegraph](https://sourcegraph.com/github.com/solywsh/go-googletrans/-/badge.svg)](https://sourcegraph.com/github.com/Conight/go-googletrans?badge)[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Conight/go-googletrans/blob/master/LICENSE)

> Transformation of t[Conight/go-googletrans](https://github.com/Conight/go-googletrans) project, support socks5/http proxy and replace it with resty client

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
