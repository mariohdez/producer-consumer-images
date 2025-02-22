package main

import (
	"context"
	"errors"
	"fmt"
	"image"
	"log/slog"
	"os"
	"playground/image/internal/filename"
	"playground/image/internal/input"
	"playground/image/internal/processing"
	"playground/image/internal/random"
	"sync"
	"time"
)

func main() {
	err := ProcessImages(context.Background(), os.Args, &random.RealGenerator{})
	if err != nil {
		slog.Error("processing image", "error", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func ProcessImages(ctx context.Context, args []string, rg random.Generator) error {
	if len(args) < 1 {
		return errors.New("command line arguments length")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	cfg, err := input.NewConfig(os.Args[0], os.Args[1:])
	if err != nil {
		return fmt.Errorf("read command line arguments: %w", err)
	}

	imgCh := make(chan image.Image, 5)
	var producerWG sync.WaitGroup
	ip := processing.New(rg, &producerWG, imgCh, 1000, 1000)
	for i := 0; i < cfg.ProducerCount; i++ {
		producerWG.Add(1)
		go func(wn int) {
			ip.Produce(ctx, wn)
		}(i)
	}
	go func() {
		producerWG.Wait()
		close(imgCh)
	}()

	var consumerWG sync.WaitGroup
	var fnGenerator = &filename.RealGenerator{}
	ic := processing.NewConsumerPool(fnGenerator, &consumerWG, imgCh, cfg.ImageLocation)
	for i := 0; i < cfg.ConsumerCount; i++ {
		consumerWG.Add(1)
		go ic.Consume(ctx)
	}
	consumerWG.Wait()

	return nil
}
