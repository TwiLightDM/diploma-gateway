package handlers

import (
	"context"
	"github.com/TwiLightDM/diploma-gateway/internal/dto"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc/user-service"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type UserHandler struct {
	userClient *user_service.UserClient
}

func NewUserHandler(userClient *user_service.UserClient) *UserHandler {
	return &UserHandler{userClient: userClient}
}

func (h *UserHandler) Login(c echo.Context) error {
	var request dto.UserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.LoginResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.userClient.Login(ctx, request.Email, request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.LoginResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.LoginResponse{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	})
}

func (h *UserHandler) SignUp(c echo.Context) error {
	var request dto.UserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.SignUpResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := h.userClient.SignUp(ctx, request.FullName, request.Role, request.Email, request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.SignUpResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SignUpResponse{})
}

func (h *UserHandler) ReadUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.UserResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.userClient.ReadUser(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.UserResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.UserResponse{
		Email:    response.Email,
		FullName: response.FullName,
	})
}

func (h *UserHandler) ReadSelf(c echo.Context) error {
	id := c.Get("user_id").(string)
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.UserResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.userClient.ReadUser(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.UserResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.UserResponse{
		Email:    response.Email,
		FullName: response.FullName,
	})
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var request dto.UserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.UserResponse{Error: "invalid request"})
	}

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.UserResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.userClient.UpdateUser(ctx, id, request.FullName, request.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.UserResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.UserResponse{
		Email:    response.Email,
		FullName: response.FullName,
	})
}

func (h *UserHandler) ChangePassword(c echo.Context) error {
	var request dto.UserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.UserResponse{Error: "invalid request"})
	}

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.UserResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := h.userClient.ChangePassword(ctx, id, request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.UserResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.UserResponse{})
}
