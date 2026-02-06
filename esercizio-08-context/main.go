package main

import (
	"context"
	"fmt"
	"time"
)

func withCancellationExample()      { /* context.WithCancel + ticker */ }
func workerPoolWithContextExample() { /* workers che ascoltano ctx.Done() */ }
func pipelineWithContextExample()   { /* stage che esce su ctx.Done() */ }
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
