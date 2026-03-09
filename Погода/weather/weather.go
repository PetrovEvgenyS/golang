package weather

import (
	"demo/weather/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Package weather инкапсулирует обращение к сервису wttr.in и возврат строки с погодой.

// ErrWrongFormat возвращается, если указан неподдерживаемый формат вывода (не 1–4).
var ErrWrongFormat = errors.New("WRONG_FORMAT")

// GetWeather запрашивает у сервиса wttr.in строку с погодой по названию города.
// Параметр format пробрасывается как параметр "format" в запросе (см. документацию wttr.in).
func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrWrongFormat
	}

	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_URL")
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_HTTP")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_READBODY")
	}

	return string(body), nil
}
