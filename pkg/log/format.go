package log

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type Formatter struct {
	pid string
}

// 日志颜色
var colors = map[logrus.Level]func(format string, a ...interface{}) string{
	logrus.WarnLevel:  color.New(color.FgYellow).Add(color.Bold).SprintfFunc(),
	logrus.PanicLevel: color.New(color.BgHiRed).Add(color.Bold).SprintfFunc(),
	logrus.FatalLevel: color.New(color.BgRed).Add(color.Bold).SprintfFunc(),
	logrus.ErrorLevel: color.New(color.FgRed).Add(color.Bold).SprintfFunc(),
	logrus.InfoLevel:  color.New(color.FgCyan).Add(color.Bold).SprintfFunc(),
	logrus.DebugLevel: color.New(color.FgWhite).Add(color.Bold).SprintfFunc(),
}

const (
	FieldLevel      = "level"
	FieldError      = "error"
	FieldGinContext = "gin_context"
	FieldUsername   = "username"
)

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	colorFunc := colors[entry.Level]

	level := entry.Level.String()
	if entry.Data[FieldLevel] != nil {
		level = entry.Data[FieldLevel].(string)
	}

	switch entry.Logger.Level {
	case logrus.DebugLevel:
		level = colorFunc("%-11s", "["+strings.ToUpper(level)+"]")
	default:
		level = colorFunc("%-7s", "["+strings.ToUpper(level)+"]")
	}

	message := entry.Message

	if entry.Data[FieldGinContext] != nil {
		message = fmt.Sprintf("%s, request id: %s", message, entry.Data[FieldGinContext].(*gin.Context).MustGet("request_id").(string))
	}

	if entry.Data[FieldUsername] != nil {
		message = fmt.Sprintf("%s, user: %s", message, entry.Data[FieldUsername])
	}

	if entry.Data[FieldError] != nil {
		message = fmt.Sprintf("%s, err: \"%s\"", message, entry.Data[FieldError].(error))
	}

	return []byte(fmt.Sprintf(
		"%s %s | %s | %s\n",
		level,
		f.pid,
		entry.Time.Format("2006-01-02 15:04:05.000"),
		message,
	)), nil
}

func NewFormatter() *Formatter {
	return &Formatter{
		pid: color.New(color.FgHiMagenta).Sprint(os.Getpid()),
	}
}
