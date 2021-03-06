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
	"errors"
	"os"
)

var ErrUnsupportedDriver = errors.New("logging: unsupported logging driver")

func NewLogger() (logger Logger, cleanup func(), err error) {
	hooks := make([]logHook, 0, 1)
	if os.Getenv("AWS_CLOUDWATCHLOGS_GROUP_NAME") != "" {
		hook, err := newCloudWatchLoggerHook()
		if err != nil {
			return nil, nil, err
		}
		hooks = append(hooks, hook)
	}

	loggerDriver := os.Getenv("LOGGER_DRIVER")
	if loggerDriver == "" {
		loggerDriver = "zap"
	}
	switch loggerDriver {
	case "logrus":
		return newLogrusLogger(hooks)
	case "zap":
		return newZapLogger(hooks)
	}
	return nil, nil, ErrUnsupportedDriver
}
