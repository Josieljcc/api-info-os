package config

var (
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	if logger == nil {
		logger = NewLogger(p)
	}
	return logger
}
