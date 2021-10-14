package controllers

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/karan-vk/mongo-crud/config"
	"github.com/karan-vk/mongo-crud/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User

	filter := bson.M{}
	findOptions := options.Find()

	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"Name": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
				{
					"user": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limitVal, _ := strconv.Atoi(c.Query("limit", "10"))
	var limit int64 = int64(limitVal)

	total, _ := userCollection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page) - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, err := userCollection.Find(ctx, filter, findOptions)
	if err != nil {
		defer cursor.Close(ctx)
	}

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Users Not found",
			"error":   err,
		})
	}

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	last := int32(float64(total / limit))
	if last < 1 && total > 0 {
		last = 1
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":      users,
		"total":     total,
		"page":      page,
		"last_page": last,
		"limit":     limit,
	})
}

func GetUser(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var user models.User
	objId, _ := primitive.ObjectIDFromHex(c.Params("id"))
	findResult := userCollection.FindOne(ctx, bson.M{"_id": objId})
	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}

	err := findResult.Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"success": true,
	})
}

func AddUser(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "User failed to insert",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "User inserted successfully",
	})

}

func UpdateUser(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error":   err,
		})
	}

	update := bson.M{
		"$set": user,
	}
	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error":   err,
		})
	}
	_, err = userCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "User failed to delete",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully",
	})
}
