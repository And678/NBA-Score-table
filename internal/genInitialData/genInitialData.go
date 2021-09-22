package genInitialData

import (
    "database/sql"
    "fmt"
    "github.com/jackc/pgtype"
    _ "github.com/jackc/pgx/v4/stdlib"
)


// Placeholder hack to populate database with fake teams and games
func GenerateInitialData(db *sql.DB) {
    var teams []string

    db.Exec("TRUNCATE TABLE game")
    db.Exec("DELETE FROM team")
    for i := 0; i < 30; i++ {
        id := pgtype.UUID{}
        err := db.QueryRow(fmt.Sprintf(`INSERT INTO team(name) VALUES ('Team %d') RETURNING id;`, i + 1)).Scan(&id)
        if err != nil {
            println(err.Error())
        }
        idStr, _ := id.EncodeText(nil, nil)
        teams = append(teams, string(idStr))
    }
    for i := 0; i < len(teams); i += 2 {
        _, err := db.Exec(fmt.Sprintf(`INSERT INTO game(location, team_a, team_b) VALUES ('Location %d', '%s', '%s') RETURNING id;`, i + 1, teams[i], teams[i + 1]))
        if err != nil {
            println(err.Error())
        }
    }

}