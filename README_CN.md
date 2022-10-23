# Googletrans

Language: 简体中文 | [English](./README.md)

[![Sourcegraph](https://sourcegraph.com/github.com/solywsh/go-googletrans/-/badge.svg)](https://sourcegraph.com/github.com/Conight/go-googletrans?badge)[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Conight/go-googletrans/blob/master/LICENSE)

> 改造自[Conight/go-googletrans](https://github.com/Conight/go-googletrans)项目，新增支持socks5代理，更换为resty客户端

## 从Github下载

```shell script
GO111MODULE=on go get github.com/solywsh/go-googletrans
```

## 快速开始

### 简单翻译

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

### 使用代理

```go
c := translator.Config{
    Proxy: "http://PROXY_HOST:PROXY_PORT",
    // Proxy: "socks5://PROXY_HOST:PROXY_PORT"
}
t := translate.New(c)
```

### 使用自定义服务 urls 或者 user agent

```go
c := translator.Config{
    UserAgent: []string{"Custom Agent"},
    ServiceUrls: []string{"translate.google.com.hk"},
}
t := translate.New(c)
```

完整示例： [Examples](./examples)

## Special thanks

* [Conight/go-googletrans](https://github.com/Conight/go-googletrans)
* [py-googletrans](https://github.com/ssut/py-googletrans)

## 协议

This SDK is distributed under the [The MIT License](https://opensource.org/licenses/MIT), see [LICENSE](./LICENSE) for more information.