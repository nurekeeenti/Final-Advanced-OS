package main

import (
    "context"
    // "log"
    "github.com/gin-gonic/gin"
    "check/database"
    "check/routes"
)

func main() {
    // Подключение к БД
    db := database.ConnectDB()
    defer db.Client().Disconnect(context.Background())

    r := gin.Default()
    r.Static("/static", "./static")

    // Подключение маршрутов
    routes.AuthRoutes(r, db)
    routes.ProductRoutes(r, db)
    routes.CartRoutes(r, db)

    

    r.GET("/search", func(c *gin.Context) {
        routes.SearchProducts(c, db)
    })
    
    r.GET("/cart", func(c *gin.Context) {
        c.File("./static/cart.html") // Загружает cart.html
    })


    r.Run(":8080")
}
