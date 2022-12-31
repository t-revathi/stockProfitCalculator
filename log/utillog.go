package log

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LoggerKey struct{}

type UtilLogger struct {
	fileLogger zerolog.Logger
}

func Newlogger(logLevel zerolog.Level, tempFileName string) *UtilLogger {
	zerolog.SetGlobalLevel(logLevel)
	if tempFileName == "" {
		tempFileName = "deletelog"
	}
	tempFile, err := ioutil.TempFile(os.TempDir(), tempFileName)
	if err != nil {
		log.Error().Err(err).Msg("error creating log temp file")
	}
	fileLogger := zerolog.New(tempFile).With().Logger()
	//Enable when api ,so can send request id as fields
	//fileLogger = fileLogger.With().Str("foo", "bar").Logger()
	return &UtilLogger{fileLogger}
}
func AddLoggerToContext(ctx context.Context, logger *UtilLogger) context.Context {
	ctx = context.WithValue(ctx, &LoggerKey{}, logger)
	return ctx
}

func getLoggerFromContext(ctx context.Context) *UtilLogger {

	ctxloggerInterface := ctx.Value(&LoggerKey{})

	if ctxloggerInterface == nil {
		return nil
	}

	return ctxloggerInterface.(*UtilLogger)

}

func (ul *UtilLogger) Debug(msg string) {
	ul.fileLogger.Debug().Msg(msg)
}

func Debug(ctx context.Context, msg string) {

	if ctxlogger1 := getLoggerFromContext(ctx); ctxlogger1 != nil {

		ctxlogger1.Debug(msg)
	}
}

func (ul *UtilLogger) Info(msg string) {
	ul.fileLogger.Info().Msg(msg)
}

func Info(ctx context.Context, msg string) {
	if ctxlogger1 := getLoggerFromContext(ctx); ctxlogger1 != nil {
		ctxlogger1.Info(msg)
	}

}
func (ul *UtilLogger) Error(msg string, err error) {
	ul.fileLogger.Error().Err(err).Msg(msg)
}
func Error(ctx context.Context, msg string, err error) {

	if ctxlogger1 := getLoggerFromContext(ctx); ctxlogger1 != nil {
		ctxlogger1.Error(msg, err)
	}
}
func (ul *UtilLogger) Warning(msg string) {
	ul.fileLogger.Warn().Msg(msg)
}
func Warning(ctx context.Context, msg string) {
	if ctxlogger1 := getLoggerFromContext(ctx); ctxlogger1 != nil {
		ctxlogger1.Warning(msg)
	}
}
func (ul *UtilLogger) Fatal(msg string, err error) {
	ul.fileLogger.Fatal().Err(err).Msg(msg)
}
func Fatal(ctx context.Context, msg string, err error) {
	ctxlogger := ctx.Value(&LoggerKey{})
	ctxlogger1 := ctxlogger.(*UtilLogger)
	ctxlogger1.Fatal(msg, err)
}
