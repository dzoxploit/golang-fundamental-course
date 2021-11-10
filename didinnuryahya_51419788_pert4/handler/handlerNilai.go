package handler

import (
	"didinnuryahya_51419788_pert4/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HandlerNilaiGet(w http.ResponseWriter, r *http.Request) {
	var data interface{}

	var err error

	npm := r.URL.Query()["npm"]

	if len(npm) != 0 {
		data, err = model.GetNilai(db, npm[0])
	} else {
		data, err = model.GetAllNilai(db)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonData, _ := json.Marshal(data)

	w.Write(jsonData)
}

func HandlerNilaiPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var data model.Nilai
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

func HandlerNilaiDelete(w http.ResponseWriter, r *http.Request) {
	npm := r.URL.Query()["npm"]

	if len(npm) != 0 {
		data := model.Nilai{NPM: npm[0]}
		if err := data.Delete(db); err != nil {
			http.Error(w, "ID Tidak Ditemukan", http.StatusBadRequest)
			return
		}
		w.Write([]byte("Data Telah dihapus"))
	} else {
		http.Error(w, "ID Tidak ditemukan", http.StatusInternalServerError)
	}
}

func HandlerNilaiPut(w http.ResponseWriter, r *http.Request) {
	npm := r.URL.Query()["npm"]

	if len(npm) == 0 {
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

	data := model.Nilai{NPM: npm[0]}

	err = data.Update(db, jsonmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := model.GetNilai(db, npm[0])

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
