package entity

import "database/sql"

type Home struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Logo        sql.NullString `json:"logo"`
	Favicon     sql.NullString `json:"favicon"`
	Phone       string         `json:"phone"`
	Email       string         `json:"email"`
}
