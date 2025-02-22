package processing

import (
	"context"
	"image"
	"image/color"
	"log/slog"
	"math"
	"math/rand"
	"sync"
)

const randRangeBound = math.MaxUint16 + 1

type Producer struct {
	waitGroup *sync.WaitGroup
	imgCh     chan<- image.Image
	maxWidth  int
	maxHeight int
}

func New(waitGroup *sync.WaitGroup, imgCh chan<- image.Image, maxWidth, maxHeight int) *Producer {
	return &Producer{
		waitGroup: waitGroup,
		imgCh:     imgCh,
		maxWidth:  maxWidth,
		maxHeight: maxHeight,
	}
}

func (pw *Producer) Produce(ctx context.Context, wn int) {
	defer pw.waitGroup.Done()

	for {
		img := pw.createImage()
		select {
		case <-ctx.Done():
			slog.Info("stopping production due to cancellation", "workerNum", wn)
			return
		case pw.imgCh <- img:
			slog.Info("wrote val to ch", "val", img)
		}
	}
}

func (pw *Producer) createImage() *image.RGBA64 {
	img := image.NewRGBA64(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: pw.maxWidth, Y: pw.maxHeight},
	})
	for x := 0; x < pw.maxWidth; x++ {
		for y := 0; y < pw.maxHeight; y++ {
			img.SetRGBA64(x, y, color.RGBA64{
				R: uint16(rand.Int31n(randRangeBound)),
				G: uint16(rand.Int31n(randRangeBound)),
				B: uint16(rand.Int31n(randRangeBound)),
			})
		}
	}
	return img
}
