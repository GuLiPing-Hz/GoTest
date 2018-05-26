package main

import (
	"github.com/sirupsen/logrus"
	"fmt"
	"os"
)

type MyHook struct {
}

func (hook *MyHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel,
		logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
}

func (hook *MyHook) Fire(entry *logrus.Entry) error {
	//fmt.Println(entry.Time.Date())
	//go time Format 函数是依据2006年的这一天开始算起
	fmt.Printf("%s %s %s %s\n",entry.Time.Format("2006-01-02 15:04:05.000"),entry.Level,entry.Message,entry.Data)
	return nil
}

func main() {

	// Log as JSON instead of the default ASCII formatter.
	//logrus.SetFormatter(&logrus.JSONFormatter{})

	var log = logrus.New()

	log.Hooks.Add(new(MyHook))

	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	//log.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("robot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
	 log.Out = file
	} else {
	 log.Info("Failed to log to file, using default stderr")
	}

	// do something here to set environment depending on an environment variable
	// or command-line flag
	Environment := "debug"
	if Environment == "production" {
		log.Formatter = &logrus.JSONFormatter{}
	} else {
		log.Formatter = &logrus.TextFormatter{}
	}

	logWithFields := log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	})
	logWithFields.Info("A group of walrus emerges from the ocean")
	log.SetLevel(logrus.DebugLevel)

	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
}
