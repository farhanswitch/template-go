package postgresql

import (
	"database/sql"
	"net/http"
	"strings"
	"template/connections"
	"template/models"
	utility "template/utilities"
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
func (AuthorPostgresRepo) GetAuthorByID(authorID string) (models.Author, errUtility.CustomError) {
	var author models.Author
	var nullableUpdatedAt sql.NullTime
	err := connections.DbPostgres().QueryRow("SELECT id, name, created, updated FROM public.authors WHERE id=$1;", authorID).Scan(&author.ID, &author.Name, &author.Created, &nullableUpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Author{}, errUtility.CustomError{
				Code:          http.StatusNotFound,
				Message:       "Author not found",
				MessageToSend: "Author not found",
			}
		} else if strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			return models.Author{}, errUtility.CustomError{
				Code:          http.StatusNotFound,
				Message:       "Invalid UUID",
				MessageToSend: "Author not found",
			}
		}

		return models.Author{}, errUtility.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	author.StrCreated = utility.FormatTimeMillis(author.Created)
	if nullableUpdatedAt.Valid {
		author.Updated = nullableUpdatedAt.Time
		author.StrUpdated = utility.FormatTimeMillis(nullableUpdatedAt.Time)
	} else {
		author.Updated = author.Created
		author.StrUpdated = utility.FormatTimeMillis(author.Created)
	}
	return author, errUtility.CustomError{}
}
func FactoryAuthorPostgresRepo() AuthorPostgresRepo {
	if repo == (AuthorPostgresRepo{}) {
		repo = AuthorPostgresRepo{}
	}
	return repo
}
