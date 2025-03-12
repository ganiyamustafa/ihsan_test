package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	type user struct {
		ID uuid.UUID `gorm:"type:uuid;primary_key"`
	}

	type wallets struct {
		ID        uuid.UUID `gorm:"type:uuid;primary_key"`
		UserID    uuid.UUID `gorm:"type:uuid"`
		Balance   float64
		Currency  string `gorm:"type:varchar(5);default:'IDR'"`
		User      user   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
		CreatedAt time.Time
		UpdatedAt time.Time `gorm:"index"`
	}

	newMigration := gormigrate.Migration{
		ID: "0002",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.AutoMigrate(wallets{}); err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(wallets{})
		},
	}
	migrations = append(migrations, &newMigration)
}
