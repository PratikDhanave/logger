package main

import (
	"fmt"
	"os"

	"github.com/PratikDhanave/logger/logging"
)

var build = "Production"

func main() {

	log, err := logging.New("API")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer log.Sync()

	log.Infow("starting service", "version", build)
	log.Infow("stopping service", "version", build)
}
