//go:build !cover

package main

import (
	"bufio"
	"context"
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/uponusolutions/go-smtp/mailer"
)

//go:embed testdata/example.eml
var eml string

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter recipients: ")
	recipientsRaw, _ := reader.ReadString('\n')
	recipients := strings.Split(strings.Trim(recipientsRaw, "\n\r "), ",")

	res, err := mailer.Send(context.Background(), recipients[0], recipients, strings.NewReader(eml))
	if err != nil {
		panic(err)
	}

	for _, fail := range res.Failures {
		fmt.Printf("Failed to send mail to %s: %s \n", fail.Rcpts, fail.Error.Error())
	}

	for _, res := range res.Responses {
		fmt.Printf("Success to send mail to %s: %d %s \n", res.Rcpts, res.Code, res.Msg)
	}
}
