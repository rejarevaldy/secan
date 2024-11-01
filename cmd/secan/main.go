package main

import "github.com/rejarevaldy/secan/internal/runner"

func main() {
	options := runner.ParseOptions()
	runner.New(options)
}
