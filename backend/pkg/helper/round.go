package helper

import "fmt"

func Round(f float64) int64 {
	return int64(f + 0.5)
}

func CalculatePersentase(data1, data2 float64) string {
	var presentase float64
	data2 = float64(Round(data2 / 1.11))
	data1 = float64(Round(data1 / 1.11))
	var selisih float64

	if data2 != 0 {
		selisih = float64(data1 - data2)
		presentase = (selisih / float64(data2)) * 100
	}

	strPresentase := fmt.Sprintf("%.2f%%", presentase)

	return strPresentase
}

func CalculatePersentaseSKU(data1, data2 float64) string {
	var presentase float64
	var selisih float64

	if data2 != 0 {
		selisih = float64(data1 - data2)
		presentase = (selisih / float64(data2)) * 100
	}

	strPresentase := fmt.Sprintf("%.2f%%", presentase)

	return strPresentase
}

func CalculatePersentaseTransaksi(data1, data2 int64) string {
	var presentase float64
	var selisih float64

	if data2 != 0 {
		selisih = float64(data1 - data2)
		presentase = (selisih / float64(data2)) * 100
	}

	strPresentase := fmt.Sprintf("%.2f%%", presentase)

	return strPresentase
}
