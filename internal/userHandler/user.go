package handler

import (
	"net/http"
	internal "w2/d3/internal/userDto"
	config "w2/d3/config/database"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"context"
	"time"
)

func Register(c echo.Context) error {
	var req internal.RegisterUser
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":"Invalid Request"})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Internal Server Error",
		})
	}

	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	var userID int

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = config.Pool.QueryRow(ctx, query, req.Name, req.Email, string(hashPassword)).Scan(&userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Registration failed",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User registered successfully",
		"user_id": string(userID),
	})
}