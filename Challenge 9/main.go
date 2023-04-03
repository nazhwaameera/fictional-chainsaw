package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 15)

	c := make(chan string)

	go func() {
		for {
			select {
			case <-c:
				return

			case tm := <-ticker.C:
				log.Println("Current time is: ", tm)
				realTimeChecking()
			}
		}
	}()

	time.Sleep(time.Minute * 1)

	ticker.Stop()

	c <- "true"

	// go doInterval()
	// time.Sleep(time.Minute * 1)
}

// func doInterval() {
// 	for range time.Tick(time.Second * 15) {
// 		realTimeChecking()
// 	}
// }

func generateNum() (int, int) {
	water, wind := rand.Intn(100), rand.Intn(100)

	return water, wind
}

func createData(water int, wind int) map[string]interface{} {
	data := map[string]interface{}{
		"water": &water,
		"wind":  &wind,
	}

	return data
}

func checkStatus(water int, wind int) (string, string) {
	var statusWater, statusWind string
	if water < 5 {
		statusWater = "aman"
	} else if water >= 6 && water <= 8 {
		statusWater = "siaga"
	} else if water > 8 {
		statusWater = "bahaya"
	}

	if wind < 6 {
		statusWind = "aman"
	} else if wind >= 7 && wind <= 15 {
		statusWind = "siaga"
	} else if wind > 15 {
		statusWind = "bahaya"
	}

	return statusWater, statusWind
}

func realTimeChecking() {
	// menyiapkan data yang akan dijadikan sebagai request body
	// data := map[string]interface{}{
	// 	"water": 13,
	// 	"wind":  8,
	// }

	water, wind := generateNum()
	data := createData(water, wind)
	statusWater, statusWind := checkStatus(water, wind)

	// mengubah data yang sudah disiapkan menjadi sebuah data JSON.
	requestJson, err := json.Marshal(data)
	client := &http.Client{}

	if err != nil {
		log.Fatalln(err)
	}

	// menyiapkan request dengan menggunakan function http.NewRequest
	// menggunakan function bytes.NewBuffer untuk mengubah data menjaid interface io.reader.
	// perlu diubah karena parameter ketiga dari fungsi NewRequest menerima nilai dengan tipe data interface io.Reader.
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	// eksekusi request menggunakan method Do
	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
	log.Println("status water :", statusWater)
	log.Println("status wind :", statusWind)
}
