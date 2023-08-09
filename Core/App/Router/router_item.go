package Router

import (
	"net/http"
)

type Route struct {
	Pattern string
	Handler http.Handler
}
