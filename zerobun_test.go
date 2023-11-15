package zerobun

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rzajac/zltest"
	"github.com/uptrace/bun"
)

func TestQueryHookDebug(t *testing.T) {
	zlt := zltest.New(t)
	logger := zerolog.New(zlt).With().Timestamp().Logger()

	qh := NewQueryHook(QueryHookOptions{Logger: &logger})

	event := &bun.QueryEvent{
		StartTime: time.Now(),
		Query:     "SELECT * FROM products",
		Err:       nil,
	}

	qh.AfterQuery(
		context.Background(),
		event,
	)

	ent := zlt.LastEntry()

	bunInfo := map[string]interface{}{
		OperationQueryName: event.Query,
		OperationFieldName: event.Operation(),
	}

	tBunInfo, _ := ent.Map(BunInfoFieldName)

	for k, v := range bunInfo {
		if v != tBunInfo[k] {
			t.Errorf(
				"key ['%s'] returned '%s', expected '%s'",
				k,
				bunInfo[k],
				v,
			)
		}
	}

	ent.ExpLevel(zerolog.DebugLevel)
}

func TestQueryHookError(t *testing.T) {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	qh := NewQueryHook(QueryHookOptions{Logger: &logger})

	event := &bun.QueryEvent{
		StartTime: time.Now(),
		Query:     "SELECT * FROM products WHERE ID = 5",
		Err:       errors.New("database error"),
	}
	qh.AfterQuery(
		context.Background(),
		event,
	)
}
