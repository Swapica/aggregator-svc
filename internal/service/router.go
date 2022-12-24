package service

import (
	"github.com/Swapica/aggregator-svc/internal/data/mem"
	"github.com/Swapica/aggregator-svc/internal/service/handlers"
	"github.com/Swapica/aggregator-svc/internal/service/helpers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxNoder(mem.NewNodesQ(s.cfg.Nodes())),
		),
	)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/*", handlers.GetHandler)
		r.Post("/*", handlers.PostHandler)
	})

	return r
}
