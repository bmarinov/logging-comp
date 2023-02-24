package deepdown

import (
	"github.com/rs/zerolog"
)

type SpicyBusinessLogic struct {
	state  map[int]string
	logger zerolog.Logger
}

func NewSpicyBusinessLogic(logger *zerolog.Logger) *SpicyBusinessLogic {
	foo := logger.With().Str("package_name", "deepdown").Logger()
	return &SpicyBusinessLogic{
		state:  make(map[int]string),
		logger: foo,
	}
}
