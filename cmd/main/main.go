package main

import (
	"log"
	"os"
	"maria/init"
	"maria/internal/pkg/cli"
)

func main() {
	sysInit.Init()

	err := cli.InitCli().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
