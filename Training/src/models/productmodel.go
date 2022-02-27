package models

import (
	"fmt"
	"training/src/config"
	"training/src/entitas"
)

type ProductModel struct {
}

func (*ProductModel) FindAll() ([]entitas.Product, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from product")
		if err2 != nil {
			return nil, err2
		} else {
			var products []entitas.Product
			for rows.Next() {
				var product entitas.Product
				rows.Scan(&product.Id, &product.Nama, &product.Harga, &product.Jumlah, &product.Deskripsi)
				products = append(products, product)
			}
			return products, nil
		}
	}
}

func (*ProductModel) Find(id int64) (entitas.Product, error) {
	db, err := config.GetDB()

	if err != nil {
		return entitas.Product{}, err
	} else {
		rows, err2 := db.Query("select * from product where id = ?", id)
		if err2 != nil {
			return entitas.Product{}, err2
		} else {
			var product entitas.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Nama, &product.Harga, &product.Jumlah, &product.Deskripsi)

			}
			fmt.Println(product)
			return product, nil
		}
	}
}

func (*ProductModel) Create(product *entitas.Product) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("insert into product(nama, harga, jumlah, deskripsi) values(?, ?, ?, ?)",
			product.Nama, product.Harga, product.Jumlah, product.Deskripsi)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*ProductModel) Update(product entitas.Product) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("update product set nama = ?, harga = ?,jumlah = ?, deskripsi= ? where id = ?", product.Nama, product.Harga, product.Jumlah, product.Deskripsi, product.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*ProductModel) Delete(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("delete from product where id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}
