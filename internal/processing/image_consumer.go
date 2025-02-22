package processing

import (
	"context"
	"image"
	"image/color"
	"image/png"
	"log/slog"
	"os"
	"playground/image/internal/filename"
	"sync"
)

type ConsumerPool struct {
	fileNameGenerator filename.Generator
	waitGroup         *sync.WaitGroup
	imgCh             <-chan image.Image
	storagePath       string
}

func NewConsumerPool(fnGenerator filename.Generator, waitGroup *sync.WaitGroup, imgCh <-chan image.Image, storagePath string) *ConsumerPool {
	return &ConsumerPool{
		fileNameGenerator: fnGenerator,
		waitGroup:         waitGroup,
		imgCh:             imgCh,
		storagePath:       storagePath,
	}
}

func (cp *ConsumerPool) Consume(ctx context.Context) {
	defer cp.waitGroup.Done()

	for {
		select {
		case <-ctx.Done():
			slog.Info("shutting consumer down because received cancellation")
			return
		case img, ok := <-cp.imgCh:
			if !ok {
				slog.Info("shutting consumer down because image channel closed")
				return
			}

			cp.convertImgToGreyscale(img)
			cp.saveImg(img)
		}
	}
}

func (cp *ConsumerPool) convertImgToGreyscale(img image.Image) {
	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		slog.Error("convert image.Image to image.RGBA")
		return
	}

	for x := 0; x < rgbaImg.Bounds().Max.X; x++ {
		for y := 0; y < rgbaImg.Bounds().Max.Y; y++ {
			// Found this algo for grayscale conversion https://en.wikipedia.org/wiki/Rec._709
			r, g, b, _ := rgbaImg.At(x, y).RGBA()
			gs := uint8(0.2126*float64(r>>8) + 0.7152*float64(g>>8) + 0.0722*float64(b>>8))
			rgbaImg.SetRGBA(x, y, color.RGBA{R: gs, G: gs, B: gs, A: 255})
		}
	}
}

func (cp *ConsumerPool) saveImg(img image.Image) {
	f, err := os.Create(cp.storagePath + cp.fileNameGenerator.Generate())
	if err != nil {
		// TODO: return error.=
	}
	defer func() {
		_ = f.Close()
	}()

	if err := png.Encode(f, img); err != nil {
		// TODO return err
		// slog.Error("image encoding", "error", err)
		// os.Exit(1)
	}
}
