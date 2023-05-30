//go:generate mockery --all --inpackage --case snake

package health

import (
	"fmt"

	"gorm.io/gorm"
)

type versionGetterRepository struct {
	db *gorm.DB
}

type VersionGetterRepository interface {
	GetVersion() (string, error)
}

func (r versionGetterRepository) GetVersion() (string, error) {
	var version string
	err := r.db.Raw("SELECT VERSION()").Scan(&version).Error
	if err != nil {
		err = fmt.Errorf("db.Scan: %w", err)
	}
	return version, err
}
