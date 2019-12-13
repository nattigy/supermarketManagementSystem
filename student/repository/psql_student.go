package repository

import (
	"database/sql"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type PsqlStudentRepository struct {
	conn *sql.DB
}

func NewPsqlStudentRepository(Conn *sql.DB) *PsqlStudentRepository {
	return &PsqlStudentRepository{conn: Conn}
}

func (ps *PsqlStudentRepository) ViewTasks(c models.ClassRoom, s models.Subject) ([]models.Task, error) {
	return nil, nil
}

func (ps *PsqlStudentRepository) Comment(t models.Task, student models.Student) error {
	return nil
}

func (ps *PsqlStudentRepository) StudentUpdateProfile(student models.Student) error {
	return nil
}

func (ps *PsqlStudentRepository) ViewClass(classRoom models.ClassRoom) ([]models.Student, error) {
	return nil, nil
}

func (ps *PsqlStudentRepository) ViewResources(subject models.Subject) ([]models.Resource, error) {
	return nil, nil
}

func (ps *PsqlStudentRepository) ViewResult(s models.Student) ([]models.Result, error)  {
	return nil, nil
}