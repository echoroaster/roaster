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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(hooks []logHook) (logger Logger, cleanup func(), err error) {
	var zapLogger *zap.Logger
	zapLogger, cleanup, err = newZap(hooks)
	if err != nil {
		return
	}
	logger = &zapFormatAdapter{&zapAdapter{zapLogger}}
	return
}

func newZap(hooks []logHook) (logger *zap.Logger, cleanup func(), err error) {
	zapHooks := make([]func(entry zapcore.Entry) error, len(hooks))
	for idx, hook := range hooks {
		zapHooks[idx] = zapHookAdapter(hook)
	}

	logger, err = zap.NewProduction(
		zap.AddCaller(),
		zap.Hooks(zapHooks...),
	)

	if err != nil {
		return
	}
	cleanup = func() {
		_ = logger.Sync()
	}
	return
}
