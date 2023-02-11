package user

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Register(router fiber.Router)
}

type handler struct {
	s Service
	v RequestValidator
}

func NewHandler(s Service) Handler {
	return &handler{s: s, v: NewValidator()}
}

func (h *handler) Register(router fiber.Router) {
	router.Post("/registration", h.registration)
}

func (h *handler) registration(ctx *fiber.Ctx) error {
	var req RequestRegistration
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := h.v.Validate(req)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	res, err := h.s.Register(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(res)
}

type RequestValidator interface {
	Validate(req interface{}) []ErrorResponse
}

type requestValidator struct {
	v *validator.Validate
}

func NewValidator() RequestValidator {
	v := validator.New()
	_ = v.RegisterValidation("notblank", func(fl validator.FieldLevel) bool {
		return strings.TrimSpace(fl.Field().String()) != ""
	})

	return &requestValidator{v: v}
}

func (v *requestValidator) Validate(req interface{}) []ErrorResponse {
	var errs []ErrorResponse
	err := v.v.Struct(req)
	var validationErrors validator.ValidationErrors
	if err != nil && errors.As(err, &validationErrors) {
		for _, err := range validationErrors {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errs = append(errs, element)
		}
	}

	return errs
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type RequestRegistration struct {
	FirstName string `json:"firstName" validate:"required,notblank"`
	LastName  string `json:"lastName"  validate:"required,notblank"`
	Email     string `json:"email"     validate:"required,notblank,email"`
	Password  string `json:"password"  validate:"required,notblank"`
}
