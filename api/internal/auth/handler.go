package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service,
	}
}

type SignInBody struct {
	Email    string
	Password string
}

func (h *Handler) SignIn(c *gin.Context) {
	body := SignInBody{}

	if err := c.Bind(body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}

type SignUpBody struct {
	Email           string
	FirstName       string
	LastName        string
	Password        string
	ConfirmPassword string
}

func (h *Handler) SignUp(c *gin.Context) {
	body := SignUpBody{}

	if err := c.Bind(body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	h.service.SignUp(SignUpParams{
		Email:           body.Email,
		FirstName:       body.FirstName,
		LastName:        body.LastName,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	})
}

func (h *Handler) SignOut() {}
