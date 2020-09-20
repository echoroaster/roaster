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
	"net/http"

	"github.com/echoroaster/roaster/pkg/monitoring"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
)

func init() {
	view.Register(ochttp.ServerRequestCountView,
		ochttp.ServerRequestBytesView,
		ochttp.ServerResponseBytesView,
		ochttp.ServerLatencyView,
		ochttp.ServerRequestCountByMethod,
		ochttp.ServerResponseCountByStatusCode,
	)
}

var MonitorMiddleware = MiddlewareFunc(func(next http.Handler) http.Handler {
	return &ochttp.Handler{
		Propagation: monitoring.HTTPFormat,
		Handler:     next,
	}
})
