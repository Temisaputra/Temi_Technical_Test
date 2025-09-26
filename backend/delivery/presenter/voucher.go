package presenter

type VoucherRequest struct {
	VoucherCode string  `json:"voucher_code"`
	Discount    float64 `json:"discount" validate:"required"`
	ExpiredDate string  `json:"expired_date" validate:"required" `
}

type VoucherResponse struct {
	No          int     `json:"no"`
	VoucherCode string  `json:"voucher_code"`
	Discount    float64 `json:"discount"`
	ExpiredDate string  `json:"expired_date"`
}

type VoucherParams struct {
	VoucherCode string `json:"voucher_code"`
}
