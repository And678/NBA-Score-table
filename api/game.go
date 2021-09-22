package api

import (
    "github.com/jackc/pgtype"
    "time"
)

type Game struct {
    ID          pgtype.UUID `json:"-"`
    Location    string      `json:"location"`
    TeamAID     pgtype.UUID `json:"-"`
    TeamBID     pgtype.UUID `json:"-"`
    TeamA       Team        `json:"teamA"`
    TeamB       Team        `json:"teamB"`
    ScoreA      int         `json:"scoreA"`
    ScoreB      int         `json:"scoreB"`
    Attacks     int         `json:"attacks"`
    Assists     int         `json:"assists"`
    StartsAt    *time.Time  `json:"startsAt,omitempty"`
    FinishedAt  *time.Time  `json:"finishedAt,omitempty"`
}
