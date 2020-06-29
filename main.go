package main

import (
	"flag"
	"github.com/PauloLeal/go-iris-app-template/app"
	"github.com/sirupsen/logrus"
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

func initLogger(logLevel logrus.Level) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logLevel)
}

func main() {
	var serverPort = flag.Int("p", getDefaultPort(), "web server http port")
	var logAll = flag.Bool("log-all", false, "Log all messages (trace level)")

	flag.Parse()

	logLevel := logrus.InfoLevel
	if *logAll {
		logLevel = logrus.TraceLevel
	}

	initLogger(logLevel)

	logrus.Debug("START")

	_ = app.App().RunServer(*serverPort)
}
