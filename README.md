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

## Options
- these are configurable elements for output or hide
  - log level
  - timestamp
  - called cursor

- you can specify output format
  - using indent or not (json format only)
  - json or text (not implement yet)
  - custom format (not implement yet)


## sample

### string case
- usage

```go
l.Trace("hello,world")
```

- output

```json
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
