package main

import "coinrunner/pkg/coinrunner"

func main() {

	cfg := coinrunner.GetConfig()
	coinrunner.InitializeGame(cfg)
}
