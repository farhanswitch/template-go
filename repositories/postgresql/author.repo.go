package postgresql

import (
	"template/connections"
	"template/models"
	"template/utilities"
)

type AuthorPostgresRepo struct{}

var repo AuthorPostgresRepo

func (a AuthorPostgresRepo) CreateAuthor(param models.CreateAuthorRequest) error {
	uuid, err := utilities.GenerateUUIDv7()
	if err != nil {
		return err
	}
	_, err = connections.DbPostgres().Query("INSERT INTO public.authors(id, name) VALUES($1,$2);", uuid, param.Name)
	return err
}
func FactoryAuthorPostgresRepo() AuthorPostgresRepo {
	if repo == (AuthorPostgresRepo{}) {
		repo = AuthorPostgresRepo{}
	}
	return repo
}
