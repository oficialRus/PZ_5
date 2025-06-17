package main

import (
	"cadet_project/config"
	"cadet_project/services"
	"log"
)

func main() {
	///Загружаем конфигурацию (.env - > переменные окружения)
	config.LoadEnv()
	cadetSvc := services.NewCadetService()
	log.Println(cadetSvc.GetCadetInfo())

	// solo := models.NewCadet("Пётр", "курсант 4-го курса")
	// log.Println(solo.String())
}
