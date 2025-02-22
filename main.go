package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"playground/image/internal/input"
	"playground/image/internal/processing"
	"sync"
	"time"
)

func main() {
	err := ProcessImages(context.Background(), os.Args)
	if err != nil {
		slog.Error("processing image", "error", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func ProcessImages(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return errors.New("command line arguments length")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	cfg, err := input.NewConfig(os.Args[0], os.Args[1:])
	if err != nil {
		return fmt.Errorf("read command line arguments: %w", err)
	}

	imgCh := make(chan int, 10)
	var wg sync.WaitGroup
	ip := processing.NewProducerWorker(cfg.ProducerCount, &wg, imgCh)
	for i := 0; i < cfg.ProducerCount; i++ {
		wg.Add(1)
		go func(wn int) {
			ip.Produce(ctx, wn)
		}(i)
	}

	wg.Add(1)
	go Consumer(ctx, &wg, imgCh)

	wg.Wait()

	close(imgCh)

	return nil
}

func Consumer(ctx context.Context, wg *sync.WaitGroup, imgCh <-chan int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			slog.Info("stopping consumption due to cancellation")
			return
		case item, ok := <-imgCh:
			if !ok {
				fmt.Printf("ch closed...")
				return
			}
			slog.Info("read item", "item", item)
			time.Sleep(time.Millisecond * 100)
		}
	}
}
