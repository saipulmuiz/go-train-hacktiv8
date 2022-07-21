package main

import (
	"assignment3/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", getHtml)
	http.HandleFunc("/status", getData)

	port := ":4444"

	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}

func getHtml(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Println("[INFO]", r.Method, r.URL.Path)
		rw.Header().Set("Content-Type", "text/html")
		tpl, err := template.ParseFiles("index.html")
		if err != nil {
			log.Println("[ERROR]", r.Method, r.URL.Path, "error :", err.Error())
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(rw, nil)
		return
	}

	errMsg := fmt.Sprintf("method %s not allowed !", r.Method)
	log.Println("[ERROR]", r.Method, r.URL.Path, "error :", errMsg)
	rw.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(rw).Encode(map[string]string{
		"error": "Selain method GET tidak bisa!",
	})
}

func getData(rw http.ResponseWriter, r *http.Request) {
	var newData *models.Data = readData()
	var statusWater, statusWind string
	water := rand.Intn(100)
	wind := rand.Intn(100)
	newData.Status.Water = water
	newData.Status.Wind = wind
	result, err := json.Marshal(newData)
	statusWater = checkWater(water)
	statusWind = checkWind(wind)

	if err != nil {
		json.NewEncoder(rw).Encode(map[string]string{
			"status": "false",
			"error":  err.Error(),
		})
	}
	err = writeData(result)
	if err != nil {
		json.NewEncoder(rw).Encode(map[string]string{
			"status": "false",
			"error":  err.Error(),
		})
	}
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"status":      true,
		"payload":     newData,
		"statusWater": statusWater,
		"statusWind":  statusWind,
	})
}

func readData() *models.Data {
	file, err := ioutil.ReadFile("file.json")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	var data models.Data
	json.Unmarshal(file, &data)
	return &data
}

func writeData(res []byte) error {
	err := ioutil.WriteFile("file.json", res, 0644)
	if err != nil {
		return err
	}
	return nil
}

func checkWater(data int) string {
	var statusWater string
	if data < 5 {
		statusWater = "Aman"
	} else if data >= 6 && data <= 8 {
		statusWater = "Siaga"
	} else if data > 8 {
		statusWater = "Bahaya"
	}
	return statusWater
}

func checkWind(data int) string {
	var statusWind string
	if data < 6 {
		statusWind = "Aman"
	} else if data >= 7 && data <= 15 {
		statusWind = "Siaga"
	} else if data > 15 {
		statusWind = "Bahaya"
	}
	return statusWind
}
