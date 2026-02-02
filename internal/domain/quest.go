package domain

import (
	"fmt"
	"time"
)

// Quest — основной объект квеста
type Quest struct {
	ID       string       // уникальный идентификатор
	Title    string       // название квеста
	Intro    string       // вводный текст
	Tasks    []Task       // список заданий
	Location Location     // исходная точка квеста
	Signals  []UserSignal // действия пользователя (ответы, оценки, комментарии)
}

// Task — отдельное задание в квесте
type Task struct {
	ID         int
	Text       string
	Hint       string
	Difficulty TaskDifficulty
	Answer     TaskAnswer
}

type TaskDifficulty struct {
	Level int    // 1–5 (или 1–10)
	Label string // easy | medium | hard | expert
}

type TaskAnswer struct {
	Type       string   // choice | text | photo | offline
	Expected   string   // ожидаемое значение (если применимо)
	Options    []string // варианты ответа (для choice)
	Evaluation string   // стратегия проверки: exact | fuzzy | ai | manual
}

// Location — географическая точка и её описание
type Location struct {
	Latitude  float64 // широта
	Longitude float64 // долгота
	Name      string  // название места
	Type      string  // тип местности (парк, исторический центр, туристическая зона)
}

// UserSignal — сигнал от пользователя (ответ, оценка, комментарий)
type UserSignal struct {
	UserID    string    // уникальный идентификатор пользователя
	TaskID    int       // к какому заданию относится сигнал
	Type      string    // тип сигнала: "answer", "rating", "comment"
	Value     string    // значение сигнала
	Timestamp time.Time // время сигнала
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
