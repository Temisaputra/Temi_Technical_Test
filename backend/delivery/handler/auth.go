package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/pkg/helper"
)

type authUsecase interface {
	Login(ctx context.Context, req *presenter.LoginRequest) (res *presenter.LoginResponse, err error)
	Register(ctx context.Context, req *presenter.RegisterRequest) (err error)
}

type AuthHandler struct {
	authUsecase authUsecase
}

func NewAuthHandler(authUsecase authUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body presenter.RegisterRequest true "Register request"
// @Success 201 {object} helper.Response "success"
// @Failure 400 {object} helper.Response "bad request"
// @Router /register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request presenter.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	log.Printf("Registering user: %v", request)

	err := h.authUsecase.Register(r.Context(), &request)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusCreated
	response.Message = "success"
	response.Data = nil

	helper.WriteResponse(w, nil, &response)
}

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body presenter.LoginRequest true "Login request"
// @Success 200 {object} helper.Response{data=presenter.LoginResponse{}} "success"
// @Failure 400 {object} helper.Response "bad request"
// @Router /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request presenter.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	res, err := h.authUsecase.Login(r.Context(), &request)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"
	response.Data = res

	helper.WriteResponse(w, nil, &response)
}
