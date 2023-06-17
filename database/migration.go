package database

import (
	"fmt"
	"landtick/models"
	"landtick/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Transaksi{},
		&models.Train{},
		&models.Tiket{},
		&models.Station{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed!")
	}

	fmt.Println("Migration Success !")

}
