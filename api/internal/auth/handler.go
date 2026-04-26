package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/noel-vega/finances/api/internal/apierr"
	"github.com/noel-vega/finances/api/internal/user"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service,
	}
}

type SignInBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(ctx *gin.Context) {
	body := SignInBody{}

	if err := ctx.Bind(body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
}

type SignUpBody struct {
	Email           string `json:"email" binding:"required,email,max=254"`
	FirstName       string `json:"firstName" binding:"required,max=100"`
	LastName        string `json:"lastName" binding:"required,max=100"`
	Password        string `json:"password" binding:"required,min=12,max=72"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
}

func (h *Handler) SignUp(ctx *gin.Context) {
	body := SignUpBody{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(err)
		var ve validator.ValidationErrors

		problems := []apierr.ProblemDetailError{}
		if errors.As(err, &ve) {
			for _, fe := range ve {
				problems = append(problems, apierr.FromFieldError(fe))
			}
			apierr.NewBadRequest(ctx, "One or more fields failed validation", problems)
			return
		}

		apierr.NewBadRequest(ctx, "Malformed request body.", nil)
		return
	}

	u, err := h.service.SignUp(ctx, SignUpParams{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  body.Password,
	})
	if err != nil {
		ctx.Error(err)
		if errors.Is(err, user.ErrEmailExists) {
			apierr.NewConflict(ctx, "User with email exists", nil)
			return
		}
		apierr.NewInternalServerError(ctx)
		return
	}

	ctx.JSON(http.StatusCreated, u)
}

func (h *Handler) SignOut() {}
