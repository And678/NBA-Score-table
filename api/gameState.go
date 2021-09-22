package api

type GameState struct {
    Time int `json:"time"`
    Game Game `json:"game"`
}