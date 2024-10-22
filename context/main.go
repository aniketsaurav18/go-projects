package main

import (
	"context"
	"fmt"
	"time"
)

func SleepandTalk(ctx context.Context, t time.Duration, msg string) {
	select {
	case <-time.After(t):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println("context cancel executed")
		fmt.Println(ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	ctx , cancle := context.WithTimeout(ctx, 1000)
	defer cancle()
	SleepandTalk(ctx, 2*time.Second, "hello")
}
