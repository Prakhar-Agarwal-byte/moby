package common // import "github.com/Prakhar-Agarwal-byte/moby/integration/plugin/common"

import (
	"context"
	"os"
	"testing"

	"github.com/Prakhar-Agarwal-byte/moby/testutil"
	"github.com/Prakhar-Agarwal-byte/moby/testutil/environment"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

var (
	testEnv     *environment.Execution
	baseContext context.Context
)

func TestMain(m *testing.M) {
	shutdown := testutil.ConfigureTracing()
	ctx, span := otel.Tracer("").Start(context.Background(), "integration/plugin/common.TestMain")
	baseContext = ctx

	var err error
	testEnv, err = environment.New(ctx)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.End()
		shutdown(ctx)
		panic(err)
	}
	testEnv.Print()
	ec := m.Run()
	if ec != 0 {
		span.SetStatus(codes.Error, "m.Run() returned non-zero exit code")
	}
	span.SetAttributes(attribute.Int("exit", ec))
	shutdown(ctx)
	os.Exit(ec)
}

func setupTest(t *testing.T) context.Context {
	ctx := testutil.StartSpan(baseContext, t)
	environment.ProtectAll(ctx, t, testEnv)
	t.Cleanup(func() {
		testEnv.Clean(ctx, t)
	})
	return ctx
}
