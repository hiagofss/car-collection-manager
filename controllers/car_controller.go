package controllers

import (
	"car-collection-manager/database"
	"car-collection-manager/domain/dto"
	"car-collection-manager/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Insert new Car resource
// @Description Desc
// @Tags car
// @Param request body dto.CarRequest.request true "query params"
// @Accept json
// @Produce json
// @Success 200 {object} dto.CarResponse
// @Router /car [post]
func PostCar(ctx *gin.Context) {
	carModel := models.Car{}

	if err := ctx.ShouldBindJSON(&carModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	carsCollection := database.CarsCollection

	insertResult, err := carsCollection.InsertOne(ctx, carModel)
	if err != nil {
		// panic(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	fmt.Println(insertResult.InsertedID)

	objectId := insertResult.InsertedID.(primitive.ObjectID)

	ctx.JSON(http.StatusOK, gin.H{
		"id": objectId,
	})
}

// @Summary List Car's resource
// @Description Desc
// @Tags car
// @Accept json
// @Produce json
// @Success 200 {object} []dto.CarResponse
// @Router /car [get]
func GetCar(ctx *gin.Context) {

	carsCollection := database.CarsCollection

	var ctxTodo = context.TODO()

	cursor, err := carsCollection.Find(ctxTodo, bson.D{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer cursor.Close(ctxTodo)

	var results []dto.CarResponse

	for cursor.Next(ctx) {
		var user dto.CarResponse
		if err := cursor.Decode(&user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding cars"})
			return
		}
		results = append(results, user)
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	fmt.Println(len(results))

	ctx.JSON(http.StatusOK, results)
}
