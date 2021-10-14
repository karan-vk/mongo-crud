package controllers

import (
	"context"
	"time"

	"github.com/karan-vk/mongo-crud/config"
	"github.com/karan-vk/mongo-crud/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetToken(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var auth models.Auth
	c.BodyParser(&auth)
	y, _ := primitive.ObjectIDFromHex(auth.Id)
	findResult := userCollection.FindOne(ctx, bson.M{"_id": y})

	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	var user models.User
	if err := findResult.Decode(&user); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	claims := jwt.RegisteredClaims{
		ID: user.ID.Hex(),
		IssuedAt: &jwt.NumericDate{
			Time: time.Now(),
		},
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(time.Hour * 24),
		},
		Audience: jwt.ClaimStrings{
			user.Name,
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.MI.JwtSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":   t,
		"user":    user,
		"success": true,
	})
}
