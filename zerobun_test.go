package zerobun

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

func TestQueryHookDebug(t *testing.T) {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	qh := NewQueryHook(QueryHookOptions{Logger: &logger})

	event := &bun.QueryEvent{
		StartTime: time.Now(),
		Query:     "",
		Err:       nil,
	}

	qh.AfterQuery(
		context.Background(),
		event,
	)
}
