package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

// main — точка входа CLI‑приложения.
// Разбирает параметры командной строки, определяет город
// и выводит строку с погодой в выбранном формате.
func main() {
	fmt.Println("Погода")

	// Флаг city позволяет явно указать город (иначе он будет получен по IP).
	city := flag.String("city", "", "Город пользователя")
	// Флаг format отвечает за формат строки ответа сервиса wttr.in (1–4).
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println("ошибка определения местоположения:", err.Error())
		return
	}

	weatherData, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Println("ошибка получения погоды:", err.Error())
		return
	}
	fmt.Println(weatherData)
}
