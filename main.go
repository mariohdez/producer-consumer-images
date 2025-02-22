package main

import (
	"fmt"
	"log/slog"
	"os"
	"playground/image/internal/input"
)

func main() {
	err := ProcessImage(os.Args)
	if err != nil {
		slog.Error("Error processing image", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func ProcessImage(args []string) error {
	config, err := input.NewConfig(os.Args[0], os.Args[1:])
	if err != nil {
		return fmt.Errorf("read command line arguments: %w", err)
	}

	fmt.Print(config.ConsumerCount)
	return nil
}
