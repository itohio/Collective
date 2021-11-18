package api

//go:generate gqlgen

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"github.com/itohio/collective/backend/pkg/api/graph"
	"github.com/itohio/collective/backend/pkg/api/graph/generated"
	"github.com/itohio/collective/backend/pkg/auth"
	"github.com/itohio/collective/backend/pkg/db"
)

func New(db db.DBType) http.Handler {
	domain := os.Getenv("AUTH0_DOMAIN")
	clientId := os.Getenv("AUTH0_CLIENT_ID")
	authMiddleware := auth.Middleware(clientId, domain)
	router := chi.NewRouter()

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					Orm: db,
				},
				Directives: generated.DirectiveRoot{
					IsAuthenticated: IsAuthenticated,
					HasScope:        HasScope,
				},
			},
		),
	)

	router.Handle("/", authMiddleware.Handler(Gzip(srv)))
	router.HandleFunc("/play", playground.Handler("GraphQL playground", "/api"))

	return router
}
