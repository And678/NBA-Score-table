# NBA Score table

Live score table built using websockets. Simulation starts on first websocket connection. 

## Running

To start the database run `docker-compose up`, then run `go mod download` and `go run cmd/server/main.go` to start the project. 
After that, open `localhost:1111` in any web browser.

## Used technologies

* Postgresql
* Sqrl - https://github.com/elgris/sqrl
* Gin - https://github.com/gin-gonic/gin
* Gorilla-websocket - https://github.com/gorilla/websocket
* golang-migrate - https://github.com/golang-migrate/migrate
