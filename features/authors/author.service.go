package authors

import (
	"net/http"
	"template/models"
	repoPostgres "template/repositories/postgresql"
	"template/utilities"
	errUtility "template/utilities/errors"
)

type authorService struct {
	repo repoPostgres.AuthorPostgresRepo
}

var service authorService

func (s authorService) createAuthorPostgresSrvc(param models.CreateAuthorRequest) (bool, errUtility.CustomError) {
	uuid, err := utilities.GenerateUUIDv7()
	if err != nil {
		return true, errUtility.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	param.UUID = uuid
	return s.repo.CreateAuthor(param)
}
func (s authorService) getAuthorByIDPostgresSrvc(authorID string) (models.Author, errUtility.CustomError) {
	return s.repo.GetAuthorByID(authorID)
}
func factoryAuthorService(repo repoPostgres.AuthorPostgresRepo) authorService {
	if service == (authorService{}) {
		service = authorService{
			repo,
		}
	}
	return service
}
