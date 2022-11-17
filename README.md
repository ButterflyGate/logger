# Logger
ButterflyGate's standard logger

## Language
- [Japanese](./readme.d/README-Japanese.md)
- [English](./readme.d/README-English.md)

## Usage
- simple usage

```go
import (
	"github.com/ButterflyGate/logger"
	"github.com/ButterflyGate/logger/levels"
)

func main(){
	// definition and specify log level
	logger := logger.NewLogger(
		levels.Trace,
	)

	// output message
	logger.Info("hello,world\nhello\nhello")
}
```

### Definition
- use `NewLogger`.
- Log level must be specified to use this logger at definition.
- log levels are definition in logger/levels directory.

```go
logger := logger.NewLogger(
	levels.Trace,
)
```

### Options
- You can customize the output.
- The function `NewLogger` can take Optional Arguments controlling output format.

```go
outputOption := logger.DefaultOutputOption().HideCursor().HideLevel().HideTimestamp()
formatOption := logger.DefaultFormatOption().FormatMessageRowLimit(0)
l := logger.NewLogger(
	levels.Trace,
	outputOption, formatOption,
)
```

#### Output Option
- `OutputOption` can control these elements for output or hide
  - log level
  - timestamp
  - called cursor
  - message (change source)
  - struct data (change source)
  - struct name (change source)

```go
option := logger.DefaultOutputOption().HideLevel().HideTimestamp().HideCursor()
l := logger.NewLogger(
	levels.Trace,
	option,
)
```

#### Format Option
- `FormatOption` can control output format
  - using indent or not (json format only)
  - limit message rows count (string type log only)
  - json or text (not implement yet)
  - custom format (not implement yet)

```go
option := logger.DefaultFormatOption().FormatMessageRowLimit(1)
l := logger.NewLogger(
	levels.Trace,
	option,
)
```

## Sample Output
### string case
- usage

```go
logger.Info("hello\nworld")
```

- output

```json
{
  "level": "Trace",
  "timestamp": "2022-10-15T13:35:07.987882308+09:00",
  "cursor": "~/go/src/logger/test/log_test.go:14",
  "message": [
    "hello",
    "world"
  ]
}
```

### structure case
- usage

```go
yourStruct := something.NewYourStruct()
l.Notice(yourStruct)
```

- output

```json
{
  "level": "Notice",
  "timestamp": "2022-10-15T14:00:16.889360303+09:00",
  "cursor": "/go/src/logger/test/log_test.go:28",
  "structure_data": {
    "ID": 10,
    "FullName": "kyota tahsiro",
    "Birth": "2022-10-15T14:00:15.527879915+09:00"
  },
  "structure_name": "unknown"
}
```

- it takes second argument for `structure_name`
- usage

```go
yourStruct := something.NewYourStruct()
l.Notice(yourStruct,"yourStruct")
```

- output

```json
{
  "level": "Notice",
  "timestamp": "2022-10-15T14:00:16.889360303+09:00",
  "cursor": "/go/src/logger/test/log_test.go:28",
  "structure_data": {
    "ID": 10,
    "FullName": "kyota tahsiro",
    "Birth": "2022-10-15T14:00:15.527879915+09:00"
  },
  "structure_name": "yourStruct"
}
```
