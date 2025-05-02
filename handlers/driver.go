package handlers

import (
	"encoding/json"
	"f1fastapi/model"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllDrivers(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("data/drivers.json")
	if err != nil {
		http.Error(w, "Не удалось открыть файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	//читаем в слайс структур Driver
	var drivers []model.Driver
	err = json.NewDecoder(file).Decode(&drivers)
	if err != nil {
		http.Error(w, "Ошибка чтения JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drivers)

}
func GetDriverByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Не удалось превратить строку в число", http.StatusBadRequest)
		return
	}

	file, err := os.Open("data/drivers.json")
	if err != nil {
		http.Error(w, "Не удалось открыть файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	//читаем в слайс структур Driver по ID
	var driversId []model.Driver
	err = json.NewDecoder(file).Decode(&driversId)
	if err != nil {
		http.Error(w, "Ошибка чтения JSON", http.StatusInternalServerError)
		return
	}
	for _, d := range driversId {
		if d.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(d)
			return
		}

	}
	//если не нашли ни одного гонщика с заданным ID
	http.Error(w, "Такой гонщик не найден", 404)

}

func DeleteDriver(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Не удалось превратить строку в число", http.StatusBadRequest)
		return
	}

	file, err := os.Open("data/drivers.json")
	if err != nil {
		http.Error(w, "Ошибка чтения JSON", 500)
		return
	}
	defer file.Close()

	var drivers []model.Driver

	err = json.NewDecoder(file).Decode(&drivers)
	if err != nil {
		http.Error(w, "Ошибка чтения JSON", 500)
		return
	}
	var newList []model.Driver
	found := false
	for _, d := range drivers {
		if d.ID != id {
			newList = append(newList, d)

		} else {
			found = true
		}
	}
	if !found {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	newFile, err := os.Create("data/drivers.json")
	if err != nil {
		http.Error(w, "Не удалось сохранить файл", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	err = json.NewEncoder(newFile).Encode(newList)
	if err != nil {
		http.Error(w, "Ошибка при записи JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func CreateDriver(w http.ResponseWriter, r *http.Request) {

	var newDriver model.Driver
	err := json.NewDecoder(r.Body).Decode(&newDriver)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	file, err := os.Open("data/drivers.json")
	if err != nil {
		http.Error(w, "Ошибка чтения JSON", 500)
		return
	}
	defer file.Close()

	var drivers []model.Driver
	err = json.NewDecoder(file).Decode(&drivers)
	if err != nil {
		http.Error(w, "Ошибка чтения JSON", http.StatusInternalServerError)
		return
	}

	maxID := 0
	for _, d := range drivers {
		if d.ID > maxID {
			maxID = d.ID

		}

	}
	newDriver.ID = maxID + 1
	drivers = append(drivers, newDriver)

	newFile, err := os.Create("data/drivers.json")
	if err != nil {
		http.Error(w, "Не удалось сохранить файл", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	err = json.NewEncoder(newFile).Encode(&drivers)
	if err != nil {
		http.Error(w, "Не удалось записать данные", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
}
