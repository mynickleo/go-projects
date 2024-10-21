package tasks

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateTask(task *Task) error {
	return r.DB.Create(task).Error
}

func (r *Repository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *Repository) GetTaskByID(id uint) (Task, error) {
	var task Task
	err := r.DB.First(&task, id).Error
	return task, err
}

func (r *Repository) UpdateTask(task *Task) error {
	return r.DB.Save(task).Error
}

func (r *Repository) DeleteTask(id uint) error {
	return r.DB.Delete(&Task{}, id).Error
}

func NewRepository() (*Repository, error) {
	db, err := gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Task{})
	return &Repository{DB: db}, nil
}
