package api

import "github.com/scando/ColdChainTrackerServer/main/src/logging"

var logger *logging.SeelogWrapper

func init() {
	var err error
	logger, err = logging.New()
	if err != nil {
		panic(err)
	}
	Debug(false)
}

func Debug(debugMode bool) {
	if debugMode {
		logger.SetLevel("debug")
	} else {
		logger.SetLevel("info")
	}
}
