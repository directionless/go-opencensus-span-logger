package main

import (
	"context"
	"os"
	"strings"

	"github.com/directionless/go-opencensus-span-logger/pkg/ctxlog"
	"github.com/directionless/go-opencensus-span-logger/pkg/logspanner"
	"github.com/go-kit/kit/log/level"
	"github.com/kolide/kit/logutil"
	"go.opencensus.io/trace"
)

func main() {
	debug := false
	if len(os.Args) > 1 {
		debug = strings.Contains("debug", os.Args[1])
	}

	ctx := context.Background()
	logger := logutil.NewServerLogger(debug)
	ctx = ctxlog.NewContext(ctx, logger)

	thing := 123

	doThing(ctx, thing)

}

func doThing(ctx context.Context, thing int) {
	ctx, span := trace.StartSpan(ctx, "simple-ctxlog.doThing")
	defer span.End()

	logger := logspanner.New(ctxlog.FromContext(ctx))
	defer logger.End(level.Info)

	level.Debug(logger).Log(
		"msg", "Starting a thing",
		"thingId", thing,
	)

	level.Debug(logger).Log(
		"msg", "Sending a thing to upload",
		"thingId", thing,
	)
	uploadThing(ctx, thing)

	level.Debug(logger).Log(
		"msg", "Finished a thing elsewhere",
		"thingId", thing,
	)

}

func uploadThing(ctx context.Context, thing int) {
	ctx, span := trace.StartSpan(ctx, "simple-ctxlog.uploadThing")
	defer span.End()

	logger := logspanner.New(ctxlog.FromContext(ctx))
	defer logger.End(level.Info)

	level.Debug(logger).Log(
		"msg", "Uploading thing",
		"thingId", thing,
	)

}
