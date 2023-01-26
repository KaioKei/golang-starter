package gin

import (
	"github.com/gin-gonic/gin"
	"log"
)

type image struct {
	ID string ``
}

func run() {
	router := gin.Default()
	log.Println(router)
}
