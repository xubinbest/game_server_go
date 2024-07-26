package main

import (
	"game_server_go/src/server/conf"
	"game_server_go/src/server/game"
	"game_server_go/src/server/gate"
	"game_server_go/src/server/login"
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)

}
