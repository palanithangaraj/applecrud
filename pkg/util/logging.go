package util

import (
	log "github.com/sirupsen/logrus"
)

func SetLoggingFormat() {
	Formatter := new(log.TextFormatter)
	// You can change the Timestamp format. But you have to use the same date and tim    e.
	// "2006-02-02 15:04:06" Works. If you change any digit, it won't work
	// ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
	Formatter.TimestampFormat = "2006-02-01 15:04:05"
	//Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
}

//TODO - improve timing where the logging level is set?

//TODO - set logging level here

//TODO - set logging output file here

//TODO - set other options e.g. colors

//TODO -- add unit tests
