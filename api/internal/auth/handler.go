package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/noel-vega/finances/api/internal/apierrors"
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

func (h *Handler) SignIn(c *gin.Context) {
	body := SignInBody{}

	if err := c.Bind(body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}

type SignUpBody struct {
	Email           string `json:"email" binding:"required,email"`
	FirstName       string `json:"firstName" binding:"required"`
	LastName        string `json:"lastName" binding:"required"`
	Password        string `json:"password" binding:"required,min=12"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
}

func (h *Handler) SignUp(ctx *gin.Context) {
	body := SignUpBody{}

	if err := ctx.ShouldBind(&body); err != nil {
		var ve validator.ValidationErrors

		problems := []apierrors.ProblemDetailError{}
		if errors.As(err, &ve) {
			for _, fe := range ve {
				problems = append(problems, apierrors.FromFieldError(fe))
			}
			apierrors.NewBadRequest(ctx, "One or more fields failed validation", problems)
			return
		}

		apierrors.NewBadRequest(ctx, "Malformed request body.", nil)
		return
	}

	u, err := h.service.SignUp(ctx, SignUpParams{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  body.Password,
	})
	if err != nil {
		if errors.Is(err, user.ErrEmailExists) {
			apierrors.NewConflict(ctx, "User with email exists", nil)
			return
		}
		apierrors.NewInternalServerError(ctx)
		return
	}

	ctx.JSON(http.StatusCreated, u)
}

func (h *Handler) SignOut() {}
