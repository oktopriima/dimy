package config

import (
	"github.com/oktopriima/dimy/app/seeder"
	"gorm.io/gorm"
)

type Seeders struct {
	db *gorm.DB
}

func NewSeedersInstance(db *gorm.DB) *Seeders {
	return &Seeders{db: db}
}

func (s *Seeders) RunSeeders() {
	seeder.Seed(s.db)
}
