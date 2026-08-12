package domain

import (
	"fmt"
	"time"
)

// TaskDifficulty хранит уровень сложности задания
type TaskDifficulty struct {
	Level int    // 1–5
	Label string // easy | medium | hard
}

// TaskAnswer хранит ожидаемый ответ на задание, тип и стратегию проверки
type TaskAnswer struct {
	Type       string   // choice | text | photo | offline
	Expected   string   // ожидаемый ответ
	Options    []string // варианты для тестовых заданий
	Evaluation string   // exact | fuzzy | ai | manual
}

// Task представляет задание и привязку к месту
type Task struct {
	ID         int
	Text       string
	Hint       string
	Difficulty TaskDifficulty
	Answer     TaskAnswer
	LocationID int // ссылка на место, где выполняется задание
}

// Location хранит информацию о месте (POI)
type Location struct {
	ID        int
	Name      string
	Latitude  float64
	Longitude float64
	Type      string // исторический, развлекательный, природный и т.д.
}

// Quest представляет квест с заданиями, локацией и рейтингом
type Quest struct {
	ID       string
	Title    string
	Intro    string
	Tasks    []Task
	Location Location
	Signals  []UserSignal
	Rating   float64 // средний рейтинг квеста
}

// UserSignal фиксирует действия пользователя
type UserSignal struct {
	UserID    string
	TaskID    int
	Answer    string
	Correct   bool
	Timestamp time.Time
}

// UserProgress хранит прогресс пользователя по квесту
type UserProgress struct {
	UserID    string
	QuestID   string
	Completed bool
	Score     float64
	Timestamp time.Time
}

// AddTask добавляет новое задание в квест
func (q *Quest) AddTask(t Task) {
	q.Tasks = append(q.Tasks, t)
}

// AddSignal добавляет сигнал пользователя к квесту
func (q *Quest) AddSignal(s UserSignal) {
	q.Signals = append(q.Signals, s)
}

// Validate проверяет, что квест корректен
func (q *Quest) Validate() error {
	if q.Title == "" {
		return fmt.Errorf("quest title is empty")
	}
	if len(q.Tasks) == 0 {
		return fmt.Errorf("quest has no tasks")
	}
	return nil
}
