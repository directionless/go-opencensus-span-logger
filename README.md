# Opencensus Span Logger

In playing
[ctxlog](https://github.com/kolide/launcher/blob/master/pkg/contexts/ctxlog/ctxlog.go),
I found myself wanting a slightly different pattern of logging. Something suitable to support:

* Production logs
* Debug Logs
* Trace level / printf logs
* Structured logging
* clear support for trace spans

I thought I'd start playing...

## An evolution

1. [simple-ctxlog](cmd/01-simple-ctxlog/simple-ctxlog.go) Simplest example of ctxlog
2. [manysteps-ctxlog](cmd/02-manysteps-ctxlog/manysteps-ctxlog.go) multistep process. You can see the printf style things start becoming quite verbose
3. [manysteps-logspan](cmd/03-manysteps-logspan/manysteps-logspan.go) Now with logspan. It gets a bit cleaner to see both the end state, and the intermediary printf
