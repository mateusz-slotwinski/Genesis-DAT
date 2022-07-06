package src

import (
	Config "GenesisDAT/src/config"
	Router "GenesisDAT/src/router"
)

func App() {
	App := Router.CreateServer()

	App.Run(Config.Port)
}

func Init() {
	Config.Init()

	Logs()
	App()

}
