package main

import (
	"context"
	"fmt"
	"time"

	"github.com/KRR19/retry/retry"
)

func createAction() retry.RetryAction {
	c := 1
	return func(ctx context.Context) error {
		if c == 5 {
			fmt.Println("200 OK")
			return nil
		}

		c++
		fmt.Println("408 Request Timeout")
		return fmt.Errorf("408 Request Timeout")
	}
}

func main() {
	a := createAction()
	r := retry.New(5, 1000*time.Millisecond)
	ctx := context.Background()
	r.Execute(ctx, a)
}
