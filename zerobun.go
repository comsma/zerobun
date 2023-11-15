package zerobun

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

const (
	BunInfoFieldName       = "bun_info"
	OperationFieldName     = "operation"
	OperationQueryName     = "query"
	OperationTimeFieldName = "operation_time_ms"
)

type QueryHook struct {
	bun.QueryHook

	logger       *zerolog.Logger
	slowDuration time.Duration
}

type QueryHookOptions struct {
	Logger       *zerolog.Logger
	SlowDuration time.Duration
}

// NewQueryHook returns a new bun.QueryHook that outputs database transactions using an instance of zerolog.Logger
func NewQueryHook(options QueryHookOptions) QueryHook {
	return QueryHook{
		logger:       options.Logger,
		slowDuration: options.SlowDuration,
	}
}

// BeforeQuery called before a database transaction but currently does nothing
func (qh QueryHook) BeforeQuery(ctx context.Context, _ *bun.QueryEvent) context.Context {
	return ctx
}

// AfterQuery logs an entry after a database transaction
func (qh QueryHook) AfterQuery(_ context.Context, event *bun.QueryEvent) {
	queryDuration := time.Since(event.StartTime)

	fields := map[string]interface{}{
		OperationQueryName:     event.Query,
		OperationFieldName:     event.Operation(),
		OperationTimeFieldName: queryDuration.Milliseconds(),
	}

	bunLogger := qh.logger.With().Dict(
		BunInfoFieldName,
		zerolog.Dict().Fields(fields),
	).Logger()

	if event.Err != nil {
		bunLogger.Error().Err(event.Err).Msg("")
		return
	}

	bunLogger.Debug().Msg("")

}
