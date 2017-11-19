package requestHandlers

import (
	"fmt"

	"github.com/alaaattya/recipy-admin-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-bongo/bongo"
)

// HandleCreateNewRecipy create new recipy handler
func HandleCreateNewRecipy(context *gin.Context) {

	connection, ok := context.Keys["DBConnection"].(*bongo.Connection)
	if !ok {
		fmt.Println("error: database object is not found...") //Handle case of no *DB instance
	}

	recipy := &models.Recipe{
		Title:       "title",
		CategoryID:  "1",
		Description: "super easy delicios meal",
		CookingTime: 20,
	}

	connection.Collection("recipies").Save(recipy)

	context.JSON(200, gin.H{
		"message": "heeeey",
	})
}

func transformRequestParams() {

}
