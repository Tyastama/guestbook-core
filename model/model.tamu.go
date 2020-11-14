package model

import (
	"context"
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	Tamu struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	}

	TamuResponse struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Email       string    `json:"email"`
		Date        time.Time `json:"date"`
		Discription string    `json:"discription"`
	}

	AllTamu struct {
		FilterBy    string `json:"filter_by"`
		FilterValue string `json:"filter_value"`
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (u *Tamu) Response() TamuResponse {
	return TamuResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Date:  u.Date,
	}
}

func (u *Tamu) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {

}

func (u *Tamu) All(ctx context.Context, db *sql.DB, param AllTamu) ([]Tamu, error) {

}

func (u *Tamu) One(ctx context.Context, db *sql.DB) (Tamu, error) {

}

func (u *Tamu) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {

}

func (u *Tamu) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {

}
