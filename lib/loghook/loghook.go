package loghook

import (
	"runtime"
	"strings"

	"github.com/rs/zerolog"
)

var Caller = zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
	pc, file, line, ok := runtime.Caller(zerolog.CallerSkipFrameCount + 2)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	fnName := fn.Name()[strings.LastIndex(fn.Name(), "/")+1:]

	e.Str(zerolog.CallerFieldName, zerolog.CallerMarshalFunc(file, line))
	e.Str("func", fnName)
})

var Severity = zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.NoLevel {
		e.Str("severity", level.String())
	}
})
