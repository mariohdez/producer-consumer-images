package processing

import (
	"context"
	"image"
	"sync"
)

type Consumer struct {
	waitGroup   *sync.WaitGroup
	imgCh       <-chan image.Image
	storagePath string
}

func NewConsumer(waitGroup *sync.WaitGroup, imgCh <-chan image.Image, storagePath string) *Consumer {
	return &Consumer{
		waitGroup:   waitGroup,
		imgCh:       imgCh,
		storagePath: storagePath,
	}
}

func (c *Consumer) Consume(ctx context.Context) {
	defer c.waitGroup.Done()
}
