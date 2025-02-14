package routes

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"


    "time"
    "log"
)

func ProductRoutes(r *gin.Engine, db *mongo.Database) {
    productsCollection := db.Collection("Products")

    r.GET("/api/products", func(c *gin.Context) {
        pipeline := mongo.Pipeline{
            bson.D{{"$lookup", bson.D{{"from", "Categories"}, {"localField", "category"}, {"foreignField", "_id"}, {"as", "categoryDetails"}}}},
            bson.D{{"$lookup", bson.D{{"from", "Brands"}, {"localField", "brand"}, {"foreignField", "_id"}, {"as", "brandDetails"}}}},
        }

        cursor, err := productsCollection.Aggregate(context.Background(), pipeline)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
            return
        }
        defer cursor.Close(context.Background())

        var products []bson.M
        if err = cursor.All(context.Background(), &products); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse products"})
            return
        }

        c.JSON(http.StatusOK, products)
    })

    r.GET("/dashboard", func(c *gin.Context) {
        c.File("./static/dashboard.html")
    })
}




func SearchProducts(c *gin.Context, db *mongo.Database) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Введите поисковый запрос"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productCollection := db.Collection("Products")

	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"$text", bson.D{{"$search", query}}}}}},
		bson.D{{"$lookup", bson.D{{"from", "Brands"}, {"localField", "brand"}, {"foreignField", "_id"}, {"as", "brandInfo"}}}},
		bson.D{{"$lookup", bson.D{{"from", "Categories"}, {"localField", "category"}, {"foreignField", "_id"}, {"as", "categoryInfo"}}}},
		bson.D{{"$unwind", bson.D{{"path", "$brandInfo"}, {"preserveNullAndEmptyArrays", true}}}},
		bson.D{{"$unwind", bson.D{{"path", "$categoryInfo"}, {"preserveNullAndEmptyArrays", true}}}},
		bson.D{{"$project", bson.D{
			{"name", 1},
			{"price", 1},
			{"brand", "$brandInfo.name"},
			{"category", "$categoryInfo.name"},
		}}},
	}

	cursor, err := productCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println("Ошибка поиска:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка поиска"})
		return
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		log.Println("Ошибка обработки результатов:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки данных"})
		return
	}

	c.JSON(http.StatusOK, results)
}












