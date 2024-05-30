package student

import (
	"itw-be/src/config"
	"itw-be/src/database/entities"

	"github.com/google/uuid"
)

func SVGetStudents() ([]entities.Student, error) {
	var students []entities.Student
	err := config.DB.Preload("User").Preload("Class").Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func SVCreateStudent(input StudentInput) (entities.Student, error) {
	student := entities.Student{
		UserID:        uuid.New(),
		AdmissionDate: input.AdmissionDate,
		Dob:           input.Dob,
		Sex:           input.Sex,
		CurrentStatus: input.CurrentStatus,
		ClassID:       input.ClassID,
	}

	err := config.DB.Create(&student).Error
	if err != nil {
		return entities.Student{}, err
	}

	return student, nil
}