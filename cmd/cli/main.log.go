package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello, name: %s, age: %d", "Hautran", 23)

	// // logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "Hautran"), zap.Int("age", 10))

	// logger := zap.NewExample()
	// logger.Info("heloo")

	// // Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Heloo Development")

	// // Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Helooo Production")

}

// format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriter() {

}
