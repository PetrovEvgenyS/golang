package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Package geo отвечает за получение и валидацию информации о городе пользователя.

// GeoData содержит минимальные геоданные, необходимые для запроса погоды.
type GeoData struct {
	City string `json:"city"`
}

// CityPopulationResponce — сокращённый ответ сервиса,
// по которому определяем, существует ли указанный город.
type CityPopulationResponce struct {
	Error bool `json:"error"`
}

// ErrNoCity возвращается, если переданный город не найден во внешнем сервисе.
var ErrNoCity = errors.New("NOCITY")

// ErrNot200 возвращается, если сервис определения местоположения ответил не 200 OK.
var ErrNot200 = errors.New("NOT200")

// GetMyLocation определяет город пользователя.
// Если город передан явно, он сначала проверяется во внешнем API,
// иначе город запрашивается по IP через сервис ipapi.co.
func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			return nil, ErrNoCity
		}
		return &GeoData{
			City: city,
		}, nil
	}

	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, ErrNot200
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

// checkCity проверяет, существует ли город в базе сервиса countriesnow.space.
// При любой сетевой или JSON‑ошибке функция возвращает false.
func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})

	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var populationResponce CityPopulationResponce
	json.Unmarshal(body, &populationResponce)
	return !populationResponce.Error
}
