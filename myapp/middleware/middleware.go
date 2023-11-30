package middleware

import (
	"myapp/data"

	"github.com/daddy2054/celeritas"
)

type Middleware struct {
	App *celeritas.Celeritas
	Models data.Models
}