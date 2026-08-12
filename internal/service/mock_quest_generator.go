package service

import (
	"fmt"
	"geo_quest/internal/domain"
	"time"
)

// MockQuestGenerator — пример генератора, который возвращает фиксированный квест
type MockQuestGenerator struct{}

// GenerateQuest возвращает Quest на основе координат и параметров
func (g *MockQuestGenerator) GenerateQuest(lat, lng float64, style, difficulty string) (*domain.Quest, error) {
	// Создаем локацию
	location := domain.Location{
		ID:        1,
		Name:      "Старый город",
		Latitude:  lat,
		Longitude: lng,
		Type:      "historical_center",
	}

	// Создаем квест
	q := &domain.Quest{
		ID:       fmt.Sprintf("qst_%d", time.Now().Unix()),
		Title:    "Тайна старых дворов",
		Intro:    fmt.Sprintf("Вы находитесь рядом с координатами %.4f, %.4f. Исследуйте окрестности!", lat, lng),
		Location: location,
		Rating:   4.5, // пример рейтинга
	}

	// Мок-задания с TaskAnswer и привязкой к LocationID
	task1 := domain.Task{
		ID:         1,
		Text:       "Найди старую арку рядом",
		Hint:       "Она старше 100 лет",
		Difficulty: domain.TaskDifficulty{Level: 2, Label: "medium"},
		Answer: domain.TaskAnswer{
			Type:       "text",
			Expected:   "Арка найдена",
			Evaluation: "exact",
		},
		LocationID: location.ID,
	}

	task2 := domain.Task{
		ID:         2,
		Text:       "Сфотографируй памятник",
		Hint:       "Он находится на площади",
		Difficulty: domain.TaskDifficulty{Level: 3, Label: "hard"},
		Answer: domain.TaskAnswer{
			Type:       "photo",
			Evaluation: "ai",
		},
		LocationID: location.ID,
	}

	q.Tasks = append(q.Tasks, task1, task2)

	// Пример пользовательского сигнала
	signal := domain.UserSignal{
		UserID:    "u_001",
		TaskID:    task1.ID,
		Answer:    "Найдено",
		Correct:   true,
		Timestamp: time.Now(),
	}

	q.Signals = append(q.Signals, signal)

	return q, nil
}
