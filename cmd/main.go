package main

import "hh-scraper/internal/parsers"

// import "../internal/parsers"

func main() {
	//TODO:  Сделать http сервер, который будет возвращаться список вакансий постранично. Так же добавить редис для кэширования результатов
	// Далее добавить сокеты для отслеживания прогресса собирания вакансий
	// Настроить nginx на vps и сделать CI/CD через github

	parsers.Parse()
}

