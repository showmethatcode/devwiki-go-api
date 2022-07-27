package main

import (
	"context"
	"devwiki/controllers"
	"devwiki/db"
	"devwiki/ent/migrate"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	d, err := db.Conn()

	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	err = d.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.POST("/term/write", func(c *gin.Context) {
		err = controllers.CreateTerm(c)
		//if err != nil {
		//	return
		//}
	})

	if err = r.Run(":3000"); err != nil {
		log.Fatalln(err)
	}
}
