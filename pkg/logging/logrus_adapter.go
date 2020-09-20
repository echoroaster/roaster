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
	"encoding/json"

	"github.com/sirupsen/logrus"
)

var logrusLevelMap = map[logrus.Level]Level{
	logrus.PanicLevel: Panic,
	logrus.FatalLevel: Fatal,
	logrus.ErrorLevel: Error,
	logrus.WarnLevel:  Warn,
	logrus.InfoLevel:  Info,
	logrus.DebugLevel: Debug,
	logrus.TraceLevel: Debug,
}

type logrusHookAdapter struct {
	base logHook
}

func (h *logrusHookAdapter) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *logrusHookAdapter) Fire(entry *logrus.Entry) error {
	type logrusMessage struct {
		Fields  logrus.Fields
		Message string
	}
	message := logrusMessage{
		entry.Data,
		entry.Message,
	}
	bytea, err := json.Marshal(&message)
	if err != nil {
		return err
	}
	return h.base.Log(&logEntry{
		Level:   logrusLevelMap[entry.Level],
		Time:    entry.Time,
		Message: string(bytea),
		Frame: logFrame{
			Function: entry.Caller.Function,
			File:     entry.Caller.File,
			Line:     entry.Caller.Line,
		},
	})
}

type logrusAdapter struct {
	logrus.FieldLogger
}

func (a *logrusAdapter) WithField(key string, value interface{}) Logger {
	return &logrusAdapter{a.FieldLogger.WithField(key, value)}
}

func (a *logrusAdapter) WithFields(fields map[string]interface{}) Logger {
	return &logrusAdapter{a.FieldLogger.WithFields(fields)}
}

func (a *logrusAdapter) WithError(err error) Logger {
	return &logrusAdapter{a.FieldLogger.WithError(err)}
}
