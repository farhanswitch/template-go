package postgresql

import (
	"template/connections"
	"template/models"
)

type AuthorPostgresRepo struct{}

var repo AuthorPostgresRepo

func (a AuthorPostgresRepo) CreateAuthor(param models.CreateAuthorRequest) error {
	uuid := "418ce6c3-7df1-780a-8b76-1e0a4e01a89f"
	_, err := connections.DbPostgres().Query("INSERT INTO public.authors(id, name) VALUES($1,$2);", uuid, param.Name)
	return err
}
func FactoryAuthorPostgresRepo() AuthorPostgresRepo {
	if repo == (AuthorPostgresRepo{}) {
		repo = AuthorPostgresRepo{}
	}
	return repo
}
