package main

import "github.com/acore2026/nas/internal/tools/generator"

func main() {
	generator.ParseSpecs()

	generator.GenerateNasMessage()

	generator.GenerateNasEncDec()

	generator.GenerateTestLarge()
}
