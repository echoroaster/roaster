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
	"time"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	WithError(err error) Logger
}

type FormatLogger interface {
	Logger

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

//go:generate stringer -type=Level
type Level uint8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	Debug Level = iota
	// InfoLevel is the default logging priority.
	Info
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	Warn
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	Error
	// PanicLevel logs a message, then panics.
	Panic
	// FatalLevel logs a message, then calls os.Exit(1).
	Fatal
)

type logEntry struct {
	Level   Level
	Time    time.Time
	Message string
	Frame   logFrame
}

type logFrame struct {
	// Function is the package path-qualified function name of
	// this call frame. If non-empty, this string uniquely
	// identifies a single function in the program.
	// This may be the empty string if not known.
	// If Func is not nil then Function == Func.Name().
	Function string

	// File and Line are the file name and line number of the
	// location in this frame. For non-leaf frames, this will be
	// the location of a call. These may be the empty string and
	// zero, respectively, if not known.
	File string
	Line int
}

type logHook interface {
	Log(entry *logEntry) error
}
