package db

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"errors"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/delivery/presenter/request"
	"github.com/Temisaputra/warOnk/delivery/presenter/response"
	irepository "github.com/Temisaputra/warOnk/delivery/repository"
	"github.com/Temisaputra/warOnk/internal/entity"
	"gorm.io/gorm"
)

type VoucherRepository struct {
	*TransactionRepository
}

func NewVoucherRepo(db *gorm.DB) irepository.VoucherRepository {
	return &VoucherRepository{
		TransactionRepository: NewTransactionRepo(db),
	}
}

func (r *VoucherRepository) GetListVoucher(ctx context.Context, params presenter.VoucherParams, pagination *request.Pagination) (vouchers []presenter.VoucherResponse, meta response.Meta, err error) {
	db := r.Conn(ctx).WithContext(ctx).Model(&entity.Voucher{}).Where("deleted_at IS NULL")

	if pagination.Keyword != "" {
		keywordStr := "%" + pagination.Keyword + "%"
		db = db.Where("voucher_code ILIKE ? ", keywordStr)
	}

	if pagination.OrderBy != "" || pagination.OrderType != "" {
		splitOrder := strings.Split(pagination.OrderBy, "|")
		if len(splitOrder) > 1 {
			db = db.Order(splitOrder[0] + " " + splitOrder[1] + " " + pagination.OrderType)
		} else {
			db = db.Order(pagination.OrderBy + " " + pagination.OrderType)
		}
	} else {
		db = db.Order("updated_at DESC NULLS LAST")
	}

	offset := pagination.GetOffset()
	limit := pagination.GetLimit()

	if err = db.Count(&meta.TotalData).Error; err != nil {
		return nil, meta, err
	}

	var result []entity.Voucher

	if err = db.Offset(offset).Limit(limit).Find(&result).Error; err != nil {
		return nil, meta, err
	}

	meta.Page = int64(pagination.Page)
	meta.PageSize = int64(pagination.PageSize)
	meta.TotalPage = int64(math.Ceil(float64(meta.TotalData) / float64(pagination.PageSize)))

	for _, item := range result {
		vouchers = append(vouchers, item.ToPresenter())
	}

	return
}

func (r *VoucherRepository) GetVoucherByID(ctx context.Context, no int) (voucher presenter.VoucherResponse, err error) {
	var result entity.Voucher
	db := r.Conn(ctx).WithContext(ctx).Model(&entity.Voucher{}).Where("no = ?", no)
	if err := db.Where("deleted_at IS NULL").First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return presenter.VoucherResponse{}, fmt.Errorf("[VoucherRepository-GetVoucherByID] Voucher not found: %w", err)
		}
		return presenter.VoucherResponse{}, fmt.Errorf("[VoucherRepository-GetVoucherByID] Error when getting voucher by id: %w", err)
	}

	return result.ToPresenter(), nil
}

func (r *VoucherRepository) CreateVoucher(ctx context.Context, params entity.Voucher) error {
	conn := r.Conn(ctx)
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	storedData := map[string]interface{}{
		"voucher_code": params.VoucherCode,
		"discount":     params.Discount,
		"expired_date": params.ExpiredDate,
		"created_at":   currentTime,
	}
	db := conn.WithContext(ctx).Model(&entity.Voucher{})
	if err := db.Create(&storedData).Error; err != nil {
		return err
	}

	return nil
}

func (r *VoucherRepository) UpdateVoucher(ctx context.Context, params entity.Voucher) error {
	conn := r.Conn(ctx)
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	updatedData := map[string]interface{}{
		"voucher_code": params.VoucherCode,
		"discount":     params.Discount,
		"expired_date": params.ExpiredDate,
		"updated_at":   currentTime,
	}

	db := conn.WithContext(ctx).Model(&entity.Voucher{})
	if err := db.Where("no = ?", params.No).Updates(&updatedData).Error; err != nil {
		return err
	}

	return nil
}

func (r *VoucherRepository) DeleteVoucher(ctx context.Context, id int) error {
	conn := r.Conn(ctx).WithContext(ctx)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	deletedData := map[string]interface{}{
		"deleted_at": currentTime,
	}

	db := conn.WithContext(ctx).Model(&entity.Voucher{})
	if err := db.Where("no = ?", id).Updates(&deletedData).Error; err != nil {
		return err
	}

	return nil
}

func (r *VoucherRepository) CheckVoucherCode(ctx context.Context, code string) (bool, error) {
	var count int64
	db := r.Conn(ctx).WithContext(ctx).Model(&entity.Voucher{}).Where("voucher_code = ? AND deleted_at IS NULL", code)
	if err := db.Count(&count).Error; err != nil {
		return false, nil
	}
	return count > 0, nil
}

func (r *VoucherRepository) ExportVoucher(ctx context.Context) (vouchers []presenter.VoucherResponse, err error) {
	db := r.Conn(ctx).WithContext(ctx).Model(&entity.Voucher{}).Where("deleted_at IS NULL")

	var result []entity.Voucher

	if err = db.Find(&result).Error; err != nil {
		return nil, err
	}

	for _, item := range result {
		vouchers = append(vouchers, item.ToPresenter())
	}

	return
}
