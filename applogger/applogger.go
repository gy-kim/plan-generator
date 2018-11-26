package applogger

import (
	"log"
	"os"
	"sync"
	"time"
)

type appLogger struct {
	*log.Logger
	filename string
	createAt time.Time
}

var alogger *appLogger
var once sync.Once

// GetInstance create a singleton instance of the app logger
func GetInstance() *appLogger {
	once.Do(func() {
		alogger = createLogger("applogger.log")
	})
	return alogger
}

func createLogger(fname string) *appLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0777)

	return &appLogger{
		filename: fname,
		Logger:   log.New(file, "App", log.Lshortfile),
		createAt: time.Now(),
	}
}
