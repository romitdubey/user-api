package handler

import (
	"strconv"
	"time"
	"go.uber.org/zap"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/romitdubey1/user-api/internal/models"
	"github.com/romitdubey1/user-api/internal/repository"
	"github.com/romitdubey1/user-api/internal/service"
)

/*
	Request DTOs
	(API input should NEVER bind directly to DB/domain models)
*/

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob" validate:"required"`
}

type UserHandler struct {
	repo      *repository.UserRepository
	service   *service.UserService
	validator *validator.Validate
	logger     *zap.Logger
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob" validate:"required"`
}

func NewUserHandler(
	r *repository.UserRepository,
	s *service.UserService,
	l *zap.Logger,

) *UserHandler {
	return &UserHandler{
		repo:      r,
		service:   s,
		validator: validator.New(),
		logger:    l,
	}
}

// POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest

	// Parse JSON body
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	// Validate input
	if err := h.validator.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Parse DOB string into time.Time
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"dob must be in YYYY-MM-DD format",
		)
	}

	// Create user
	user, err := h.repo.Create(
		c.Context(),
		req.Name,
		dob,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"failed to create user",
		)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GET /users/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	// Parse ID
	id64, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}
	id := int32(id64)

	// Fetch user
	user, err := h.repo.GetByID(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	// Build response with calculated age
	resp := models.User{
		ID:   int64(user.ID),
		Name: user.Name,
		DOB:  user.Dob,
		Age:  h.service.CalculateAge(user.Dob),
	}

	return c.JSON(resp)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id64, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}
	id := int32(id64)

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"dob must be in YYYY-MM-DD format",
		)
	}

	user, err := h.repo.Update(c.Context(), id, req.Name, dob)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update user")
	}

	h.logger.Info("user updated", zap.Int32("id", id))

	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id64, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}
	id := int32(id64)

	if err := h.repo.Delete(c.Context(), id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to delete user")
	}

	h.logger.Info("user deleted", zap.Int32("id", id))

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.repo.List(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to fetch users")
	}

	var resp []models.User
	for _, u := range users {
		resp = append(resp, models.User{
			ID:   int64(u.ID),
			Name: u.Name,
			DOB:  u.Dob,
			Age:  h.service.CalculateAge(u.Dob),
		})
	}

	return c.JSON(resp)
}
