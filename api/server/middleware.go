package server // import "github.com/Prakhar-Agarwal-byte/moby/api/server"

import (
	"github.com/Prakhar-Agarwal-byte/moby/api/server/httputils"
	"github.com/Prakhar-Agarwal-byte/moby/api/server/middleware"
	"github.com/containerd/log"
)

// handlerWithGlobalMiddlewares wraps the handler function for a request with
// the server's global middlewares. The order of the middlewares is backwards,
// meaning that the first in the list will be evaluated last.
func (s *Server) handlerWithGlobalMiddlewares(handler httputils.APIFunc) httputils.APIFunc {
	next := handler

	for _, m := range s.middlewares {
		next = m.WrapHandler(next)
	}

	if log.GetLevel() == log.DebugLevel {
		next = middleware.DebugRequestMiddleware(next)
	}

	return next
}
