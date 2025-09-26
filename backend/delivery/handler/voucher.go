package handler

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/delivery/presenter/request"
	"github.com/Temisaputra/warOnk/delivery/presenter/response"
	"github.com/Temisaputra/warOnk/pkg/helper"
	"github.com/gorilla/mux"
)

type voucherUsecase interface {
	GetListVoucher(ctx context.Context, params presenter.VoucherParams, pagination *request.Pagination) (vouchers []presenter.VoucherResponse, meta response.Meta, err error)
	GetVoucherByID(ctx context.Context, id int) (voucher presenter.VoucherResponse, err error)
	CreateVoucher(ctx context.Context, params presenter.VoucherRequest) error
	UpdateVoucher(ctx context.Context, params presenter.VoucherRequest, no int) error
	DeleteVoucher(ctx context.Context, id int) error
	UploadVouchersFromCSV(ctx context.Context, vouchers []presenter.VoucherRequest) error
	ExportVouchersToCSV(ctx context.Context) (fileName string, err error)
}

type VoucherHandler struct {
	voucherUsecase voucherUsecase
}

func NewVoucherHandler(voucherUsecase voucherUsecase) *VoucherHandler {
	return &VoucherHandler{
		voucherUsecase: voucherUsecase,
	}
}

// GetListVoucher godoc
// @Tags Voucher
// @Summary Get List Voucher
// @Description Get List Voucher
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param keyword query string false "Keyword for search"
// @Param order_by query string false "Order by field"
// @Param order_type query string false "Order type (asc/desc)"
// @Success 200 {object} helper.Response{data=[]presenter.VoucherResponse,meta=response.Meta}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Security BearerAuth
// @Router /vouchers [get]
func (h *VoucherHandler) GetListVoucher(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	var params presenter.VoucherParams
	params.VoucherCode = r.URL.Query().Get("keyword")

	pagination := &request.Pagination{
		Keyword:   r.URL.Query().Get("keyword"),
		OrderBy:   r.URL.Query().Get("order_by"),
		OrderType: r.URL.Query().Get("order_type"),
		Page:      page,
		PageSize:  pageSize,
	}

	data, meta, err := h.voucherUsecase.GetListVoucher(r.Context(), params, pagination)
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

// GetProductByID godoc
// @Tags Product
// @Summary Get Product by ID
// @Description Get Voucher by ID
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Success 200 {object} helper.Response{data=presenter.VoucherResponse}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Security BearerAuth
// @Router /vouchers/{id} [get]
func (h *VoucherHandler) GetVoucherByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	if idInt == 0 {
		helper.WriteResponse(w, helper.NewErrBadRequest("id is required"), nil)
		return
	}
	data, err := h.voucherUsecase.GetVoucherByID(r.Context(), idInt)
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

// CreateProduct godoc
// @Tags Voucher
// @Summary Create a new voucher
// @Description Create a new voucher
// @Accept json
// @Produce json
// @Param request body presenter.VoucherRequest true "Voucher data"
// @Success 201 {object} helper.Response{data=presenter.VoucherResponse}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Security BearerAuth
// @Router /voucher-create [post]
func (h *VoucherHandler) CreateVoucher(w http.ResponseWriter, r *http.Request) {
	var params presenter.VoucherRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.voucherUsecase.CreateVoucher(r.Context(), params)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusCreated
	response.Message = "success"

	helper.WriteResponse(w, nil, &response)
}

// UpdateVoucher godoc
// @Tags Voucher
// @Summary Update a voucher
// @Description Update a voucher
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Param request body presenter.VoucherRequest true "Voucher data"
// @Success 200 {object} helper.Response{data=presenter.VoucherResponse}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Security BearerAuth
// @Router /voucher-update/{id} [put]
func (h *VoucherHandler) UpdateVoucher(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	log.Printf("idInt: %d", idInt)

	if idInt == 0 {
		helper.WriteResponse(w, helper.NewErrBadRequest("no is required"), nil)
		return
	}

	var params presenter.VoucherRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.voucherUsecase.UpdateVoucher(r.Context(), params, idInt)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"

	helper.WriteResponse(w, nil, &response)
}

// DeleteVoucher godoc
// @Tags Voucher
// @Summary Delete a voucher
// @Description Delete a voucher
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Success 200 {object} helper.Response{data=presenter.VoucherResponse}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Security BearerAuth
// @Router /voucher-delete/{id} [delete]
func (h *VoucherHandler) DeleteVoucher(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	if idInt == 0 {
		helper.WriteResponse(w, helper.NewErrBadRequest("id is required"), nil)
		return
	}

	err := h.voucherUsecase.DeleteVoucher(r.Context(), idInt)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"

	helper.WriteResponse(w, nil, &response)
}

// UploadVouchersFromCSV godoc
// @Tags Voucher
// @Summary Upload Vouchers from CSV
// @Description Upload Vouchers from CSV
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Security BearerAuth
// @Router /vouchers/upload-csv [post]
func (h *VoucherHandler) UploadVouchersFromCSV(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(100 << 20) // 100 MB

	// Dapatkan file dari form data
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Fail When Read CSV : ", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Buat pembaca CSV dari file
	reader := csv.NewReader(file)

	// Baca header CSV untuk melewatkan baris pertama
	_, err = reader.Read()
	if err != nil {
		http.Error(w, "Fail When Read CSV Header: ", http.StatusInternalServerError)
		return
	}

	// Baca semua data dari CSV
	records, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "Fail When Read CSV: ", http.StatusInternalServerError)
		return
	}

	dataVouchers := []presenter.VoucherRequest{}

	for _, record := range records {
		discount, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			http.Error(w, "Invalid discount value in CSV", http.StatusBadRequest)
			return
		}
		voucher := presenter.VoucherRequest{
			VoucherCode: record[0],
			Discount:    discount,
			ExpiredDate: record[2],
		}
		dataVouchers = append(dataVouchers, voucher)
	}

	err = h.voucherUsecase.UploadVouchersFromCSV(r.Context(), dataVouchers)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"

	helper.WriteResponse(w, nil, &response)
}

// ExportVouchersToCSV godoc
// @Tags Voucher
// @Summary Export Vouchers to CSV
// @Description Export Vouchers to CSV
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response{data=string}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Security BearerAuth
// @Router /vouchers/export-csv [get]
func (h *VoucherHandler) ExportVouchersToCSV(w http.ResponseWriter, r *http.Request) {
	fileName, err := h.voucherUsecase.ExportVouchersToCSV(r.Context())
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)

	file, err := os.Open(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}

	file.Close()
	os.Remove(fileName)
}
