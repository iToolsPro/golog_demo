package elog

import (
	"bytes"
	"fmt"
	"github.com/k0kubun/go-ansi"
	"github.com/op/go-logging"
	"github.com/sirupsen/logrus"
	"github.com/vbauerster/mpb/v5/decor"
	"golog/vars"
	"io"
	"time"
)

var (
	Log *logrus.Logger
)

func init() {
	logger := &logrus.Logger{
		Out:   ansi.NewAnsiStdout(),
		Level: logrus.InfoLevel,
		Formatter: &logrus.TextFormatter{
			TimestampFormat: "15:04:05",
			FullTimestamp:   true,
		},
	}

	Log = logger
}

var format logging.Formatter

func init() {
	format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{longfunc} ▶ %{level:.4s} %{color:reset} %{message}`)
}

type LogFiller struct {
	//record *logging.Record
	msg string
}

func NewFiller(record *logging.Record) *LogFiller {
	return &LogFiller{msg: toMsg(record)}

}
func toMsg(record *logging.Record) string {
	buf := new(bytes.Buffer)
	format.Format(4, record, buf)
	return buf.String()
}

func (f *LogFiller) Fill(w io.Writer, _ int, st decor.Statistics) {
	fmt.Fprint(w, f.msg)
	//data = f.Message()

}

func makeLogBar(msg string) {

	//limit := "%%.%ds"
	//base_format := fmt.Sprintf("%s - %s")

	//return mpb.BarFillerFunc(func(w io.Writer, _ int, st decor.Statistics) {
	//	//format := logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{longfunc} ▶ %{level:.4s} %{color:reset} %{message}`)
	//
	//})
}

func LogWithLevel(msg string, level logging.Level) {
	record := &logging.Record{
		Time: time.Now(),
		//Module: "",
		Args:  []interface{}{msg},
		Level: level,
	}
	vars.ProcessBar.Add(0, NewFiller(record)).SetTotal(0, true)
}
func Info(msg string) {
	LogWithLevel(msg, logging.INFO)
}

func Debug(msg string) {
	LogWithLevel(msg, logging.DEBUG)
}

func Warn(msg string) {
	LogWithLevel(msg, logging.WARNING)
}

func Error(msg string) {
	LogWithLevel(msg, logging.ERROR)
}
