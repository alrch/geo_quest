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
	q := &domain.Quest{
		ID:    fmt.Sprintf("qst_%d", time.Now().Unix()),
		Title: "Тайна старых дворов",
		Intro: fmt.Sprintf("Вы находитесь рядом с координатами %.4f, %.4f. Исследуйте окрестности!", lat, lng),
		Location: domain.Location{
			Latitude:  lat,
			Longitude: lng,
			Name:      "Старый город",
			Type:      "historical_center",
		},
	}

	// Мок-задания
	q.AddTask(domain.Task{ID: 1, Text: "Найди старую арку рядом", Hint: "Она старше 100 лет"})
	q.AddTask(domain.Task{ID: 2, Text: "Сфотографируй памятник", Hint: "Он находится на площади"})

	// Пример пользовательского сигнала
	q.AddSignal(domain.UserSignal{
		UserID:    "u_001",
		TaskID:    1,
		Type:      "answer",
		Value:     "Найдено",
		Timestamp: time.Now(),
	})

	return q, nil
}
