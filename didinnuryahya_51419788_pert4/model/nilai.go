package model

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var TableNilai string = `
	CREATE TABLE nilai(
		id_nilai INT PRIMARY KEY AUTO_INCREMENT,
		kd_mk VARCHAR(10),
		npm VARCHAR(10),
		uts REAL NOT NULL,
		uas REAL NOT NULL, 
		total REAL NOT NULL, 
		grade VARCHAR(1) NOT NULL
	);
`

type Nilai struct {
	Id_nilai int     `json:"idnilai"`
	Kd_mk    string  `json:"kd_mk"`
	NPM      string  `json:"npm"`
	UAS      float64 `json:"uas"`
	UTS      float64 `json:"uts"`
	Total    float64 `json:"total"`
	Grade    string  `json:"grade"`
}

func (n *Nilai) Fields() ([]string, []interface{}) {
	fields := []string{"id_nilai", "kd_mk", "npm", "uas", "uts", "total", "grade"}
	temp := []interface{}{&n.Id_nilai, &n.Kd_mk, &n.NPM, &n.UAS, &n.UTS, &n.Total, &n.Grade}
	return fields, temp
}

func (n *Nilai) Structur() *Nilai {
	return &Nilai{}
}

func (n *Nilai) Insert(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO %v values (?, ?, ?, ?, ?, ?, ?)", "nilai")
	_, err := db.Exec(query, &n.Id_nilai, &n.Kd_mk, &n.NPM, &n.UAS, &n.UTS, &n.Total, &n.Grade)
	return err
}

func (n *Nilai) Update(db *sql.DB, data map[string]interface{}) error {
	var kolom = []string{}
	var args []interface{}

	i := 1

	for key, value := range data {
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
		i++
	}

	dataUpdate := strings.Join(kolom, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", "nilai", dataUpdate, "id_nilai")

	args = append(args, n.Id_nilai)

	_, err := db.Exec(query, args...)
	return err
}

func (n *Nilai) Delete(db *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", "nilai", "id_nilai")

	_, err := db.Exec(query, n.Id_nilai)
	return err
}

func GetNilai(db *sql.DB, id string) (*Nilai, error) {
	n := &Nilai{}
	each := n.Structur()
	_, dst := each.Fields()
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", "nilai", "id_nilai")

	err := db.QueryRow(query, id).Scan(dst...)

	if err != nil {
		return nil, err
	}
	return each, nil
}

func GetAllNilai(db *sql.DB) ([]*Nilai, error) {
	n := &Nilai{}
	query := fmt.Sprintf("SELECT * FROM %s", "nilai")
	data, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	var result []*Nilai

	for data.Next() {
		each := n.Structur()
		_, dst := each.Fields()

		err := data.Scan(dst...)

		if err != nil {
			return nil, err
		}
		fmt.Println(each)
		result = append(result, each)
	}
	return result, nil
}
