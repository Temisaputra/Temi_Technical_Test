package usecase

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/delivery/presenter/request"
	"github.com/Temisaputra/warOnk/delivery/presenter/response"
	"github.com/Temisaputra/warOnk/delivery/repository"
	"github.com/Temisaputra/warOnk/internal/entity"
	"github.com/Temisaputra/warOnk/pkg/helper"
)

type VoucherUsecase struct {
	voucherRepo     repository.VoucherRepository
	transactionRepo repository.TransactionRepository
}

func NewVoucherUsecase(voucherRepository repository.VoucherRepository, transactionRepository repository.TransactionRepository) *VoucherUsecase {
	return &VoucherUsecase{
		voucherRepo:     voucherRepository,
		transactionRepo: transactionRepository,
	}
}

func (u *VoucherUsecase) GetListVoucher(ctx context.Context, params presenter.VoucherParams, pagination *request.Pagination) (vouchers []presenter.VoucherResponse, meta response.Meta, err error) {
	if pagination.Page == 0 {
		pagination.Page = 1
	}

	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}

	res, meta, err := u.voucherRepo.GetListVoucher(ctx, params, pagination)
	if err != nil {
		errMsg := fmt.Errorf("[VoucherUsecase-GetAllVoucher] Error when getting all voucher: %w", err)
		return nil, meta, errMsg
	}

	return res, meta, nil
}

func (u *VoucherUsecase) GetVoucherByID(ctx context.Context, no int) (res presenter.VoucherResponse, err error) {
	res, err = u.voucherRepo.GetVoucherByID(ctx, no)
	if err != nil {
		errMsg := fmt.Errorf("[VoucherUsecase-GetVoucherByID] Error when getting voucher by id: %w", err)
		return presenter.VoucherResponse{}, errMsg
	}

	return res, nil
}

func (u *VoucherUsecase) CreateVoucher(ctx context.Context, params presenter.VoucherRequest) (err error) {
	// Mulai transaksi opsional
	return u.transactionRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		storedData := entity.Voucher{
			VoucherCode: helper.GenerateVoucherCode(),
			Discount:    params.Discount,
			ExpiredDate: params.ExpiredDate,
		}
		if err := u.voucherRepo.CreateVoucher(txCtx, storedData); err != nil {
			return err // rollback otomatis
		}
		return nil // commit otomatis
	})
}

func (u *VoucherUsecase) UpdateVoucher(ctx context.Context, params presenter.VoucherRequest, no int) (err error) {
	return u.transactionRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		// Ambil voucher dulu, pakai txCtx
		voucher, err := u.voucherRepo.GetVoucherByID(txCtx, no)
		if err != nil {
			return fmt.Errorf("[VoucherUsecase-UpdateVoucher] Error when getting voucher by no: %w", err)
		}

		if voucher.No == 0 {
			return fmt.Errorf("[VoucherUsecase-UpdateVoucher] Voucher not found")
		}

		updatedData := entity.Voucher{
			No:          no,
			VoucherCode: voucher.VoucherCode,
			Discount:    params.Discount,
			ExpiredDate: params.ExpiredDate,
		}

		// Update voucher pakai txCtx
		if err := u.voucherRepo.UpdateVoucher(txCtx, updatedData); err != nil {
			return fmt.Errorf("[VoucherUsecase-UpdateVoucher] Error when updating voucher: %w", err)
		}

		return nil // commit otomatis kalau tidak ada error
	})
}

func (u *VoucherUsecase) DeleteVoucher(ctx context.Context, id int) (err error) {
	return u.transactionRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		// Ambil voucher dulu, pakai txCtx
		voucher, err := u.voucherRepo.GetVoucherByID(txCtx, id)
		if err != nil {
			return fmt.Errorf("[VoucherUsecase-DeleteVoucher] Error when getting voucher by id: %w", err)
		}

		if voucher.No == 0 {
			return fmt.Errorf("[ProductUsecase-DeleteProduct] Product not found")
		}

		if err := u.voucherRepo.DeleteVoucher(txCtx, id); err != nil {
			return fmt.Errorf("[VoucherUsecase-DeleteVoucher] Error when deleting voucher: %w", err)
		}

		return nil // commit otomatis kalau tidak ada error
	})
}

func (u *VoucherUsecase) UploadVouchersFromCSV(ctx context.Context, vouchers []presenter.VoucherRequest) (err error) {
	// Mulai transaksi opsional
	return u.transactionRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		for _, v := range vouchers {
			// Cek apakah voucher code sudah ada
			exists, err := u.voucherRepo.CheckVoucherCode(txCtx, v.VoucherCode)
			if err != nil {
				return fmt.Errorf("[VoucherUsecase-UploadVouchersFromCSV] Error when checking voucher code: %w", err)
			}
			if exists {
				continue // Lewati jika sudah ada
			}

			storedData := entity.Voucher{
				VoucherCode: v.VoucherCode,
				Discount:    v.Discount,
				ExpiredDate: v.ExpiredDate,
			}
			if err := u.voucherRepo.CreateVoucher(txCtx, storedData); err != nil {
				return err // rollback otomatis
			}
		}
		return nil // commit otomatis
	})
}

func (u *VoucherUsecase) ExportVouchersToCSV(ctx context.Context) (fileName string, err error) {
	// Implementasi fungsi ekspor voucher ke CSV
	exportData, err := u.voucherRepo.ExportVoucher(ctx)
	if err != nil {
		return "", fmt.Errorf("[VoucherUsecase-ExportVouchersToCSV] Error when exporting vouchers: %w", err)
	}

	if len(exportData) < 1 {
		err = errors.New("data not found")
		return "", err
	}

	// Membuat nama file dengan format yang sesuai
	currentDate := time.Now().Format("2006-01-02")
	fileName = fmt.Sprintf("Vouchers_%s.csv", currentDate)

	// Membuat file CSV
	file, err := os.Create(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer file.Close()

	// Inisialisasi writer CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Menulis header kolom ke file CSV
	headers := []string{"VOUCHER_CODE", "DISCOUNT", "EXPIRED_DATE"}
	if err := writer.Write(headers); err != nil {
		err = fmt.Errorf("failed to write headers to CSV: %v", err)
		return "", err
	}

	// Menulis data ke file CSV
	for _, voucher := range exportData {
		record := []string{
			voucher.VoucherCode,
			fmt.Sprintf("%.0f", voucher.Discount),
			voucher.ExpiredDate,
		}
		if err := writer.Write(record); err != nil {
			err = fmt.Errorf("failed to write record to CSV: %v", err)
			return "", err
		}
	}

	// Menyimpan file hasil generate CSV
	err = file.Sync()
	if err != nil {
		err = fmt.Errorf("failed to save CSV file: %v", err)
		return
	}

	return fileName, nil
}
