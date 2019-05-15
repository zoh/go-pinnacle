# Go Pinnacle
Golang Wrapper for Pinnacle Sports API

[Pinnacle Documentation](https://www.pinnacle.com/en/api/manual)

# Installation

**go mod**
```
requer (
    github.com/zoh/go-pinnacle v1.0.0
)
```
or
```
go get github.com/zoh/go-pinnacle
```

# Usage

```go
package main

import "github.com/zoh/go-pinnacle"
import "log"

func main()  {
    client := go_pinnacle.NewApiClient(username, pass)
    res, _ := r.Account.GetClientBalance()
    log.Printf("%+v\n", res)
}
```