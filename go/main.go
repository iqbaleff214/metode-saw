package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Dataset struct {
	Criteria     []*Criterion
	Alternatives []*Alternative
}

type Criterion struct {
	Code             string  `json:"code"`
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	Weight           float64 `json:"weight"`
	NormalizedWeight float64
	DivisorValue     float64
}

type Alternative struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	C1           float64 `json:"C1"`
	NormalizedC1 float64
	C2           float64 `json:"C2"`
	NormalizedC2 float64
	C3           float64 `json:"C3"`
	NormalizedC3 float64
	C4           float64 `json:"C4"`
	NormalizedC4 float64
	C5           float64 `json:"C5"`
	NormalizedC5 float64
	Result       float64
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
		return dataset.Alternatives[i].Result > dataset.Alternatives[j].Result
	})

	// menampilkan hasil menggunakan std out CLI
	fmt.Println("================================")
	fmt.Println("No.\tNama \t\tHasil")
	fmt.Println("================================")
	for i, a := range dataset.Alternatives {
		fmt.Printf("%2d\t%-10s \t %.2f\n", i+1, a.Name, a.Result)
	}
}

// method untuk menormalisasi nilai bobot
func (dataset *Dataset) weightNormalization() {
	// Menghitung total bobot
	var total float64
	for _, c := range dataset.Criteria {
		total += c.Weight
	}

	// Menghitung hasil normalisasi masing-masing bobot
	for _, c := range dataset.Criteria {
		c.NormalizedWeight = c.Weight / total
	}
}

// method untuk menormalisasi nilai matriks
func (dataset *Dataset) matrixNormalization() {
	// menentukan nilai pembagi untuk masing-masing kriteria berdasarkan nilai-nilai alternatif
	for _, a := range dataset.Alternatives {
		for _, c := range dataset.Criteria {
			switch c.Code {
			case "C1":
				c.DivisorValue = divisorValue(a.C1, c.DivisorValue, c.Type)
			case "C2":
				c.DivisorValue = divisorValue(a.C2, c.DivisorValue, c.Type)
			case "C3":
				c.DivisorValue = divisorValue(a.C3, c.DivisorValue, c.Type)
			case "C4":
				c.DivisorValue = divisorValue(a.C4, c.DivisorValue, c.Type)
			case "C5":
				c.DivisorValue = divisorValue(a.C5, c.DivisorValue, c.Type)
			}
		}
	}

	// normalisasi matriks berdasarkan nilai pembaginya
	for _, a := range dataset.Alternatives {
		for _, c := range dataset.Criteria {
			switch c.Code {
			case "C1":
				a.NormalizedC1 = normalize(a.C1, c.DivisorValue, c.Type)
			case "C2":
				a.NormalizedC2 = normalize(a.C2, c.DivisorValue, c.Type)
			case "C3":
				a.NormalizedC3 = normalize(a.C3, c.DivisorValue, c.Type)
			case "C4":
				a.NormalizedC4 = normalize(a.C4, c.DivisorValue, c.Type)
			case "C5":
				a.NormalizedC5 = normalize(a.C5, c.DivisorValue, c.Type)
			}
		}
	}
}

// method untuk menghitung pembobotan matriks yang telah dinormalisasi
func (dataset *Dataset) resultCalculation() {
	for _, a := range dataset.Alternatives {
		for _, c := range dataset.Criteria {
			switch c.Code {
			case "C1":
				a.Result += (a.NormalizedC1 * c.NormalizedWeight)
			case "C2":
				a.Result += (a.NormalizedC2 * c.NormalizedWeight)
			case "C3":
				a.Result += (a.NormalizedC3 * c.NormalizedWeight)
			case "C4":
				a.Result += (a.NormalizedC4 * c.NormalizedWeight)
			case "C5":
				a.Result += (a.NormalizedC5 * c.NormalizedWeight)
			}
		}
	}
}

func normalize(matrix, divisor float64, criteriaType string) float64 {
	if criteriaType == "cost" {
		return divisor / matrix
	}

	return matrix / divisor
}

func divisorValue(value, initial float64, criteriaType string) float64 {
	if criteriaType == "cost" && (value < initial || initial == 0) {
		return value
	}

	if criteriaType == "benefit" && value > initial {
		return value
	}

	return initial
}
