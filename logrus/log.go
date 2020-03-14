package logrus

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func InitLog() (*logrus.Logger, error) {
	//dir, _ := os.Executable()
	//exPath := filepath.Dir(dir)
	//logPath := path.Join(exPath, "test.log")
	logPath := "/Users/yuchanns/Coding/golang/gobyexample/logrus/log/test"

	writer, err := rotatelogs.New(
		strings.Join([]string{logPath, "%Y%m%d%H%M", "log"}, "."),
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err != nil {
		return nil, err
	}

	jsonFormatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel: writer,
	}, jsonFormatter)

	logger := logrus.New()

	logger.AddHook(lfHook)

	logger.SetFormatter(jsonFormatter)

	return logger, nil
}
