package goldilocks

import (
	"log"
	"os"
)

type Logger struct {
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
	fatalLog   *log.Logger
}

/*
Return a default logger.
*/
func NewLogger() (logger *Logger) {
	return &Logger{
		infoLog:    log.New(os.Stdout, "INFO: ", log.LstdFlags),
		warningLog: log.New(os.Stdout, "WARNING: ", log.LstdFlags),
		errorLog:   log.New(os.Stderr, "ERROR: ", log.LstdFlags),
		fatalLog:   log.New(os.Stderr, "FATAL: ", log.LstdFlags),
	}
}

/*
Print output to Stdout by default.
*/
func (logger *Logger) Info(v ...interface{}) {
	logger.infoLog.Println(v...)
}

/*
Print formatted output to Stdout by default.
*/
func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.infoLog.Printf(format, v...)
}

/*
Print output to Stdout by default.
*/
func (logger *Logger) Warning(v ...interface{}) {
	logger.warningLog.Println(v...)
}

/*
Print formatted output to Stdout by default.
*/
func (logger *Logger) Warningf(format string, v ...interface{}) {
	logger.warningLog.Printf(format, v...)
}

/*
Print output to Stderr by default.
*/
func (logger *Logger) Error(v ...interface{}) {
	logger.errorLog.Println(v...)
}

/*
Print formatted output to Stderr by default.
*/
func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.errorLog.Printf(format, v...)
}

/*
Print output to Stderr and exit with os.Exit(1).
*/
func (logger *Logger) Fatal(v ...interface{}) {
	logger.fatalLog.Fatalln(v...)
}

/*
Print formatted output to Stderr and exit with os.Exit(1).
*/
func (logger *Logger) Fatalf(format string, v ...interface{}) {
	logger.fatalLog.Fatalf(format, v...)
}
