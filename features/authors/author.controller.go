package authors

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"template/models"
	repoPostgres "template/repositories/postgresql"

	"github.com/go-playground/validator/v10"
)

type authorController struct {
	service authorService
}

var controller authorController

func (a authorController) CreateAuthorCtrl(w http.ResponseWriter, r *http.Request) {
	var param models.CreateAuthorRequest
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"message":"Invalid Data"}`)
		return
	}
	val := validator.New()
	err = val.Struct(param)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"message":"Invalid Data"}`)
		return
	}
	err = a.service.createAuthorSrvc(param)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message":"Something Went Wrong. Please contact our administrator!"}`)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message":"Author created successfully"}`)

}

func factoryAuthorController(repo repoPostgres.AuthorPostgresRepo) authorController {
	if controller == (authorController{}) {
		controller = authorController{
			service: factoryAuthorService(repo),
		}
	}
	return controller
}
