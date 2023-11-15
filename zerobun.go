package zerobun

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

const (
	OperationFieldName     = "operation"
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

func NewQueryHook(options QueryHookOptions) QueryHook {
	return QueryHook{
		logger:       options.Logger,
		slowDuration: options.SlowDuration,
	}
}

func (qh QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (qh QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	queryDuration := time.Since(event.StartTime)

	fields := map[string]interface{}{
		OperationFieldName:     event.Operation(),
		OperationTimeFieldName: queryDuration.Milliseconds(),
	}

	qh.logger.Debug().Fields(fields).Msg("")

}
