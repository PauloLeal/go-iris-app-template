package main

import (
	"flag"
	"getzoop/zec-companion-app/app"
	"os"
	"strconv"
)

func getDefaultPort() int {
	defaultEnvPort := 8080
	envPort := os.Getenv("SERVER_PORT")
	if len(envPort) > 0 {
		p, err := strconv.Atoi(envPort)
		if err == nil {
			defaultEnvPort = p
		}
	}

	return defaultEnvPort
}

func main() {
	var serverPort = flag.Int("p", getDefaultPort(), "web server http port")
	flag.Parse()

	app.Logger().Debug("START")

	_ = app.App().RunServer(*serverPort)
}
