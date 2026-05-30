package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/TwiLightDM/diploma-gateway/internal/dto"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc/user-service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userClient *user_service.UserClient
}

func NewUserHandler(userClient *user_service.UserClient) *UserHandler {
	return &UserHandler{userClient: userClient}
}

// Login
// @Summary Авторизация пользователя
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "Email и пароль"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /auth/login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var request dto.UserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.userClient.Login(ctx, request.Email, request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.LoginResponse{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	})
}

// SignUp
// @Summary Регистрация пользователя
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "Данные пользователя"
// @Success 200 {object} dto.SignUpResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /auth/signup [post]
func (h *UserHandler) SignUp(c echo.Context) error {
	var request dto.UserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.userClient.SignUp(ctx, request.FullName, "student", request.Email, request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SignUpResponse{
		User: dto.UserResponse{
			Id:       response.User.Id,
			FullName: response.User.FullName,
			Role:     response.User.Role,
			Email:    response.User.Email,
		},
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	})
}

// CreateTeacher
// @Summary Создать преподавателя
// @Description Создание нового пользователя с ролью teacher
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "Данные преподавателя"
// @Success 200 {object} dto.SignUpResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /users/teachers [post]
func (h *UserHandler) CreateTeacher(c echo.Context) error {
	var request dto.UserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.userClient.SignUp(ctx, request.FullName, "teacher", request.Email, request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SignUpResponse{
		User: dto.UserResponse{
			Id:       response.User.Id,
			FullName: response.User.FullName,
			Role:     response.User.Role,
			Email:    response.User.Email,
		},
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	})
}

// Refresh
// @Summary Обновить токены
// @Tags Auth
// @Produce json
// @Success 200 {object} dto.LoginResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /auth/refresh [post]
func (h *UserHandler) Refresh(c echo.Context) error {
	id := c.Get("user_id").(string)
	role := c.Get("role").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.userClient.Refresh(ctx, id, role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.LoginResponse{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	})
}

// ReadUser
// @Summary Получить пользователя по ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /users/{id} [get]
func (h *UserHandler) ReadUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.userClient.ReadUser(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.UserResponse{
		Id:       response.User.Id,
		FullName: response.User.FullName,
		Role:     response.User.Role,
		Email:    response.User.Email,
	})
}

// ReadAllUser
// @Summary Получить список пользователей
// @Description Доступно только администратору
// @Tags Users
// @Produce json
// @Success 200 {object} dto.UserListResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /users [get]
func (h *UserHandler) ReadAllUser(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Error: "you need to be an admin"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.userClient.ReadUsers(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	users := make([]dto.UserResponse, 0, len(response.Users))
	for _, user := range response.Users {
		users = append(users, dto.UserResponse{
			Id:       user.Id,
			FullName: user.FullName,
			Role:     user.Role,
			Email:    user.Email,
		})
	}

	return c.JSON(http.StatusOK, dto.UserListResponse{
		Users: users,
	})
}

// ReadSelf
// @Summary Получить информацию о текущем пользователе
// @Tags Users
// @Produce json
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /users/me [get]
func (h *UserHandler) ReadSelf(c echo.Context) error {
	id := c.Get("user_id").(string)
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.userClient.ReadUser(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.UserResponse{
		Id:       response.User.Id,
		FullName: response.User.FullName,
		Role:     response.User.Role,
		Email:    response.User.Email,
	})
}

// UpdateUser
// @Summary Обновить профиль пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "Новые данные пользователя"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /users/me [put]
func (h *UserHandler) UpdateUser(c echo.Context) error {
	var request dto.UserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	id := c.Get("user_id").(string)
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.userClient.UpdateUser(ctx, id, request.FullName, request.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.UserResponse{
		Id:       response.User.Id,
		FullName: response.User.FullName,
		Role:     response.User.Role,
		Email:    response.User.Email,
	})
}

// ChangePassword
// @Summary Изменить пароль
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "Новый пароль"
// @Success 200
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /users/me/password [put]
func (h *UserHandler) ChangePassword(c echo.Context) error {
	var request dto.UserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	id := c.Get("user_id").(string)
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := h.userClient.ChangePassword(ctx, id, request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

// UpdateUserRole
// @Summary Изменить роль пользователя
// @Description Доступно только администратору
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "ID пользователя и новая роль"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /users/role [put]
func (h *UserHandler) UpdateUserRole(c echo.Context) error {
	var request dto.UserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	role := c.Get("role").(string)
	if role != "admin" {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Error: "you need to be an admin"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.userClient.UpdateUserRole(ctx, request.UserId, request.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.UserResponse{
		Id:       response.User.Id,
		FullName: response.User.FullName,
		Role:     response.User.Role,
		Email:    response.User.Email,
	})
}
