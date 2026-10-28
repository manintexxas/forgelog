package domain

type Level string

type Field struct {
	Key   string 	`json:"key"`
	Value any		`json:"value"`
}

const (
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO"
	LevelWarn  Level = "WARNING"
	LevelError Level = "ERROR"
	LevelFatal Level = "FATAL"
)

type LogEntry struct {
	TimeStamp int64 	`json:"timestamp"`
	LogLevel  Level		`json:"level"`
	Message   string	`json:"message"`
	Fields    []Field	`json:"fields,omitempty"`
	Caller    string	`json:"caller,omitempty"`
}
