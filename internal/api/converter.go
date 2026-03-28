package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	LengthURL      = "http://localhost:8080/lengths"
	WeightURL      = "http://localhost:8080/weights"
	TemperatureURL = "http://localhost:8080/temperatures"
)

type UnitConverter struct {
	FromUnit string  `json:"fromUnit"`
	ToUnit   string  `json:"toUnit"`
	Val      float32 `json:"value"`
	Ans      float32 `json:"ans"`
}

func NewUnitConverter(funit, tunit string, v, a float32) *UnitConverter {
	return &UnitConverter{
		FromUnit: funit,
		ToUnit:   tunit,
		Val:      v,
		Ans:      a,
	}
}

func ConversionRequest(url string, uc *UnitConverter) (*UnitConverter, error) {
	jsonData, err := json.Marshal(uc)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	client := &http.Client{Timeout: 1 * time.Second}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server error: status %d", resp.StatusCode)
	}

	var result UnitConverter
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
