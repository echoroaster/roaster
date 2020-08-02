module github.com/echoroaster/roaster/pkg/httpclient

go 1.14

replace (
	github.com/echoroaster/roaster/pkg/auth => ../auth
	github.com/echoroaster/roaster/pkg/logging => ../logging
	github.com/echoroaster/roaster/pkg/monitoring => ../monitoring
)

require (
	github.com/echoroaster/roaster/pkg/auth v0.0.0-00010101000000-000000000000
	github.com/echoroaster/roaster/pkg/logging v0.0.0-00010101000000-000000000000 // indirect
	github.com/echoroaster/roaster/pkg/monitoring v0.0.0-00010101000000-000000000000
	github.com/google/wire v0.4.0
	go.opencensus.io v0.22.4
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
)
