//go:build !windows

package ipvlan // import "github.com/Prakhar-Agarwal-byte/moby/integration/network/ipvlan"

import (
	"context"
	"os"
	"testing"

	"github.com/Prakhar-Agarwal-byte/moby/testutil"
	"github.com/Prakhar-Agarwal-byte/moby/testutil/environment"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var (
	testEnv     *environment.Execution
	baseContext context.Context
)

func TestMain(m *testing.M) {
	shutdown := testutil.ConfigureTracing()
	ctx, span := otel.Tracer("").Start(context.Background(), "integration/network/ipvlan/TestMain")
	baseContext = ctx

	var err error
	testEnv, err = environment.New(ctx)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.End()
		shutdown(ctx)
		panic(err)
	}
	err = environment.EnsureFrozenImagesLinux(ctx, testEnv)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.End()
		shutdown(ctx)
		panic(err)
	}

	testEnv.Print()
	code := m.Run()
	if code != 0 {
		span.SetStatus(codes.Error, "m.Run() returned non-zero exit code")
	}
	span.End()
	shutdown(ctx)
	os.Exit(code)
}
