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

package httpserver // import "github.com/echoroaster/roaster/pkg/httpserver"

import (
	"io"
	"net/http"
	"os"

	"github.com/echoroaster/roaster/pkg/logging"
	"github.com/gorilla/handlers"
)

var (
	ProxyMiddleware   Middleware = MiddlewareFunc(handlers.ProxyHeaders)
	LoggingMiddleware Middleware = MiddlewareFunc(func(next http.Handler) http.Handler {
		return handlers.CombinedLoggingHandler(os.Stdout, next)
	})
)

func NewLoggingMiddleware(w io.Writer) Middleware {
	return MiddlewareFunc(func(next http.Handler) http.Handler {
		return handlers.CombinedLoggingHandler(w, next)
	})
}

func NewPanicMiddleware(logger logging.Logger) Middleware {
	return MiddlewareFunc(handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(true),
		handlers.RecoveryLogger(&panicLogger{logger}),
	))
}

type panicLogger struct {
	logging.Logger
}

func (l *panicLogger) Println(args ...interface{}) {
	l.Logger.Panic(args...)
}
