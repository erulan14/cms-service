package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

type writeHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writeHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
	}

	return err
}

func (hook *writeHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithField(key string, value interface{}) Logger {
	return Logger{l.WithField(key, value)}
}

func Init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0755)

	if err != nil || os.IsNotExist(err) {
		panic("Failed to create log directory")
	} else {
		errorFile, err := os.OpenFile("logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("Failed to open log file")
		}

		warningFile, err := os.OpenFile("logs/warning.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("Failed to open log file")
		}

		infoFile, err := os.OpenFile("logs/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("Failed to open log file")
		}

		l.SetOutput(ioutil.Discard)

		l.AddHook(&writeHook{
			Writer:    []io.Writer{errorFile},
			LogLevels: []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel},
		})

		l.AddHook(&writeHook{
			Writer:    []io.Writer{warningFile},
			LogLevels: []logrus.Level{logrus.WarnLevel},
		})

		l.AddHook(&writeHook{
			Writer:    []io.Writer{infoFile},
			LogLevels: []logrus.Level{logrus.InfoLevel},
		})

	}

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
