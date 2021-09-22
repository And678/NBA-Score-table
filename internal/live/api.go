package live

import (
    "backend/api"
    "backend/internal/fixture"
    simulator "backend/internal/simulation"
    "database/sql"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "log"
    "time"
)

var upgrader = websocket.Upgrader{}
const timeScale float64 = 5.0/60.0
var simulations []*simulator.Simulator
var currentGames []*api.Game

func SimulateGames(games []*api.Game) {
    currentGames = make([]*api.Game, 0, len(games))
    for _, v := range games {
        sim := simulator.SimulateGame(v, timeScale)
        simulations = append(simulations, sim)
        currentGames = append(currentGames, v)
    }
}

func Route(router gin.IRouter, db *sql.DB) {
    router.GET("/live", func(context *gin.Context) {
        fixturesDb := fixture.NewRepository(db)

        fix, err := fixturesDb.GetFullFixture()
        if err != nil {
            panic(err)
        }

        c, err := upgrader.Upgrade(context.Writer, context.Request, nil)
        if err != nil {
            log.Print("upgrade:", err)
            return
        }
        defer c.Close()

        if len(currentGames) == 0 {
            SimulateGames(fix.Games)
        }

        for {
            time.Sleep(time.Second)
            var gameState = make([]api.GameState, 0, len(simulations))
            for _, v := range simulations {
                gameState = append(gameState, api.GameState{
                    Time: v.GetTime(),
                    Game: v.GetGame(),
                })
            }
            message, err := json.Marshal(gameState)
            if err != nil {
                log.Print(err.Error())
                break
            }
            err = c.WriteMessage(websocket.TextMessage, message)
        }
    })
}
