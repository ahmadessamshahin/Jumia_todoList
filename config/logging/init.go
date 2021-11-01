package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Log *zerolog.Logger

func init() {
	fmt.Println("START LOADING LOGGER ....")

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := log.With().Logger()
	Log = &logger

	fmt.Println("FINISH LOADING LOGGER 🚀")
}
