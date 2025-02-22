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
		slog.Error("processing image", "error", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func ProcessImage(args []string) error {
	programName := os.Args[0]
	argWithoutProgramName := os.Args[1:]
	config, err := input.NewConfig(programName, argWithoutProgramName)
	if err != nil {
		return fmt.Errorf("read command line arguments: %w", err)
	}

	fmt.Print(config.ConsumerCount)
	return nil
}
