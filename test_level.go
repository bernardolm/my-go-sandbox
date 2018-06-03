package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func logrusLevelToSylog(level logrus.Level) int32 {
	// Till warn, logrus levels are lower than syslog by 1
	// (logrus has no equivalent of syslog LOG_NOTICE)
	if level <= logrus.WarnLevel {
		fmt.Printf("%d, increse 1\n", level)
		return int32(level) + 1
	}
	// From info, logrus levels are lower than syslog by 2
	fmt.Printf("%d, increse 2\n", level)
	return int32(level) + 2
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	fmt.Printf("Till warn, logrus levels are lower than syslog by 1\n\n")

	logrus.Debugf("DebugLevel (in logrus %d) output to graylog as %d\n",
		logrus.DebugLevel,
		logrusLevelToSylog(logrus.DebugLevel))

	logrus.Infof("InfoLevel (in logrus %d) output to graylog as %d\n",
		logrus.InfoLevel,
		logrusLevelToSylog(logrus.InfoLevel))

	logrus.Warnf("WarnLevel (in logrus %d) output to graylog as %d\n",
		logrus.WarnLevel,
		logrusLevelToSylog(logrus.WarnLevel))

	logrus.Errorf("ErrorLevel (in logrus %d) output to graylog as %d\n",
		logrus.ErrorLevel,
		logrusLevelToSylog(logrus.ErrorLevel))

	logrus.Errorf("FatalLevel (in logrus %d) output to graylog as %d\n",
		logrus.FatalLevel,
		logrusLevelToSylog(logrus.FatalLevel))

	logrus.Errorf("PanicLevel (in logrus %d) output to graylog as %d\n",
		logrus.PanicLevel,
		logrusLevelToSylog(logrus.PanicLevel))
}
