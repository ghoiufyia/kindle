package dogo

import (
	"github.com/jinzhu/gorm"
	"fmt"
	// Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDatabase(cfg *Config) (*gorm.DB,error) {
	if cfg.Database.Type == "mysql" {
		args := fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.DatabaseName,
		)
		db,err := gorm.Open(cfg.Database.Type,args)
		if err != nil {
			return db,err
		}
		// Max idle connections
		db.DB().SetMaxIdleConns(cfg.Database.MaxIdleConns)

		// Max open connections
		db.DB().SetMaxOpenConns(cfg.Database.MaxOpenConns)

		// Database logging
		db.LogMode(cfg.IsDevelopment)

		return db, nil
	}
	return nil, fmt.Errorf("Database type %s not supported", cfg.Database.Type)
}

