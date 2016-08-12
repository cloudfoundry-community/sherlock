package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func extractTime(time string) string {
	trimRight := strings.TrimRight(strings.TrimSpace(time), "]")
	return strings.TrimLeft(trimRight, "[")
}

func parseTime(timeString string) time.Time {
	time, err := time.Parse("2006-01-02 15:04:05.000", timeString)
	if err != nil {
		log.Fatal(err)
	}
	return time
}

func main() {
	logLocation := flag.String("log", "/var/vcap/sys/log/uaa/uaa.log", "location of log")
	timeInSeconds := flag.String("time", "60s", "how far back to look in logs in seconds")
	users := flag.String("users", "admin", "list of users seperated by commas to look for")
	flag.Parse()

	logFile, err := ioutil.ReadFile(*logLocation)
	if err != nil {
		log.Fatal(err)
	}

	duration, err := time.ParseDuration(*timeInSeconds)
	if err != nil {
		log.Fatal(err)
	}

	usersList := strings.Split(*users, ",")

	logAfter := time.Now().UTC().Add(-duration)
	splitLog := strings.Split(string(logFile), "\n")
	for i := len(splitLog) - 1; i > 0; i-- {
		if strings.HasPrefix(splitLog[i], "[") {
			timeSplit := strings.Split(splitLog[i], "uaa")
			parseTimeOfLog := parseTime(extractTime(timeSplit[0]))
			if !parseTimeOfLog.After(logAfter) {
				break
			}
			if strings.Contains(splitLog[i], "UserAuthentication") {
				for j := range usersList {
					if strings.Contains(splitLog[i], usersList[j]) {
						fmt.Println(splitLog[i])
						os.Exit(1)
					}
				}
			}
		}
	}
}
