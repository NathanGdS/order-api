package helpers

func DecimalToCents(value float64) int64 {
	return int64(value * 100)
}

func CentsToDecimal(value int64) float64 {
	return float64(value) / 100
}
