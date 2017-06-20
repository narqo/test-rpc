package logger

type Logger interface {
	Print(...interface{}) error
	Printf(...interface{}) error
	Fatal(...interface{}) error
	Fatalf(...interface{}) error
}
