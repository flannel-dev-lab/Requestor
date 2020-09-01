# Requestor

[![Build Status](https://github.com/flannel-dev-lab/Requestor/workflows/Requestor/badge.svg)](https://github.com/flannel-dev-lab/Requestor/workflows/Requestor/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/flannel-dev-lab/Requestor)](https://goreportcard.com/report/github.com/flannel-dev-lab/Requestor)
[![codecov](https://codecov.io/gh/flannel-dev-lab/Requestor/branch/master/graph/badge.svg)](https://codecov.io/gh/flannel-dev-lab/Requestor)
![GitHub repo size](https://img.shields.io/github/repo-size/flannel-dev-lab/Requestor)
![GitHub](https://img.shields.io/github/license/flannel-dev-lab/Requestor)
[![GoDoc](https://godoc.org/github.com/flannel-dev-lab/Requestor?status.svg)](https://pkg.go.dev/github.com/flannel-dev-lab/Requestor?tab=doc)

![GitHub Logo1](Gopher.jpg)

Requestor is a simple HTTP library for Developers. The main idea behind this project is to make HTTP requests fun, simple
and easy (Inspired by python requests). 

## Features
- All the HTTP methods are supported
- Proxy Support
- Retry Support - You can set the max number of times you want to retry if the request fails
- TLS Client Certificates support
- Enable/Disable Keep-Alive
- Timeouts
- Built purely using the standard library
- more coming soon

## Installation
```shell script
go get github.com/flannel-dev-lab/Requestor
```

## Usage
```go
package main

import (
    "fmt"
    "github.com/flannel-dev-lab/Requestor"
    "log"
)

func main() {
    client := requestor.New()
    headers := map[string][]string{
    	"Content-Type": {"application/json"},
    }
    
    queryParams := map[string][]string{
    	"arg1": {"test"},
    }

	response, err := client.Get("http://httpbin.org/get", headers, queryParams)
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("Do something with ", response)
}
```

The usage is similar for `POST` method as well, the only difference is that you send can send data in the request.
For JSON requests, the data can be either a struct or map, by default Requestor assumes Content-Type to be `application/json`

For `application/x-www-form-urlencoded` requests, make sure the data is in the form `map[string][]string`

**Note**: Requestor does not support XML

```go
package main

import (
    "fmt"
    "github.com/flannel-dev-lab/Requestor"
    "log"
)

func main() {
    client := requestor.New()
    headers := map[string][]string{
    	"Content-Type": {"application/json"},
    }
    
    queryParams := map[string][]string{
    	"arg1": {"test"},
    }

	response, err := client.Post("http://httpbin.org/get", headers, queryParams, map[string]string{"hello": "world"})
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("Do something with ", response)
}
```

## Advanced Usage

### Setting Timeouts
```go
client := requestor.New()
client.SetTimeout(10)
```

## Using Proxy
```
client := requestor.New()
client.SetHTTPProxy(proxyURL, username, password string)
```
or
```
client := requestor.New()
client.SetHTTPSProxy(proxyURL, username, password string)
```

### Disabling Keep-Alive
```
client := requestor.New()
client.DisableKeepAlive(true)
```

### Much-more settings can be found here [![GoDoc](https://godoc.org/github.com/flannel-dev-lab/Requestor?status.svg)](https://pkg.go.dev/github.com/flannel-dev-lab/Requestor?tab=doc)


## Contributing
Requestor loves contributions. If you find a bug, want to add a feature, you can create a PR and we will take a look.