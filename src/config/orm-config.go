package config

import (
	"fmt"
	"itw-be/src/database/entities"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	err    error
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string
	dbType string
)

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USERNAME")
	dbPass = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
}

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	} else {
		log.Println("Connected to database successfully!")
	}
	err = AutoMigrate(DB)
	if err != nil {
		log.Fatal("Failed to perform auto migration. \n", err)
	}
}

func AutoMigrate(connection *gorm.DB) error {
	tx := connection.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	if err := tx.Debug().AutoMigrate(
		&entities.AcademicYear{},
		&entities.AdminBusiness{},
		&entities.AdminSchool{},
		&entities.AppliedHistory{},
		&entities.Business{},
		&entities.Class{},
		&entities.Conversation{},
		&entities.Department{},
		&entities.DetailConversationEnity{},
		&entities.ExaminationBoard{},
		&entities.InternJob{},
		&entities.InternSubject{},
		&entities.Job{},
		&entities.JobLevel{},
		&entities.JobSkill{},
		&entities.Major{},
		&entities.Position{},
		&entities.Program{},
		&entities.RegularTodoDetail{},
		&entities.RegularTodo{},
		&entities.Report{},
		&entities.Role{},
		&entities.School{},
		&entities.Semester{},
		&entities.Skill{},
		&entities.Student{},
		&entities.StudentLearnIntern{},
		&entities.Teacher{},
		&entities.TodoAppreciation{},
		&entities.UserAccount{},
		&entities.UserPerson{},
	); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to auto migrate: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Println("Auto migration completed successfully!")
	return nil
}
