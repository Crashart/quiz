package entity

import "time"

type Rating struct {
	Name           string        `json:"name"`
	CorrectAnswers int           `json:"correctAnswers"`
	Duration       time.Duration `json:"duration"`
}
