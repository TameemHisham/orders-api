package main

import (
	"context"
	"fmt"

	"github.com/TameemHisham/orders-api/application"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("Failed to start app: %v\n", err)
	}
}