package logger

import (
	"fmt"
	"os"
	"rosenchat/src/configs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var conf = configs.Get()

// implZap implements ILogger using uber/zap package.
type implZap struct {
	zapLogger *zap.Logger
}

func (i *implZap) Debugf(format string, a ...interface{}) {
	packageInfo, file, line := getCallerDetails()

	format = fmt.Sprintf("%s %d %s: %s", file, line, packageInfo, format)
	i.zapLogger.Sugar().Debugf(format, a...)
}

func (i *implZap) Infof(format string, a ...interface{}) {
	packageInfo, file, line := getCallerDetails()

	format = fmt.Sprintf("%s %d %s: %s", file, line, packageInfo, format)
	i.zapLogger.Sugar().Infof(format, a...)
}

func (i *implZap) Warnf(format string, a ...interface{}) {
	packageInfo, file, line := getCallerDetails()

	format = fmt.Sprintf("%s %d %s: %s", file, line, packageInfo, format)
	i.zapLogger.Sugar().Warnf(format, a...)
}

func (i *implZap) Errorf(format string, a ...interface{}) {
	packageInfo, file, line := getCallerDetails()

	format = fmt.Sprintf("%s %d %s: %s", file, line, packageInfo, format)
	i.zapLogger.Sugar().Errorf(format, a...)
}

func (i *implZap) init() {
	logFilePointer, err := getLogFilePointer(conf.Logger.FilePath)
	if err != nil {
		panic(err)
	}

	zapLevelMap := map[string]zapcore.Level{
		"debug": zap.DebugLevel,
		"info":  zap.InfoLevel,
		"warn":  zap.WarnLevel,
		"error": zap.ErrorLevel,
	}

	// More logging destinations can be included here, example: Kafka.
	devOut := []zapcore.WriteSyncer{os.Stdout}
	prodOut := []zapcore.WriteSyncer{logFilePointer}

	i.zapLogger = i.createZapLogger(zapLevelMap[conf.Logger.Level], devOut, prodOut)
}

// createZapLogger creates and returns a new zap logger.
//
// 1. level controls the level at which the returned logger will log.
//
// 2. devOut are logging destinations that will receive pretty and human-readable logs.
//
// 3. prodOut are logging destinations that will receive machine-readable (JSON) logs.
func (i *implZap) createZapLogger(level zapcore.Level, devOut []zapcore.WriteSyncer, prodOut []zapcore.WriteSyncer) *zap.Logger {
	var devCores []zapcore.Core
	var prodCores []zapcore.Core

	levelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})

	for _, dest := range devOut {
		encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		devCores = append(devCores, zapcore.NewCore(encoder, zapcore.Lock(dest), levelEnabler))
	}

	for _, dest := range prodOut {
		encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		prodCores = append(prodCores, zapcore.NewCore(encoder, zapcore.Lock(dest), levelEnabler))
	}

	allCores := append(devCores, prodCores...)
	return zap.New(zapcore.NewTee(allCores...))
}
