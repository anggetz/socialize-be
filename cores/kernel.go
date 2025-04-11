package cores

import (
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/websocket/v2"
)

var PGDB *pg.DB

var TopicClients map[int64]map[*websocket.Conn]bool
