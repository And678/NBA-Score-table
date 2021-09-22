package simulator

import (
    "backend/api"
    "math/rand"
    "time"
)

type Simulator struct {
    time int
    game  *api.Game
}

func SimulateGame(game *api.Game, timeScale float64) *Simulator {
    sim := Simulator{
        time:  0,
        game: game,
    }
    t := time.Now()
    game.StartsAt = &t
    go func() {
        for {
            if sim.time >= 48 * 60 {
                tFinish := t.Add(time.Second * 48 * 60)
                game.FinishedAt = &tFinish
                return
            }
            time.Sleep(time.Duration(float64(time.Second) * timeScale))
            p := rand.Float64()

            if p < 0.004 {
                game.ScoreA += 1
                if p > 0.002 {
                    game.Assists += 1
                }
            } else if p >= 0.004 && p < 0.008 {
                game.ScoreB += 1
                if p > 0.006 {
                    game.Assists += 1
                }
            }
            if p < 0.016 {
                game.Attacks += 1
            }
            sim.time += 1

        }
    }()
    return &sim
}

func (sim *Simulator) GetTime() int {
    return sim.time
}

func (sim *Simulator) GetGame() api.Game {
    return *sim.game
}