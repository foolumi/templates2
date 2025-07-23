package main

import (
  "fmt"
  
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
  "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		config := config.New(ctx, "")
    phrase := fmt.Sprintf("Hello %s", config.Get("name"))
    fmt.Println(phrase)

		// Export the name of the bucket
		ctx.Export("phrase", pulumi.String(phrase))
		return nil
	})
}
