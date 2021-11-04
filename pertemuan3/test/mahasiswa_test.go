package test

import (
	"fmt"
	"pertemuan3/model"
	"testing"
)

func TestMahasiswa(t *testing.T) {
	var dataInsertMhs = []model.Mahasiswa{
		model.Mahasiswa{
			NPM:   "12345678",
			Nama:  "Budi Doremi",
			Kelas: "3KA20",
		},
		model.Mahasiswa{
			NPM:   "19283746",
			Nama:  "Doremi Budi",
			Kelas: "4KA20",
		},
		model.Mahasiswa{
			NPM:   "4444444",
			Nama:  "DoBud",
			Kelas: "4KA21",
		},
	}

	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("Testing insert mahasiswa", func(t *testing.T) {
		for _, dataInsert := range dataInsertMhs {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Testing update  mahasiswa", func(t *testing.T) {
		var updateData = map[string]interface{}{
			"nama": "Abdi Teh Ayuna",
		}

		data := dataInsertMhs[0]

		if err := data.Update(db, updateData); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Testing delete  mahasiswa", func(t *testing.T) {
		data := dataInsertMhs[0]

		if err := data.Delete(db); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Testing get  mahasiswa", func(t *testing.T) {
		value, err := model.GetMahasiswa(db, "4444444")

		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%v", value)
	})

	t.Run("Testing get all mahasiswa", func(t *testing.T) {
		value, err := model.GetAllMahasiswa(db)

		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%v", value)
	})
}
