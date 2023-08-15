package seeder

import (
	"fmt"
	"github.com/oktopriima/dimy/app/seeder/seeders"
	"gorm.io/gorm"
)

type Seeder interface {
	ShouldRun() bool
	Prepare() error
	Run() error
	Commit() error
	Rollback() error
}

func Seed(db *gorm.DB) {
	runner([]Seeder{
		seeders.NewCustomers(db),
		seeders.NewCustomerAddress(db),
		seeders.NewProduct(db),
		seeders.NewPaymentMethod(db),
	})
}

func runner(list []Seeder) {
	for _, s := range list {
		if !s.ShouldRun() {
			// Skip the current seeder
			continue
		}
		fmt.Printf("seeding %T\n", s)
		if err := s.Prepare(); err != nil {
			panic(err)
		}
		if err := s.Run(); err != nil {
			if err := s.Rollback(); err != nil {
				fmt.Println("Rollback failed", err)
			}
			panic(err)
		}
		if err := s.Commit(); err != nil {
			fmt.Println("Commit failed", err)
		}
	}
}
