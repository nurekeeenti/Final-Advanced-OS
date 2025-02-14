package routes

import (
    "context"
    "net/http"
 	"check/models"
    "check/utils"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func AuthRoutes(r *gin.Engine, db *mongo.Database) {
    usersCollection := db.Collection("users")

    r.GET("/", func(c *gin.Context) {
        c.File("./static/login.html")
    })

    r.POST("/login", func(c *gin.Context) {
        var user models.User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }

        var dbUser models.User
        err := usersCollection.FindOne(context.Background(), bson.M{"email": user.Email, "password": user.Password}).Decode(&dbUser)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Login successful", "redirect": "/dashboard"})
    })

    r.POST("/register", func(c *gin.Context) {
        var user models.User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }

        utils.SendVerificationCode(user.Email)
        c.JSON(http.StatusOK, gin.H{"message": "Verification code sent"})
    })

    r.POST("/verify", func(c *gin.Context) {
        var data struct {
            Email    string `json:"email"`
            Code     string `json:"code"`
            Name     string `json:"name"`
            Password string `json:"password"`
        }

        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }

        if utils.VerifyCode(data.Email, data.Code) {
            newUser := models.User{Name: data.Name, Email: data.Email, Password: data.Password}
            _, err := usersCollection.InsertOne(context.Background(), newUser)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
                return
            }
            utils.RemoveCode(data.Email)
            c.JSON(http.StatusOK, gin.H{"message": "Account created successfully"})
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid verification code"})
        }
    })
}
