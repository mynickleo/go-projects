package tasks

type Task struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
