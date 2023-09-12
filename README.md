### simple smtp

```go
package main

import (
	"github.com/fanchann/simplesmtp"
)

func main() {
	s := simplesmtp.SimpleSmtp{
		Email:    "youremailhere@xyz.com",
		Password: "XXXXXXXXXXXX",
		To:       []string{"yourfriend1@gmail.com", "yourfriend2@gmail.com"},
		Subject:  "Hello Friend!",
		Body:     "Pinjem dulu seratus",
	}

	s.Send()
} 
```

more simple
```go
package main

import (
	"github.com/fanchann/simplesmtp"
)

func main(){
	simplesmtp.NewSimpleSmtp("youremail@xyz.com", "XXXXX", "smtoHostHere", 666, []string{"yourfriend1@gmail.com", "yourfriend2@gmail.com"}, "your subject", "your message").Send()

    //666 is port
}
```