package handler

import "time"

type ClassRequest struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	StartDate    time.Time `json:"start_date"`
	GraduateDate time.Time `json:"graduate_date"`
}
