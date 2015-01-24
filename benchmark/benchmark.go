package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"

	//"github.com/davecheney/profile"

	//"github.com/Sirupsen/logrus"
	"github.com/hamaxx/logrus"
)

func main() {
	//defer profile.Start(&profile.Config{MemProfile: true, ProfilePath: "."}).Stop()

	log := logrus.New()
	log.Out = ioutil.Discard
	//log.Out = os.Stdout
	log.Formatter = &logrus.JSONFormatter{}
	log.Level = logrus.DebugLevel

	t0 := time.Now()

	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			log.WithField("AAA", "BBBB").Info("Nice message alalsjlkfldkjf dkf jlsdf")
			log.WithFields(logrus.Fields{"AAA": "BBBB"}).Info("Nice message alalsjlkfldkjf dkf jlsdf")
			log.Info("Nice message alalsjlkfldkjf dkf jlsdf")
		}
		time.Sleep(time.Millisecond)
	}

	ms := &runtime.MemStats{}
	runtime.ReadMemStats(ms)

	ml := ms.Mallocs
	ta := ms.TotalAlloc
	ngc := ms.NumGC

	fmt.Println(ml, ta, ngc)

	fmt.Println(time.Since(t0))
}
