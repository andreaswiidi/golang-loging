package loging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("Username", "awe").Info("Hello World")

	logger.WithField("username", "awe").
		WithField("fullname", "andreas widi").
		Info("hello world")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("Username", "awe").Info("Hello World")

	logger.WithFields(logrus.Fields{
		"username": "awe",
		"fullname": "andreas widi",
	}).Info("hello world")
}
