// Package logic
// @Author twilikiss 2024/12/13 17:43:43
package logic

import (
	"strconv"
	"testing"
)

func TestMeasurement_GetMeasurement(t *testing.T) {
	data := new(Measurement).GetMeasurement("http://deviceshifu-plate-reader.shifu.com/get_measurement")
	_, err := strconv.ParseFloat(data, 64)
	if err != nil {
		t.Error(err)
	}
}
