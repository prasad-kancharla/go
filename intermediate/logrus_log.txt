package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()

	log.Info("This is an info log message")
	log.Warn("This is a warn log message")
	log.Error("This is an error log message")

	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	log.WithFields(logrus.Fields{
		"name":   "John Doe",
		"method": "GET",
	}).Info("User logged in")
}
