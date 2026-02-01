package service

import "geo_quest/internal/domain"

// QuestGenerator — интерфейс генератора квестов
type QuestGenerator interface {
	GenerateQuest(lat, lng float64, style, difficulty string) (*domain.Quest, error)
}
