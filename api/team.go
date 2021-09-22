package api

import (
    "github.com/jackc/pgtype"
)

type Team struct {
    ID          pgtype.UUID `json:"-"`
    Name        string      `json:"name"`
}