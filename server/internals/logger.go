package internals

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Use ISO8601 format for timestamps
	logger, err := config.Build()
	if err != nil {
		// Fallback to a basic logger if Zap initialization fails
		log := logrus.New()
		log.SetOutput(os.Stdout)
		log.SetLevel(logrus.InfoLevel)
		log.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			TimestampFormat: "2006-01-02 15:04:05.999999999",
			FullTimestamp:   true,
		})
		log.Warn("Failed to initialize Zap logger. Falling back to logrus.")
		return nil
	}

	return logger
}
