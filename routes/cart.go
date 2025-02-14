package routes

import (
    "context"
    "net/http"
    "check/models"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func CartRoutes(r *gin.Engine, db *mongo.Database) {
    cartsCollection := db.Collection("Carts")

    r.GET("/api/cart", func(c *gin.Context) {
        var cart models.Cart
        err := cartsCollection.FindOne(context.Background(), bson.M{}).Decode(&cart)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
            return
        }
        c.JSON(http.StatusOK, cart)
    })

    r.POST("/api/cart/add", func(c *gin.Context) {
        var request struct {
            Name  string `json:"name"`
            Price int    `json:"price"`
        }
        if err := c.BindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }
        
        newProduct := bson.M{
            "name":     request.Name,
            "price":    request.Price,
            "quantity": 1,
        }
        
        result, err := cartsCollection.UpdateOne(
            context.Background(),
            bson.M{},
            bson.M{"$push": bson.M{"products": newProduct}},
        )
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product to cart"})
            return
        }
        if result.ModifiedCount == 0 {
            c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
            return
        }
        
        c.JSON(http.StatusOK, gin.H{"message": "Product added to cart successfully"})
    })

    r.POST("/api/cart/remove", func(c *gin.Context) {
        var request struct {
            Name string `json:"name"`
        }
        if err := c.BindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }
        
        result, err := cartsCollection.UpdateOne(
            context.Background(),
            bson.M{},
            bson.M{"$pull": bson.M{"products": bson.M{"name": request.Name}}},
        )
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item"})
            return
        }
        if result.ModifiedCount == 0 {
            c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
            return
        }
        
        c.JSON(http.StatusOK, gin.H{"message": "Item removed successfully"})
    })

    r.POST("/api/cart/update", func(c *gin.Context) {
        var request struct {
            Name     string `json:"name"`
            Quantity int    `json:"quantity"`
        }
        if err := c.BindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }
        
        result, err := cartsCollection.UpdateOne(
            context.Background(),
            bson.M{"products.name": request.Name},
            bson.M{"$set": bson.M{"products.$.quantity": request.Quantity}},
        )
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update quantity"})
            return
        }
        if result.ModifiedCount == 0 {
            c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
            return
        }
        
        c.JSON(http.StatusOK, gin.H{"message": "Quantity updated successfully"})
    })
}

// package routes

// import (
//     "context"
//     "net/http"
//     "check/models"

//     "github.com/gin-gonic/gin"
//     "go.mongodb.org/mongo-driver/bson"
//     "go.mongodb.org/mongo-driver/mongo"
// )

// func CartRoutes(r *gin.Engine, db *mongo.Database) {
//     cartsCollection := db.Collection("Carts")

//     r.GET("/api/cart", func(c *gin.Context) {
//         var cart models.Cart
//         err := cartsCollection.FindOne(context.Background(), bson.M{}).Decode(&cart)
//         if err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
//             return
//         }
//         c.JSON(http.StatusOK, cart)
//     })
// }