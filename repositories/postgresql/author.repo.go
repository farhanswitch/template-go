package postgresql

import (
	"net/http"
	"template/connections"
	"template/models"
	errUtility "template/utilities/errors"
)

type AuthorPostgresRepo struct{}

var repo AuthorPostgresRepo

func (a AuthorPostgresRepo) CreateAuthor(param models.CreateAuthorRequest) (bool, errUtility.CustomError) {
	_, err := connections.DbPostgres().Query("INSERT INTO public.authors(id, name) VALUES($1,$2);", param.UUID, param.Name)
	if err != nil {
		return true, errUtility.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return false, errUtility.CustomError{}
}
func FactoryAuthorPostgresRepo() AuthorPostgresRepo {
	if repo == (AuthorPostgresRepo{}) {
		repo = AuthorPostgresRepo{}
	}
	return repo
}
