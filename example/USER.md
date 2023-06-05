## Get User
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
	user, err := api.GetUser(context.Background(), "YOUR_USER_ID")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.Name)
}
```

You should replace the <YOUR_TOKEN> with the user token which created in pagerduty, and replace the <YOUR_USER_ID> with the user id of who you want to manipulate.

## Create User

```go
package main

import (
	"context"
	"fmt"
	pagerduty_go "github.com/faith2333/pagerduty-go"
	"github.com/faith2333/pagerduty-go/types"
	"log"
)

func main() {
	api := pagerduty_go.NewPagerDuty("YOUR_TOKEN")
	user, err := api.CreateUser(context.Background(), &types.CreateAndUpdateUserPayload{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.Name)
}
```

You should replace the <YOUR_TOKEN> with the user token which created in pagerduty,

and filled the Create Payload


## Update User

```go
package main

import (
	"context"
	"fmt"
	pagerduty_go "github.com/faith2333/pagerduty-go"
	"github.com/faith2333/pagerduty-go/types"
	"log"
)

func main() {
	api := pagerduty_go.NewPagerDuty("YOUR_TOKEN")
	user, err := api.UpdateUser(context.Background(), "YOUR_USER_ID", &types.CreateAndUpdateUserPayload{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.Name)
}
```

You should replace the <YOUR_TOKEN> with the user token which created in pagerduty, and replace the <YOUR_USER_ID> with the user id of who you want to manipulate,

and filled the Update Payload


