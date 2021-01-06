package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// url := "http://google.com"
	// logger, _ := zap.NewProduction()
	// logger, _ := zap.NewDevelopment()
	// defer logger.Sync() // flushes buffer, if any
	// sugar := logger.Sugar()
	// sugar.Infow("failed to fetch URL",
	// 	// Structured context as loosely typed key-value pairs.
	// 	"url", url,
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )
	// sugar.Infof("Failed to fetch URL: %s", url)
	// sugar.Error("Failed to fetch URL: ", url)
	// fmt.Println("--------------------------------")
	// logger.Info("failed to fetch URL",
	// 	// Structured context as strongly typed Field values.
	// 	zap.String("url", url),
	// 	zap.Int("attempt", 3),
	// 	zap.Duration("backoff", time.Second),
	// )
	// fmt.Println("--------------------------------")
	// logger.Error("Error occurred in file reading")
	// fmt.Println("--------------------------------")
	// logger.Warn("Warning !!!!!!!")

	// _, err := strconv.Atoi("1000f")
	// sugar.Errorw("Error in typecast", "error", err)
	// logger.Info("Processing: ", zap.String("File: ", "filename.txt"))

	log, _ := zap.Config{
		Encoding: "text",
		Level:    zap.NewAtomicLevelAt(zapcore.DebugLevel),
		//OutputPaths:      []string{"stdout"},
		OutputPaths:      []string{"file.log"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()

	log.Debug("some values", zap.String("region", "us-west-2"))

	log.Info("some processing", zap.String("bucket", "some-bucket"))

	log.Error("some error", zap.String("file", "somefile.txt"))
}
