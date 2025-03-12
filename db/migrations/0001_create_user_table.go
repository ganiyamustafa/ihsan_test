package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	type users struct {
		Id            uuid.UUID `gorm:"type:uuid;primary_key"`
		Name          string    `gorm:"type:varchar(100);not null"`
		NIK           string    `gorm:"type:varchar(16);unique;index"`
		Phone         string    `gorm:"type:varchar(16);unique:index"`
		AccountNumber string    `gorm:"type:varchar(14);unique:index"`
		CreatedAt     time.Time `gorm:"index"`
		UpdatedAt     time.Time `gorm:"index"`
	}

	newMigration := gormigrate.Migration{
		ID: "0001",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.AutoMigrate(users{}); err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(users{})
		},
	}
	migrations = append(migrations, &newMigration)
}
