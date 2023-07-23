package internals

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05.999999999",
		FullTimestamp:   true,
	})
	return log
}

func (l *Logger) InfoWithRequest(req *http.Request, message string) {
	l.WithFields(logrus.Fields{
		"method": req.Method,
		"uri":    req.RequestURI,
		"remote": req.RemoteAddr,
	}).Info(message)
}

func (l *Logger) ErrorWithRequest(req *http.Request, message string) {
	l.WithFields(logrus.Fields{
		"method": req.Method,
		"uri":    req.RequestURI,
		"remote": req.RemoteAddr,
	}).Error(message)
}
