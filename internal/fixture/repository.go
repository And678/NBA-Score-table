package fixture

import (
    "backend/api"
    "database/sql"
    sq "github.com/elgris/sqrl"
)

type Repository interface {
    GetFullFixture() (*api.Fixture, error)
    UpdateGame(*api.Game) error
}

type repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
    return &repository{
        db: db,
    }
}

func (r *repository) UpdateGame(game *api.Game) error {
    q := sq.Update("game").SetMap(
        map[string]interface{}{
            "score_a": game.ScoreA,
            "score_b": game.ScoreB,
            "starts_at": game.StartsAt,
            "finished_at": game.FinishedAt,
        },
    ).Where("id = ?", game.ID)
    _, err := q.RunWith(r.db).Exec()
    return err
}

func (r *repository) GetFullFixture() (*api.Fixture, error) {
    var fixture api.Fixture

    // Games
    q := sq.Select(
        "game.id", "location", "team_a", "team_b", "score_a", "score_b", "starts_at", "finished_at",
        "ta.id", "ta.name", "tb.id", "tb.name",
        ).From("game").
        LeftJoin("team ta ON team_a = ta.id").
        LeftJoin("team tb ON team_b = tb.id")
    rows, err := q.RunWith(r.db).Query()
    if err != nil {
        return nil, err
    }
    for rows.Next() {
        game, err := scanGame(rows)
        if err != nil {
            return nil, err
        }
        fixture.Games = append(fixture.Games, game)
    }

    return &fixture, nil
}

func scanGame(row sq.RowScanner) (*api.Game, error) {
    var game api.Game
    err := row.Scan(&game.ID, &game.Location, &game.TeamAID, &game.TeamBID,
        &game.ScoreA, &game.ScoreB, &game.StartsAt, &game.FinishedAt,
        &game.TeamA.ID, &game.TeamA.Name,
        &game.TeamB.ID, &game.TeamB.Name,
    )
    return &game, err
}