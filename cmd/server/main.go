package main

import (
    "backend/internal/genInitialData"
    "backend/internal/live"
    "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/jackc/pgx/v4/stdlib"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)


func initDB() (*sql.DB, error) {
    return sql.Open("pgx", "postgres://dev_user:dev_pwd@localhost:5432/nba?sslmode=disable")
}


func main() {
    db, err := initDB()
    if err != nil {
        panic(err)
    }

    driver, err := postgres.WithInstance(db, &postgres.Config{})
    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        "postgres", driver)
    if err != nil {
        panic(err)
    }
    m.Up()

    genInitialData.GenerateInitialData(db)

    r := gin.Default()

    live.Route(r, db)

    r.LoadHTMLFiles("web/index.html")
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })
    r.Run(":1111")
}
