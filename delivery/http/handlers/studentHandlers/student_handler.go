package studentHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices"
	"html/template"
	"net/http"
	"strconv"
)

type StudentHandler struct {
	templ    *template.Template
	Session  session.SessionUsecase
	SUsecase studentServices.StudentUsecase
}

func NewStudentHandler(templ *template.Template, session session.SessionUsecase) *StudentHandler {
	return &StudentHandler{templ: templ, Session: session}
}

type StudentInfo struct {
	User          models.User
	Tasks         []models.Task
	Task          models.Task
	UpdateProfile models.Student
	Resources     []models.Resources
	Result        []models.Result
	ClassMates    []models.Student
}

func (sh *StudentHandler) AddStudent(w http.ResponseWriter, r *http.Request) {

}

func (sh *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {

}

func (sh *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {

}

func (sh *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)

	email := r.FormValue("studentEmail")
	password := r.FormValue("studentPassword")
	profile := r.FormValue("studentProfilePic")

	in := StudentInfo{
		User:          user,
		UpdateProfile: models.Student{Email: user.Email, Password: user.Password},
	}

	if email != "" && password != "" && profile != "" {
		user.Email = email
		user.Password = password
		studentUpdateInfo := models.Student{Email: email, Password: password, ProfilePic: profile}
		newStudent, errs := sh.SUsecase.UpdateStudent(studentUpdateInfo)
		if len(errs) > 0 {
			fmt.Println(errs)
		}
		in = StudentInfo{
			User:          user,
			UpdateProfile: newStudent,
		}
	}
	err := sh.templ.ExecuteTemplate(w, "studentUpdateProfile", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (sh *StudentHandler) ViewTasks(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	subjectId, err := strconv.Atoi(r.FormValue("id"))
	//input validation
	student, _ := sh.SUsecase.GetStudentById(user.Id)
	data, _ := sh.SUsecase.ViewTasks(student.SectionId, uint(subjectId))
	if err != nil {
		fmt.Println(err)
	}
	in := StudentInfo{
		Tasks: data,
		Task:  models.Task{},
		User:  user,
	}
	err = sh.templ.ExecuteTemplate(w, "studentViewTask", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (sh *StudentHandler) Comment(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)

	comment := r.FormValue("comment")
	taskId, _ := strconv.Atoi(r.FormValue("taskId"))
	//input validation
	_ = sh.SUsecase.Comment(uint(taskId), user.Id, comment)
	http.Redirect(w, r, "/student/viewTask", http.StatusSeeOther)
}

func (sh *StudentHandler) ViewClass(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	student, _ := sh.SUsecase.GetStudentById(user.Id)
	classMates, _ := sh.SUsecase.ViewClass(student.ClassRoomId)
	in := StudentInfo{
		User:       user,
		ClassMates: classMates,
	}
	err := sh.templ.ExecuteTemplate(w, "studentClassMates", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (sh *StudentHandler) ViewResources(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	subjectId, _ := strconv.Atoi(r.FormValue("subjectId"))
	resources, errs := sh.SUsecase.ViewResources(uint(subjectId))
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	in := StudentInfo{
		User:      user,
		Resources: resources,
	}
	err := sh.templ.ExecuteTemplate(w, "studentResources", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (sh *StudentHandler) ViewResult(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	results, _ := sh.SUsecase.ViewResult(user.Id)
	in := StudentInfo{
		User:   user,
		Result: results.Result,
	}
	err := sh.templ.ExecuteTemplate(w, "studentViewResult", in)
	if err != nil {
		fmt.Println(err)
	}
}