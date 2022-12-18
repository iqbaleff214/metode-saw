package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Dataset struct {
	Criteria     []map[string]any
	Alternatives []map[string]any
}

func main() {
	// deklarasi variable dataset untuk menampung data dari dataset.json
	var dataset Dataset

	// membaca isi file dataset.json dan menampungnya ke dalam variable file dalam bentuk byte
	file, err := os.ReadFile("dataset.json")
	if err != nil {
		panic(err)
	}

	// mengubah bentuk byte menjadi tipe data konkrit dan memasukkannya ke variabel dataset
	if err := json.Unmarshal(file, &dataset); err != nil {
		panic(err)
	}

	// normalisasi nilai bobot
	dataset.weightNormalization()

	// normalisasi nilai matriks
	dataset.matrixNormalization()

	// menghitung pembobotan matriks yang telah dinormalisasi
	dataset.resultCalculation()

	// mengurutkan hasil berdasarkan pembobotan matriks tertinggi
	sort.Slice(dataset.Alternatives, func(i, j int) bool {
		return dataset.Alternatives[i]["result"].(float64) > dataset.Alternatives[j]["result"].(float64)
	})

	// menampilkan hasil menggunakan std out CLI
	fmt.Println("================================")
	fmt.Println("No.\tNama \t\tHasil")
	fmt.Println("================================")
	for i, a := range dataset.Alternatives {
		fmt.Printf("%2d\t%-10s \t %.2f\n", i+1, a["name"], a["result"])
	}
}

// method untuk menormalisasi nilai bobot
func (dataset *Dataset) weightNormalization() {
	// Menghitung total bobot
	var total float64
	for _, c := range dataset.Criteria {
		total += c["weight"].(float64)
	}

	// Menghitung hasil normalisasi masing-masing bobot
	for _, c := range dataset.Criteria {
		c["normalizedWeight"] = c["weight"].(float64) / total
	}
}

// method untuk menormalisasi nilai matriks
func (dataset *Dataset) matrixNormalization() {
	// menentukan nilai pembagi untuk masing-masing kriteria berdasarkan nilai-nilai alternatif
	for _, a := range dataset.Alternatives {
		for _, c := range dataset.Criteria {
			if c["divisorValue"] == nil {
				c["divisorValue"] = 0.0
			}
			c["divisorValue"] = divisorValue(a[c["code"].(string)].(float64), c["divisorValue"].(float64), c["type"].(string))
		}
	}

	// normalisasi matriks berdasarkan nilai pembaginya
	for _, a := range dataset.Alternatives {
		for _, c := range dataset.Criteria {
			a["normalized"+c["code"].(string)] = normalize(a[c["code"].(string)].(float64), c["divisorValue"].(float64), c["type"].(string))
		}
	}
}

// method untuk menghitung pembobotan matriks yang telah dinormalisasi
func (dataset *Dataset) resultCalculation() {
	for _, a := range dataset.Alternatives {
		var result float64
		for _, c := range dataset.Criteria {
			result += (a["normalized"+c["code"].(string)].(float64) * c["normalizedWeight"].(float64))
		}
		a["result"] = result
	}
}

// function untuk menghitung nilai normalisasi matriks
func normalize(matrix, divisor float64, criteriaType string) float64 {
	if criteriaType == "cost" {
		return divisor / matrix
	}

	return matrix / divisor
}

// function untuk mendapatkan pembagi yang akan digunakan untuk normalisasi matriks
func divisorValue(value, initial float64, criteriaType string) float64 {
	if criteriaType == "cost" && (value < initial || initial == 0) {
		return value
	}

	if criteriaType == "benefit" && value > initial {
		return value
	}

	return initial
}
