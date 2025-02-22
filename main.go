package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"playground/image/internal/input"
)

func main() {
	err := ProcessImage(os.Args)
	if err != nil {
		slog.Error("processing image", "error", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func ProcessImage(args []string) error {
	if len(args) < 1 {
		return errors.New("command line arguments length")
	}

	_, err := input.NewConfig(os.Args[0], os.Args[1:])
	if err != nil {
		return fmt.Errorf("read command line arguments: %w", err)
	}

	return nil
}
