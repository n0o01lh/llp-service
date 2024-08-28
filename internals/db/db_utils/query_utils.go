package dbutils

import (
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func IsRecordExists(database *gorm.DB, table string, query string, args ...any) (bool, error) {
	var exists bool
	err := database.Table(table).
		Select("count(*) > 0").
		Where(query, args...).
		Find(&exists).Error

	if err != nil {
		log.Error(err)
		return true, err
	}

	return exists, nil
}
