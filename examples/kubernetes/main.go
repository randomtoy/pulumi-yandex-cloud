package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/randomtoy/pulumi-yandex-cloud/modules/config"
	"github.com/randomtoy/pulumi-yandex-cloud/modules/serviceaccount"
)

type Data struct {
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.NewConfig(ctx)
		ya := conf.GetYandexData()
		sa := conf.GetServiceAccount("serviceaccount")
		fmt.Printf("Ya: %+v", sa)
		err := serviceaccount.CreateSA(ctx, ya, sa)
		if err != nil {
			return err
		}
		return nil
	})
}
