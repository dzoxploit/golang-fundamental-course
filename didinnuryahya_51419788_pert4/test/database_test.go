package test

import (
	"database/sql"
	"didinnuryahya_51419788_pert4/model"
	"fmt"
	"testing"
)

var username, password, host, nameDB, defaultDB string

func init() {
	username = "root"
	password = ""
	host = "localhost"
	nameDB = "gunadarma788"
	defaultDB = "mysql"

}

func TestDatabase(t *testing.T) {
	t.Run("Testing untuk membuat database", func(t *testing.T) {
		db, err := model.Connect(username, password, host, defaultDB)

		defer db.Close()

		if err != nil {
			t.Fatal(err)
		}

		err = model.CreateDB(db, nameDB)

		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing untuk memerika koneksi database"+nameDB, func(t *testing.T) {
		db, err := model.Connect(username, password, host, nameDB)

		defer db.Close()

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Testing untuk menghapus database", func(t *testing.T) {
		db, err := model.Connect(username, password, host, defaultDB)

		defer db.Close()

		if err != nil {
			t.Fatal(err)
		}

		err = model.DropDB(db, nameDB)

		if err != nil {
			t.Fatal(err)
		}

	})

}

func initDatabase() (*sql.DB, error) {
	dbInit, err := model.Connect(username, password, host, defaultDB)

	if err != nil {
		fmt.Println("Gagal Melakukan Koneksi")
	}

	if err = model.DropDB(dbInit, nameDB); err != nil {
		fmt.Println("Gagal Menghapus Database")
		return nil, err
	}

	if err = model.CreateDB(dbInit, nameDB); err != nil {
		fmt.Println("Gagal Memuat Database")
		return nil, err
	}

	dbInit.Close()
	db, err := model.Connect(username, password, password, nameDB)

	if err != nil {
		fmt.Println("Gagal Melakukan Koneksi")
		return nil, err
	}

	if err = model.CreateTable(db, model.TableMahasiswa); err != nil {
		fmt.Println("Gagal Membuat Table Mahasiswa")
		return nil, err
	}

	if err = model.CreateTable(db, model.TableMatkul); err != nil {
		fmt.Println("Gagal Membuat Table Matkul")
		return nil, err
	}

	if err = model.CreateTable(db, model.TableNilai); err != nil {
		fmt.Println("Gagal Membuat Table Nilai")
		return nil, err
	}

	return db, nil
}
