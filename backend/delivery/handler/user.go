package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/delivery/presenter/request"
	"github.com/Temisaputra/warOnk/delivery/presenter/response"
	"github.com/Temisaputra/warOnk/pkg/helper"
)

type userUsecase interface {
	GetAllUser(ctx context.Context, pagination *request.Pagination) (users []*presenter.UserResponse, meta response.Meta, err error)
	GetUserByID(ctx context.Context, id int) (user *presenter.UserResponse, err error)
	CreateUser(ctx context.Context, params *presenter.UserRequest) error
	UpdateUser(ctx context.Context, params *presenter.UserRequest, id int) error
	DeleteUser(ctx context.Context, id int) error
}

type UserHandler struct {
	userUsecase userUsecase
}

func NewUserHandler(userUsecase userUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	params := &request.Pagination{
		Keyword:   r.URL.Query().Get("keyword"),
		OrderBy:   r.URL.Query().Get("order_by"),
		OrderType: r.URL.Query().Get("order_type"),
		Page:      page,
		PageSize:  pageSize,
	}

	data, meta, err := h.userUsecase.GetAllUser(r.Context(), params)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"
	response.Meta = &meta
	response.Data = data

	helper.WriteResponse(w, nil, &response)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	data, err := h.userUsecase.GetUserByID(r.Context(), id)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"
	response.Data = data

	helper.WriteResponse(w, nil, &response)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request presenter.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	if err := h.userUsecase.CreateUser(r.Context(), &request); err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusCreated
	response.Message = "success"
	response.Data = nil

	helper.WriteResponse(w, nil, &response)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var request presenter.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	if err := h.userUsecase.UpdateUser(r.Context(), &request, id); err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"
	response.Data = nil

	helper.WriteResponse(w, nil, &response)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	if err := h.userUsecase.DeleteUser(r.Context(), id); err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"
	response.Data = nil

	helper.WriteResponse(w, nil, &response)
}
