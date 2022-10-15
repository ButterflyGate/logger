# logger
ButterflyGate's standard logger

## level
logging levels

| name          | number | destination | description |
| ---:          | :----: | :---------: | :---------- |
| None          | 0      | stderr      | output nothing |
| Emergency     | 1      | stderr      | only Emergency level output |
| Alert         | 2      | stderr      | Alert and Emergency level log |
| Critical      | 3      | stderr      | Similarly, above levels and this level |
| Error         | 4      | stderr      | Similarly, above levels and this level |
| Warning       | 5      | stderr      | Similarly, above levels and this level |
| Notice        | 6      | stdout      | Similarly, above levels and this level |
| Informational | 7      | stdout      | Similarly, above levels and this level |
| Debug         | 8      | stdout      | Similarly, above levels and this level |
| Trace         | 99     | stdout      | Similarly, above levels and this level |

## sample

### string case
- usage

```go
package main
import(
    "github.com/ButterflyGate/logger"
	. "github.com/ButterflyGate/logger/levels"
)
func main(){
	l := logger.NewLogger(
		Trace,
	)
	l.Trace("hello,world")
}
```

- output

```json
{
  "level": "Info",
  "timestamp": "2022-10-15T13:34:58.084277854+09:00",
  "cursor": "/home/ampamman/go/src/logger/test/log_test.go:11",
  "message": [
    "successfly created logger struct"
  ]
}
{
  "level": "Trace",
  "timestamp": "2022-10-15T13:35:07.987882308+09:00",
  "cursor": "~/go/src/logger/test/log_test.go:14",
  "message": [
    "hello,world"
  ]
}
```

### structure case
- usage

```go
func main(){
	l := logger.NewLogger(Notice)
    yourStruct := something.NewYourStruct()
	l.Notice("hello,world")
}
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
