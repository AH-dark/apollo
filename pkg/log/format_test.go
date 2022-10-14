package log

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFormat(t *testing.T) {
	asserts := assert.New(t)

	logger := logrus.New()
	logger.SetFormatter(&Formatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	asserts.NotEmpty(logger.Formatter.(*Formatter).Format(logger.WithFields(logrus.Fields{"test": "test"})))
}

func BenchmarkTestFormatter_Format(b *testing.B) {
	logger := logrus.New()
	logger.SetFormatter(&Formatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Formatter.(*Formatter).Format(logger.WithFields(logrus.Fields{"test": "test"}))
	}
	b.StopTimer()
}
