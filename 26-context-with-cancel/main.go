package main

import (
	"context"
	"fmt"
)

type userIDKey string
type database map[string]bool

var db database = database{
	"abhinand": true,
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	processRequest(ctx, "abhinand")
}

func processRequest(ctx context.Context, userID string) {

	vctx := context.WithValue(ctx, userIDKey("userIDKey"), userID)
	ch := checkMemberShipStatus(vctx)
	status := <-ch
	fmt.Printf("memebership status of userid : %s : %v\n", userID, status)
}

func checkMemberShipStatus(ctx context.Context) <-chan bool {

	ch := make(chan bool)
	go func() {
		defer close(ch)
		userID := ctx.Value(userIDKey("userIDKey")).(string)
		status := db[userID]
		ch <- status
	}()
	return ch
}
