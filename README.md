[![Go Reference](https://pkg.go.dev/badge/github.com/comsma/zerobun.svg)](https://pkg.go.dev/github.com/comsma/zerobun)
[![Go Report Card](https://goreportcard.com/badge/github.com/comsma/zerobun)](https://goreportcard.com/report/github.com/comsma/zerobun)
# zerobun
Zerobun is a query hook for [uptrace/bun](https://github.com/uptrace/bun) that logs with [rs/zerolog](https://github.com/rs/zerolog)
```bash
go get -u github.com/comsma/zerobun
```
## Usage

```go
package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	db := bun.NewDb()
	db.AddQueryHook(
		zerobun.NewQueryHook(
			zerobun.QueryHookOptions{
				SlowDuration: 5,
				Logger:       &logger,
			},
		),
	)
}
```
## Example Output
```json
{ "level":"error", 
  "bun_info":{
    "operation":"SELECT",
    "operation_time_ms":0,
    "query":"SELECT * FROM products WHERE ID = 5",
  },
  "error":"database error",
  "time":"2023-11-15T16:12:24-05:00"}
```