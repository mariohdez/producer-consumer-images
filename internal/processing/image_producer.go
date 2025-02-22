package processing

import (
	"context"
	"log/slog"
	"math/rand"
	"sync"
)

// accept interfaces return struct

type ProducerWorker struct {
	producerCount int
	waitGroup     *sync.WaitGroup
	imgCh         chan<- int
}

func NewProducerWorker(prodCnt int, wg *sync.WaitGroup, imgCh chan<- int) *ProducerWorker {
	return &ProducerWorker{
		producerCount: prodCnt,
		waitGroup:     wg,
		imgCh:         imgCh,
	}
}

func (pw *ProducerWorker) Produce(ctx context.Context, wn int) {
	defer pw.waitGroup.Done()

	for {
		val := rand.Intn(100)
		select {
		case <-ctx.Done():
			slog.Info("stopping production due to cancellation", "workerNum", wn)
			return
		case pw.imgCh <- val:
			slog.Info("wrote val to ch", "val", val)
		}
	}
}
