package usecase

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parent/repository"
	"testing"
)

func TestParentUsecase_ViewGrade(t *testing.T) {
	gormdb, _ := database.GormConfig()
	paretRepo := repository.NewGormParentRepository(gormdb)
	v := NewParentUsecase(paretRepo)
	student := models.Student{Id: 1}
	result, err := v.ViewGrade(student)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}