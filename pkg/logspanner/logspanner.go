package logspanner

import (
	"github.com/go-kit/kit/log"
)

type LogSpanner struct {
	keyvals []interface{} // Note: array not map
	logger  log.Logger
}

func New(logger log.Logger) *LogSpanner {
	l := &LogSpanner{
		logger: logger,
	}

	return l
}

func (l *LogSpanner) Log(keyvals ...interface{}) error {
	// log as is, this should come in with a level?
	l.logger.Log(keyvals...)

	// append to the running keyvals
	// Filter out the level, so it doesn't hang around
	for i := 0; i < len(keyvals); i += 2 {
		if keyvals[i] == "level" {
			keyvals = append(keyvals[:i], keyvals[i+2:]...)
		}
	}
	l.keyvals = append(l.keyvals, keyvals...)
	return nil
}

func (l *LogSpanner) End(levelLogger func(log.Logger) log.Logger) {
	l.keyvals = append(l.keyvals, "endspan", true)
	levelLogger(l.logger).Log(l.keyvals...)
}
