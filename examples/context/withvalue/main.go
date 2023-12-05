package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	if ctx.Value("authorization") != "my-auth-token" {
		fmt.Println("Unauthorized from doSomething")
		return
	}
	fmt.Println("Authorized from doSomething")
}

func doAnother(ctx context.Context) {
	if ctx.Value("authorization") != "my-auth-token" {
		fmt.Println("Unauthorized from doAnother")
		return
	}
	fmt.Println("Authorized from doAnother")
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "authorization", "my-auth-token")
	doSomething(ctx)
	doAnother(ctx)
}
