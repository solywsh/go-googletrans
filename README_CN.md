简体中文 | [English](./README.md)

# Googletrans

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
package main

import (
	"fmt"
	translator "github.com/solywsh/go-googletrans"
)

func main() {
	t := translator.New(
		translator.WithProxy("http://127.0.0.1:7890"),
	)
	result, err := t.Translate(`你好，世界！`, "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)
}
```

### 测试各节点响应时间

```go
package main

import (
	"fmt"
	translator "github.com/solywsh/go-googletrans"
)

func main() {
	t := translator.New()
	t.Latency(t.AllServiceUrls())
}

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

### 使用自定义服务 urls 或者 user agent

```go
package main

import (
	"fmt"
	translator "github.com/solywsh/go-googletrans"
)

func main() {
	t := translator.New(
		translator.WithProxy("http://127.0.0.1:7890"),
		translator.WithUserAgent("Custom Agent"),
		translator.WithServiceUrl("translate.google.com.hk"),
	)
	result, err := t.Translate(`你好，世界！`, "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)
}
```

完整示例： [Examples](./examples)

## Special thanks

* [Conight/go-googletrans](https://github.com/Conight/go-googletrans)
* [py-googletrans](https://github.com/ssut/py-googletrans)

## 协议

This SDK is distributed under the [The MIT License](https://opensource.org/licenses/MIT), see [LICENSE](./LICENSE) for more information.