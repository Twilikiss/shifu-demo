// Package logic
// @Author twilikiss 2024/12/13 17:27:27
package logic

import (
	"context"
	"io"
	"math/big"
	"net/http"
	"shifu-demo/log"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
}

func NewMeasurement() *Measurement {
	return &Measurement{}
}

func (ms *Measurement) GetMeasurement(url string) string {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()
	path := url
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		log.Error("get measurement error", err)
		return ""
	}
	client := http.DefaultClient
	httpRsp, err := client.Do(httpReq)
	if err != nil {
		log.Error("http.DefaultClient error", err)
		return ""
	}
	defer httpRsp.Body.Close()
	rspBody, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		log.Error("io.ReadAll error", err)
		return ""
	}

	split := strings.Split(string(rspBody), "\n")
	var data []string
	for _, v := range split {
		temp := strings.Split(v, " ")
		data = append(data, temp...)
	}
	var p_data float64
	var count int
	for _, v := range data {
		if IsNumeric(v) {
			s, _ := strconv.ParseFloat(v, 64)
			p_data += s
			count++
		}
	}
	d1 := big.NewFloat(p_data)
	d2 := big.NewFloat(float64(count))
	result := new(big.Float)
	result.Quo(d1, d2)
	return result.String()
}
func IsNumeric(number string) bool {
	_, err := strconv.ParseFloat(number, 64)
	return err == nil
}
