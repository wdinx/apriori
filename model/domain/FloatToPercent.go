package domain

import (
	"fmt"
	"math"
)

func FloatToPercent(data float64) string {
	rounded := int(math.Round(data * 100))
	result := fmt.Sprint(rounded, "%") // tampilkan dengan tanda persen
	return result
}
