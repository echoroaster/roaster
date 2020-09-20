// Copyright 2020 EchoRoaster
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package logging // import "github.com/echoroaster/roaster/pkg/logging"

import (
	"fmt"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

func zapHookAdapter(h logHook) func(entry zapcore.Entry) error {
	levelMap := map[zapcore.Level]Level{
		zapcore.DebugLevel:  Debug,
		zapcore.InfoLevel:   Info,
		zapcore.WarnLevel:   Warn,
		zapcore.ErrorLevel:  Error,
		zapcore.DPanicLevel: Panic,
		zapcore.PanicLevel:  Panic,
		zapcore.FatalLevel:  Fatal,
	}
	return func(entry zapcore.Entry) error {
		return h.Log(&logEntry{
			Level:   levelMap[entry.Level],
			Time:    entry.Time,
			Message: entry.Message,
			Frame: logFrame{
				Function: entry.Caller.Function,
				Line:     entry.Caller.Line,
				File:     entry.Caller.File,
			},
		})
	}
}

type zapAdapter struct {
	logger *zap.Logger
}

func (a *zapAdapter) Debug(args ...interface{}) {
	a.logger.Debug(fmt.Sprint(args...))
}

func (a *zapAdapter) Info(args ...interface{}) {
	a.logger.Info(fmt.Sprint(args...))
}

func (a *zapAdapter) Print(args ...interface{}) {
	a.Info(args...)
}

func (a *zapAdapter) Error(args ...interface{}) {
	a.logger.Error(fmt.Sprint(args...))
}

func (a *zapAdapter) Fatal(args ...interface{}) {
	a.logger.Fatal(fmt.Sprint(args...))
}

func (a *zapAdapter) Panic(args ...interface{}) {
	a.logger.DPanic(fmt.Sprint(args...))
}

func (a *zapAdapter) WithField(key string, value interface{}) Logger {
	return &zapAdapter{a.logger.With(zap.Any(key, value))}
}

func (a *zapAdapter) WithFields(fields map[string]interface{}) Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}
	return &zapAdapter{a.logger.With(zapFields...)}
}

func (a *zapAdapter) WithError(err error) Logger {
	return &zapAdapter{a.logger.With(zap.Error(err))}
}

type zapFormatAdapter struct {
	*zapAdapter
}

func (a *zapFormatAdapter) Debugf(format string, args ...interface{}) {
	a.Debug(fmt.Sprintf(format, args...))
}

func (a *zapFormatAdapter) Infof(format string, args ...interface{}) {
	a.Info(fmt.Sprintf(format, args...))
}

func (a *zapFormatAdapter) Printf(format string, args ...interface{}) {
	a.Print(fmt.Sprintf(format, args...))
}

func (a *zapFormatAdapter) Errorf(format string, args ...interface{}) {
	a.Error(fmt.Sprintf(format, args...))
}

func (a *zapFormatAdapter) Fatalf(format string, args ...interface{}) {
	a.Fatal(fmt.Sprintf(format, args...))
}

func (a *zapFormatAdapter) Panicf(format string, args ...interface{}) {
	a.Panic(fmt.Sprintf(format, args...))
}
