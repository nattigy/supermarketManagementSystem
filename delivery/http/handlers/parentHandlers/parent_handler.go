package parentHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session/usecase"
	"html/template"
	"net/http"
)

type ParentHandler struct {
	templ   *template.Template
	Session usecase.SessionUsecase
}

func NewParentHandler(templ *template.Template, session usecase.SessionUsecase) *ParentHandler {
	return &ParentHandler{templ: templ, Session: session}
}

type ParentInfo struct {
	User   models.User
	Result []models.Result
}

func (ph *ParentHandler) AddParent(w http.ResponseWriter, r *http.Request) {

}

func (ph *ParentHandler) GetParents(w http.ResponseWriter, r *http.Request) {

}

func (ph *ParentHandler) DeleteParent(w http.ResponseWriter, r *http.Request) {

}

func (ph *ParentHandler) ViewGrade(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	in := ParentInfo{
		User:   models.User{Id: sess.ID, Email: sess.Email, Role: sess.Role, LoggedIn: true},
		Result: []models.Result{},
	}
	err := ph.templ.ExecuteTemplate(w, "parentViewResult.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}
