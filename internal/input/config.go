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

	config := Config{}
	fs.IntVar(&config.ProducerCount, "producer-count", 0, "the number of producers")
	fs.IntVar(&config.ConsumerCount, "consumer-count", 0, "the number of consumers")
	fs.StringVar(&config.ImageLocation, "image-location", "/images", "the file path to the images to process")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	return &config, nil
}
