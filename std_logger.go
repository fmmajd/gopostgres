package gopostgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"strings"
	"time"
)

type StdLogger struct{}

func (l StdLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	timestamp := time.Now().Unix()
	var lvlStr string
	switch level {
	case pgx.LogLevelNone:
		lvlStr = "None"
	case pgx.LogLevelError:
		lvlStr = "Error"
	case pgx.LogLevelWarn:
		lvlStr ="Warn"
	case pgx.LogLevelInfo:
		lvlStr = "Info"
	case pgx.LogLevelDebug:
		lvlStr = "Debug"
	case pgx.LogLevelTrace:
		lvlStr = "Trace"
	default:
		lvlStr = "Unidentified"
	}
	msg = fmt.Sprintf("%d | %s: %s", timestamp, strings.ToUpper(lvlStr), msg)
	log.Println(msg)
	log.Println(data)
}
