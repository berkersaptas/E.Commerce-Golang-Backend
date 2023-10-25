package controllers

import (
	"FluxStore/configs"
	"FluxStore/models"
	"FluxStore/responses"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")

var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.UserCreateModel
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{StatusCode: http.StatusBadRequest, StatusType: "error", Message: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{StatusCode: http.StatusBadRequest, StatusType: "error", Message: validationErr.Error()})
	}

	newUser := models.UserCreateModel{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	userExist := userCollection.FindOne(ctx, bson.M{"email": newUser.Email}).Decode(&user)
	if userExist == nil {
		return c.Status(http.StatusConflict).JSON(responses.ErrorResponse{StatusCode: http.StatusConflict, StatusType: "error", Message: "E-mail address is on record"})
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{StatusCode: http.StatusInternalServerError, StatusType: "error", Message: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func LoginUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.UserLoginModel
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{StatusCode: http.StatusBadRequest, StatusType: "error", Message: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{StatusCode: http.StatusBadRequest, StatusType: "error", Message: validationErr.Error()})
	}

	loginUser := models.UserLoginModel{
		Email:    user.Email,
		Password: user.Password,
	}

	err := userCollection.FindOne(ctx, bson.M{"email": loginUser.Email, "password": loginUser.Password}).Decode(&loginUser)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responses.ErrorResponse{StatusCode: http.StatusUnauthorized, StatusType: "error", Message: "User Unauthorized"})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User Login Success"}})
}

func ForgetPassword(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.UserForgetPasswordModel
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{StatusCode: http.StatusBadRequest, StatusType: "error", Message: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{StatusCode: http.StatusBadRequest, StatusType: "error", Message: validationErr.Error()})
	}

	newUser := models.UserForgetPasswordModel{
		Email: user.Email,
	}

	userExist := userCollection.FindOne(ctx, bson.M{"email": newUser.Email}).Decode(&user)
	if userExist != nil {
		return c.Status(http.StatusUnauthorized).JSON(responses.ErrorResponse{StatusCode: http.StatusUnauthorized, StatusType: "error", Message: "E-mail address is not exist"})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

func ResetPassword(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.UserResetPasswordModel
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{StatusCode: http.StatusBadRequest, StatusType: "error", Message: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{StatusCode: http.StatusBadRequest, StatusType: "error", Message: validationErr.Error()})
	}

	newUser := models.UserResetPasswordModel{
		Email:    user.Email,
		Password: user.Password,
	}

	filter := bson.M{"email": newUser.Email}
	update := bson.M{"$set": bson.M{"password": newUser.Password}}

	updateResult, err := userCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{StatusCode: http.StatusInternalServerError, StatusType: "error", Message: err.Error()})
	}

	if updateResult.MatchedCount == 0 {
		return c.Status(http.StatusUnauthorized).JSON(responses.ErrorResponse{StatusCode: http.StatusUnauthorized, StatusType: "error", Message: "Not matched user"})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Password changed correct"}})
}
