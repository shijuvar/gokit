package main

import (
	"context"
	"fmt"
)

type role struct {
	user string
	role string
}

func doSomething(ctx context.Context) {
	if ctx.Value("authorization") != "my-auth-token" {
		fmt.Println("Unauthorized from doSomething")
		return
	}
	fmt.Println("Authorized from doSomething for user:", ctx.Value("user"))
}

func doAnother(ctx context.Context) {
	if ctx.Value("authorization") != "my-auth-token" {
		fmt.Println("Unauthorized from doAnother")
		return
	}
	fmt.Println("Authorized from doAnother:", ctx.Value("user"))
}
func doAdmin(ctx context.Context) {
	if ctx.Value("authorization") == "my-auth-token" {
		fmt.Printf("Role:%+v", ctx.Value("role").(role))
	}
}
func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "authorization", "my-auth-token")
	ctx = context.WithValue(ctx, "user", "shijuvar")
	doSomething(ctx)
	doAnother(ctx)
	ctx = context.WithValue(ctx, "role", role{
		user: "shijuvar",
		role: "admin",
	})
	doAdmin(ctx)
}
