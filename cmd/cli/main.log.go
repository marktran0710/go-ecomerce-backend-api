package main

import (
	"os"

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

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Info error", zap.Int("line", 2))

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

func getWriterSync() zapcore.WriteSyncer {
	if _, err := os.Stat("./log"); os.IsNotExist(err) {
		os.Mkdir("./log", os.ModePerm)
	}

	file, err := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync((os.Stderr))
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
