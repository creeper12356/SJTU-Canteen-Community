package app

import (
	v1_b "SJTU-Canteen-Community/internal/api/v1/b"
	v1_c "SJTU-Canteen-Community/internal/api/v1/c"
	"fmt"

	"github.com/gin-contrib/sessions"
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

	r.Use(sessions.Sessions("JSESSIONID", RedisStore))

	v1_c.SetupCRoutes(r, DB)
	v1_b.SetupBRoutes(r, DB)

	err = r.Run(fmt.Sprintf(":%d", Conf.App.Port))
	if err != nil {
		log.Errorf("Failed to start server: %v", err)
		return
	}

}
