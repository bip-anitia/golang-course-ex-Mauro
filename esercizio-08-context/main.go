package main

import (
	"context"
	"fmt"
	"time"
)

func workerPoolWithContextExample() { /* workers che ascoltano ctx.Done() */ }
func withValueExample()             { /* key type-safe + helper getter */ }

func main() {
	withTimeoutExample()
	withCancellationExample()
	workerPoolWithContextExample()
	pipelineWithContextExample()
	withValueExample()
}

func withTimeoutExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	resultCh := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second) // lavoro lento
		resultCh <- "completed"
	}()

	select {
	case result := <-resultCh:
		fmt.Println("timeout example:", result)
	case <-ctx.Done():
		fmt.Println("timeout example:", ctx.Err())
	}
}

func withCancellationExample() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(700 * time.Millisecond)
		fmt.Println("cancellation example: cancelling now")
		cancel()
	}()

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("cancellation example: working...")
		case <-ctx.Done():
			fmt.Println("cancellation example:", ctx.Err())
			return
		}
	}
}

func pipelineWithContextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	input := []int{1, 2, 3, 4, 5, 6, 7, 8}

	out := pipeline(ctx, input)

	for value := range out {
		fmt.Println("pipeline example:", value)
		time.Sleep(150 * time.Millisecond) // simula consumer lento
	}
	fmt.Println("pipeline example done:", ctx.Err())
}

func pipeline(ctx context.Context, nums []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, number := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- number * 2:
			}
		}
	}()

	return out
}
