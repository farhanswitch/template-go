package authors

import (
	"template/models"
	repoPostgres "template/repositories/postgresql"
	"template/utilities"
)

type authorService struct {
	repo repoPostgres.AuthorPostgresRepo
}

var service authorService

func (s authorService) createAuthorSrvc(param models.CreateAuthorRequest) error {
	uuid, err := utilities.GenerateUUIDv7()
	if err != nil {
		return err
	}
	param.UUID = uuid
	return s.repo.CreateAuthor(param)
}

func factoryAuthorService(repo repoPostgres.AuthorPostgresRepo) authorService {
	if service == (authorService{}) {
		service = authorService{
			repo,
		}
	}
	return service
}
