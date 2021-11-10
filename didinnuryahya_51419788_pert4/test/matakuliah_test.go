package test

import (
	"didinnuryahya_51419788_pert4/model"
	"fmt"
	"testing"
)

func TestMatakuliah(t *testing.T) {
	var dataInsertMatkul = []model.Matkul{
		model.Matkul{
			Kd_mk:      "KA01",
			Matakuliah: "MTK Dasar",
		},
		model.Matkul{
			Kd_mk:      "KA02",
			Matakuliah: "B.Indo Dasar",
		},
		model.Matkul{
			Kd_mk:      "KA03",
			Matakuliah: "IPA Dasar",
		},
	}

	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("Testing insert matkul", func(t *testing.T) {
		for _, dataInsert := range dataInsertMatkul {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Testing update matkul", func(t *testing.T) {
		var updateData = map[string]interface{}{
			"matakuliah": "Matematika Informatika",
		}

		data := dataInsertMatkul[0]

		if err := data.Update(db, updateData); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Testing delete  matkul", func(t *testing.T) {
		data := dataInsertMatkul[0]

		if err := data.Delete(db); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Testing get matkul", func(t *testing.T) {
		value, err := model.GetMatkul(db, "KA02")

		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%v", value)
	})

	t.Run("Testing get all matkul", func(t *testing.T) {
		value, err := model.GetAllMatkul(db)

		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%v", value)
	})
}
