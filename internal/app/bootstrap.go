package app

import (
	v1 "SJTU-Canteen-Community/internal/api/v1"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Start() {
	err := LoadConfig()
	if err != nil {
		log.Errorf("Failed to load config: %v", err)
		return
	}

	err = InitAll()
	if err != nil {
		log.Errorf("Failed to initialize application: %v", err)
		return
	}

	r := gin.Default()
	v1.SetupRoutes(r, DB)
	if err != nil {
		log.Errorf("Failed to set up routes: %v", err)
		return
	}

	err = r.Run(fmt.Sprintf(":%d", Conf.App.Port))
	if err != nil {
		log.Errorf("Failed to start server: %v", err)
		return
	}

}
