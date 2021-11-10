package handler

import (
	"didinnuryahya_51419788_pert4/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HandlerMatkulGet(w http.ResponseWriter, r *http.Request) {
	var data interface{}

	var err error

	kdmk := r.URL.Query()["kd_mk"]

	if len(kdmk) != 0 {
		data, err = model.GetMatkul(db, kdmk[0])
	} else {
		data, err = model.GetAllMatkul(db)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonData, _ := json.Marshal(data)

	w.Write(jsonData)
}

func HandlerMatkulPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var data model.Matkul
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err = data.Insert(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

func HandlerMatkulDelete(w http.ResponseWriter, r *http.Request) {
	kdmk := r.URL.Query()["kd_mk"]

	if len(kdmk) != 0 {
		data := model.Matkul{Kd_mk: kdmk[0]}
		if err := data.Delete(db); err != nil {
			http.Error(w, "ID Tidak Ditemukan", http.StatusBadRequest)
			return
		}
		w.Write([]byte("Data Telah dihapus"))
	} else {
		http.Error(w, "ID Tidak ditemukan", http.StatusInternalServerError)
	}
}

func HandlerMatkulPut(w http.ResponseWriter, r *http.Request) {
	kdmk := r.URL.Query()["kd_mk"]

	if len(kdmk) == 0 {
		http.Error(w, "ID Tidak Ditemukan", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonmap := make(map[string]interface{})

	err = json.Unmarshal(body, &jsonmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := model.Matkul{Kd_mk: kdmk[0]}

	err = data.Update(db, jsonmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := model.GetMatkul(db, kdmk[0])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

}
