package Test

import (
	"fmt"
	"strconv"
)

//CalculaMedia fáz o obvio
func CalculaMedia(num ...float64) float64 {
	sum := 0.0
	for _, n := range num {
		sum += n
	}
	media := sum / float64(len(num))
	MediaAred, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return MediaAred
}
