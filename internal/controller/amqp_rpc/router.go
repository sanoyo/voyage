package amqprpc

import (
	"github.com/sanoyo/voyage/internal/usecase"
	"github.com/sanoyo/voyage/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
