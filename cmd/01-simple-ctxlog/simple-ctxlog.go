package main

import (
	"context"

	"github.com/directionless/go-opencensus-span-logger/pkg/ctxlog"
	"github.com/go-kit/kit/log/level"
	"github.com/kolide/kit/logutil"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()
	logger := logutil.NewServerLogger(true)
	ctx = ctxlog.NewContext(ctx, logger)
	doThing(ctx)
}

func doThing(ctx context.Context) {
	ctx, span := trace.StartSpan(ctx, "simple-ctxlog.doThing")
	defer span.End()

	logger := ctxlog.FromContext(ctx)

	level.Debug(logger).Log(
		"msg", "Did a thing",
		"code", 123,
	)

}
