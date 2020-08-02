module github.com/echoroaster/roaster/pkg/httpserver

go 1.14

replace (
	github.com/echoroaster/roaster/pkg/auth => ../auth
	github.com/echoroaster/roaster/pkg/common => ../common
	github.com/echoroaster/roaster/pkg/logging => ../logging
	github.com/echoroaster/roaster/pkg/monitoring => ../monitoring
)

require (
	github.com/echoroaster/roaster/pkg/auth v0.0.0-20200802182826-62af7de36742
	github.com/echoroaster/roaster/pkg/common v0.0.0-20200802182826-62af7de36742
	github.com/echoroaster/roaster/pkg/logging v0.0.0-20200802182826-62af7de36742
	github.com/echoroaster/roaster/pkg/monitoring v0.0.0-20200802182826-62af7de36742
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/graphql-go/graphql v0.7.9 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/samsarahq/go v0.0.0-20191220233105-8077c9fbaed5 // indirect
	github.com/samsarahq/thunder v0.5.0
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	go.opencensus.io v0.22.4
)
