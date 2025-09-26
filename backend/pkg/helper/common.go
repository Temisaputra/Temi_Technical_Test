package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GetExcelColumnName(col int) string {
	columnName := ""
	for col >= 0 {
		columnName = string(rune('A'+(col%26))) + columnName
		col = col/26 - 1
	}
	return columnName
}

func Contains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

// func GetExcelColumnName(col int) string {
// 	columnName := ""
// 	for col >= 0 {
// 		fmt.Println("col", col)
// 		columnName = string(rune('A'+(col%26))) + columnName
// 		fmt.Println("columnName", columnName)
// 		col = col/26 - 1
// 	}
// 	return columnName
// }

func CheckAndPrintEmptyColumnPairs(headers []string) {
	for i := 0; i < len(headers); i++ {
		if headers[i] == "" && i > 0 && headers[i-1] != "" {
			// Cetak kolom kosong dan kolom kirinya yang ada isinya
			fmt.Printf("Empty header at column: %s (Left column with value: %s)\n",
				GetExcelColumnName(i+1), GetExcelColumnName(i))
		}
	}
}

// generate random voucher code
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// generate random voucher code
func GenerateVoucherCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
