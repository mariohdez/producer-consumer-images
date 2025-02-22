package processing

import (
	"context"
	"image"
	"image/color"
	"log/slog"
	"math"
	"playground/image/internal/random"
	"sync"
)

const randRangeBound = math.MaxUint16 + 1

type Producer struct {
	randGenerator random.Generator
	waitGroup     *sync.WaitGroup
	imgCh         chan<- image.Image
	maxWidth      int
	maxHeight     int
}

func New(randGenerator random.Generator, waitGroup *sync.WaitGroup, imgCh chan<- image.Image, maxWidth, maxHeight int) *Producer {
	return &Producer{
		randGenerator: randGenerator,
		waitGroup:     waitGroup,
		imgCh:         imgCh,
		maxWidth:      maxWidth,
		maxHeight:     maxHeight,
	}
}

func (p *Producer) Produce(ctx context.Context, wn int) {
	defer p.waitGroup.Done()

	for {
		img := p.createImage()
		select {
		case <-ctx.Done():
			slog.Info("stopping production due to cancellation", "workerNum", wn)
			return
		case p.imgCh <- img:
			slog.Info("produced image", "workerNum", wn)
		}
	}
}

func (p *Producer) createImage() *image.RGBA {
	img := image.NewRGBA(
		image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: p.maxWidth, Y: p.maxHeight},
		},
	)
	for x := 0; x < p.maxWidth; x++ {
		for y := 0; y < p.maxHeight; y++ {
			img.SetRGBA(x, y, color.RGBA{
				R: uint8(p.randGenerator.Generate(256)),
				G: uint8(p.randGenerator.Generate(256)),
				B: uint8(p.randGenerator.Generate(256)),
				A: uint8(255),
			})
		}
	}
	return img
}
