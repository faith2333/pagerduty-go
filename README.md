# The Go SDK for [PagerDuty](https://applovin-com.pagerduty.com/)

## Synopsis

This is a pagerduty library For Go created by [Annan Wang](https://github.com/faith23333).
It helps you manipulate pagerduty via go easily. includes:
* CRUD for pagerduty resource object
* Incident operation 
* Provides a structured resource object definition.
* Provides a RESTClient.

## Reference
[Pagerduty API](https://developer.pagerduty.com/api-reference/e65c5833eeb07-pager-duty-api)

## Project Status

This project was just started, and will be finished as soon as possible.

Any contribution is welcomed.

If you have any question, you can make an issue or send email to me directly (wan199406@gmail.com)

## Installing

### go get
```shell
go get github.com/faith2333/pagerduty-go
```

## Example

### Start
```go
package main

import (
	"context"
	"fmt"
	pagerduty_go "github.com/faith2333/pagerduty-go"
	"log"
)

func main() {
	api := pagerduty_go.NewPagerDuty("YOUR_TOKEN")
	fmt.Println("new api client",api)
}
```

You should replace the <YOUR_TOKEN> with the user token which created in pagerduty.

### More Example
* [USER](./example/USER.md)