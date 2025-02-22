package input

import (
	"flag"
)

type Config struct {
	ProducerCount int
	ConsumerCount int
	ImageLocation string
}

func NewConfig(progName string, args []string) (*Config, error) {
	fs := flag.NewFlagSet(progName, flag.ContinueOnError)

	var pc int
	fs.IntVar(&pc, "producer-count", 0, "the number of producers")

	var cc int
	fs.IntVar(&cc, "consumer-count", 0, "the number of consumers")

	var imageLocation string
	fs.StringVar(&imageLocation, "image-lcoation", "/images", "the file path to the images to process")

	if err := fs.Parse(args[1:]); err != nil {
		return nil, err
	}

	return &Config{
		ProducerCount: pc,
		ConsumerCount: cc,
		ImageLocation: imageLocation,
	}, nil
}
