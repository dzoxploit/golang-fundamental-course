package test

import (
	"fmt"
	"pertemuan3/model"
	"testing"
)

func TestNilai(t *testing.T) {
	var dataInsertNilai = []model.Nilai{
		model.Nilai{
			Id_nilai: 1,
			Kd_mk:    "KA02",
			NPM:      "19283746",
			UAS:      70,
			UTS:      70,
			Total:    70,
			Grade:    "B",
		},
		model.Nilai{
			Id_nilai: 2,
			Kd_mk:    "KA03",
			NPM:      "4444444",
			UAS:      50,
			UTS:      50,
			Total:    50,
			Grade:    "D",
		},
	}

	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("Testing insert nilai", func(t *testing.T) {
		for _, dataInsert := range dataInsertNilai {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Testing update nilai", func(t *testing.T) {
		var updateData = map[string]interface{}{
			"uas":   40,
			"uts":   40,
			"total": 40,
			"grade": "E",
		}

		data := dataInsertNilai[0]

		if err := data.Update(db, updateData); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Testing delete nilai", func(t *testing.T) {
		data := dataInsertNilai[0]

		if err := data.Delete(db); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Testing get nilai", func(t *testing.T) {
		value, err := model.GetNilai(db, "2")

		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%v", value)
	})

	t.Run("Testing get all nilai", func(t *testing.T) {
		value, err := model.GetAllNilai(db)

		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("%v", value)
	})
}
