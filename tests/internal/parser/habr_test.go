package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"remote-jobs-parser/internal/parser"
	"remote-jobs-parser/tests/helpers"

	"github.com/stretchr/testify/assert"
)

func TestParseHabr(t *testing.T) {
	html := helpers.ReadHTML("habr-page.html")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, html)
		return
	}))
	defer server.Close()

	t.Run("Returns vacancies data", func(t *testing.T) {
		expectedResults := []parser.VacancyData{
			{Id: "1000153746", Company: "beeline", Title: "Golang-разработчик (Разработчик облачного оркестратора)", Link: "/vacancies/1000153746", Service: "habr"},
			{Id: "1000139501", Company: "Группа Астра", Title: "Специалист по статическому анализу кода (Svace, AppSec)", Link: "/vacancies/1000139501", Service: "habr"},
			{Id: "1000154044", Company: "Золотое Яблоко", Title: "Ведущий системный инженер Linux", Link: "/vacancies/1000154044", Service: "habr"},
			{Id: "1000154400", Company: "Лаборатория Касперского", Title: "Go Developer (SIEM KUMA)", Link: "/vacancies/1000154400", Service: "habr"},
			{Id: "1000155119", Company: "Флант", Title: "Tech Lead Golang в проект Deckhouse", Link: "/vacancies/1000155119", Service: "habr"},
			{Id: "1000154693", Company: "БЮРО 1440", Title: "Ведущий Go разработчик (Платформа разработки)", Link: "/vacancies/1000154693", Service: "habr"},
			{Id: "1000154696", Company: "БЮРО 1440", Title: "Ведущий SRE инженер (Платформа разработки)", Link: "/vacancies/1000154696", Service: "habr"},
			{Id: "1000152616", Company: "Яндекс", Title: "Стажёр Flutter-разработчик", Link: "/vacancies/1000152616", Service: "habr"},
			{Id: "1000154690", Company: "Apphud", Title: "Lead Fullstack Developer (удаленно)", Link: "/vacancies/1000154690", Service: "habr"},
			{Id: "1000153467", Company: "AMarkets", Title: "Ruby / Go Developer", Link: "/vacancies/1000153467", Service: "habr"},
			{Id: "1000154919", Company: "True Lab Game Studios", Title: "Product Owner (slots)", Link: "/vacancies/1000154919", Service: "habr"},
			{Id: "1000148082", Company: "Автомакон", Title: "Senior PHP Программист", Link: "/vacancies/1000148082", Service: "habr"},
			{Id: "1000155114", Company: "Inventive Retail Group", Title: "Тимлид группы анализа клиентских данных и веб-аналитики/Senior Data Scientist", Link: "/vacancies/1000155114", Service: "habr"},
			{Id: "1000155103", Company: "Флант", Title: "Senior Go-разработчик", Link: "/vacancies/1000155103", Service: "habr"},
			{Id: "1000155111", Company: "БФТ - Холдинг", Title: "Бизнес-аналитик (разработка, автоматизация)", Link: "/vacancies/1000155111", Service: "habr"},
			{Id: "1000154680", Company: "Grow Media", Title: "User acquisition manager", Link: "/vacancies/1000154680", Service: "habr"},
			{Id: "1000147548", Company: "Wanted.", Title: "Senior Golang Engineer", Link: "/vacancies/1000147548", Service: "habr"},
			{Id: "1000155106", Company: "МТС", Title: "Intern Golang Developer [Развитие инфраструктуры]", Link: "/vacancies/1000155106", Service: "habr"},
			{Id: "1000154663", Company: "БФТ - Холдинг", Title: "Системный аналитик", Link: "/vacancies/1000154663", Service: "habr"},
			{Id: "1000142037", Company: "Т-Банк", Title: "Site Reliability Engineer (SRE)", Link: "/vacancies/1000142037", Service: "habr"},
			{Id: "1000154383", Company: "БФТ - Холдинг", Title: "Бизнес-аналитик", Link: "/vacancies/1000154383", Service: "habr"},
			{Id: "1000154382", Company: "БФТ - Холдинг", Title: "Старший бизнес-аналитик", Link: "/vacancies/1000154382", Service: "habr"},
			{Id: "1000154380", Company: "БФТ - Холдинг", Title: "Старший системный аналитик", Link: "/vacancies/1000154380", Service: "habr"},
			{Id: "1000155084", Company: "Флант", Title: "Go-pазработчик", Link: "/vacancies/1000155084", Service: "habr"},
			{Id: "1000152829", Company: "МТС", Title: " Middle/Senior Go Developer [Голосовая экосистема]", Link: "/vacancies/1000152829", Service: "habr"}}

		results := make(chan []parser.VacancyData, 2)

		client := &http.Client{}
		parser.ParseHabr(results, client, server.URL+"/vacancies?q=go&sort=date&type=all")

		actualResults := []parser.VacancyData{}

		for x := range results {
			actualResults = append(actualResults, x...)
		}

		assert.Equal(t, expectedResults, actualResults, "Results should match the mocked content")
	})
}
