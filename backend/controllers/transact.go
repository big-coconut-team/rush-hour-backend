package controllers

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"scalable-final-proj/backend/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"	
	"github.com/go-redis/redis"
    "context"  
    "bytes" 
    "fmt"  
)

type CheckStockInput struct {
	ProdID 			int 		`json:"prod_id" binding:"required"`
	Amount 			int			`json:"amount" binding:"required"`
}

func checkStock(c *gin.Context) (bool, error) {

	var input CheckStockInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stock, err := DB.Query("SELECT stock FROM products WHERE prod_id = ?", input.ProdID)

	// if there is an error prod_iding, handle it
	if err != nil {
		panic(err.Error())
	}

	if stock < input.Amount {
		c.JSON(http.StatusOK, gin.H{
			"prod_id": input.ProdID,
			"message": "not enough product left"
		})
		return false, nil
		// disable the "confirm payment" button
	}

	return true, nil
}

type ConfirmInput struct {
	ProdID 			int 		`json:"prod_id" binding:"required"`
	Amount 			int			`json:"amount" binding:"required"`
	Confirmation	bool		`json:"confirm" binding:"required"`
}

type OrderRequest struct {
	ProdID 		int 		
	Amount 		int			
	UserID		int
}

func confirm(c *gin.Context) {

	var input ConfirmInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prod_id := input.ProdID
	amount := input.Amount
	user_id := utils.ExtractTokenID(c)


	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })


	ctx := context.TODO()

	orq =: OrderRequest({prod_id, amount, user_id})

	orq_as_byte =: json.Marshal(orq)

    err := redisClient.RPush(ctx, "payments", ).Err(); 

    if err != nil {
        fmt.Fprintf(w, err.Error() + "\r\n")             
    }                         

	// jsonify(OrderRequest)
	// RPUSH into redisqueue
	// worker stuff

}