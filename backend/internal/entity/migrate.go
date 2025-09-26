package entity

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Voucher{},
		&Users{},
	)
}

func Drop(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&Voucher{},
		&Users{},
	)
}
