package main

import (
	"github.com/joho/godotenv"
	"github.com/sumayu/pet/config"
	"github.com/sumayu/pet/internal/bd"
	"github.com/sumayu/pet/internal/logger"
	"github.com/sumayu/pet/internal/router"
)

func main()  {
	godotenv.Load(config.GetConfigPath())
	logger.Init()
	defer logger.Sync()
logger.Info("Приложение запущено",	)
bd.BD()
router.Router().Run()
	}