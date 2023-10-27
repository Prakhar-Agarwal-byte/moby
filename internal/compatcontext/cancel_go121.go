//go:build go1.21

package compatcontext // import "github.com/Prakhar-Agarwal-byte/moby/internal/compatcontext"

import "context"

func WithoutCancel(ctx context.Context) context.Context {
	return context.WithoutCancel(ctx)
}
