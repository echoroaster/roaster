package common // import "github.com/echoroaster/roaster/pkg/common"

import "context"

type Runner interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
