package main

import (
	"fmt"
	"log"

	"github.com/alaaattya/recipy-admin-api/middlewares"
	"github.com/alaaattya/recipy-admin-api/requestHandlers"
	"github.com/gin-gonic/gin"
	"github.com/go-bongo/bongo"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("yaml")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	}

	config := &bongo.Config{
		ConnectionString: viper.GetString("local.database.host"),
		Database:         viper.GetString("local.database.dbName"),
	}

	connection, err := bongo.Connect(config)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Use(middlewares.Database(connection))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/recipy", requestHandlers.HandleCreateNewRecipy)

	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
