package authors

import (
	repoPostgres "template/repositories/postgresql"

	"github.com/go-chi/chi/v5"
)

func InitModule(router *chi.Mux) {
	repo := repoPostgres.FactoryAuthorPostgresRepo()
	controller := factoryAuthorController(repo)

	router.Post("/api/v1/author/pg", controller.CreateAuthorPostgresCtrl)
}
