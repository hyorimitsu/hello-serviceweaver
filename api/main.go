package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ServiceWeaver/weaver"

	"github.com/hyorimitsu/hello-serviceweaver/api/notificationservice"
	"github.com/hyorimitsu/hello-serviceweaver/api/paymentservice"
	"github.com/hyorimitsu/hello-serviceweaver/api/productservice"
	"github.com/hyorimitsu/hello-serviceweaver/api/server"
)

//go:generate weaver generate ./...

func main() {
	root := weaver.Init(context.Background())

	notification, err := weaver.Get[notificationservice.T](root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	payment, err := weaver.Get[paymentservice.T](root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	product, err := weaver.Get[productservice.T](root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	srv := server.NewServer(notification, payment, product)
	if err := srv.Run(root); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
