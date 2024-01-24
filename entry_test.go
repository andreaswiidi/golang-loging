package loging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.WithField("username", "awe").Info("hello world")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "awe")
	entry.Info("hello entry")
}
