package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"remote-jobs-parser/internal/parser"
	"remote-jobs-parser/tests/helpers"
)

func TestStart(t *testing.T) {
	hhHTML := helpers.ReadHTML("hh-page.html")
	habrHTML := helpers.ReadHTML("habr-page.html")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("PATH", r.URL.Path)
		if r.URL.Path == "/search/vacancy" {
			fmt.Fprintln(w, hhHTML)
		} else {
			fmt.Fprintln(w, habrHTML)
		}
	}))
	defer server.Close()

	t.Run("Returns jobs names", func(t *testing.T) {
		hhResults := []parser.VacancyData{
			{Id: "117345106", Company: "Zero Agency", Title: "Backend-разработчик (Golang и микросервисы)", Link: "https://hh.ru/vacancy/117345106?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117536524", Company: "ООО\u00a0GlowByte", Title: "Senior Frontend-разработчик (JS, React)", Link: "https://hh.ru/vacancy/117536524?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117530627", Company: "Multilogin Software Ltd.", Title: "Junior Python Developer", Link: "https://hh.ru/vacancy/117530627?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117347051", Company: "enKod", Title: "Junior backend-разработчик", Link: "https://hh.ru/vacancy/117347051?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117217647", Company: "Версус Сервис", Title: "Senior/Middle+ Backend Developer", Link: "https://hh.ru/vacancy/117217647?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117532131", Company: "Multilogin Software Ltd.", Title: "Junior Python Developer", Link: "https://hh.ru/vacancy/117532131?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117532180", Company: "Multilogin Software Ltd.", Title: "Middle Golang разработчик", Link: "https://hh.ru/vacancy/117532180?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117493725", Company: "ООО\u00a0Ботхаб", Title: "Front-end разработчик (React.js) junior+", Link: "https://hh.ru/vacancy/117493725?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117459539", Company: "ТОО\u00a0Неткрэкер Текнолоджи Казахстан", Title: "Go Developer", Link: "https://hh.ru/vacancy/117459539?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117398331", Company: "Дром", Title: "Backend-разработчик (Middle) в команду открытых данных", Link: "https://hh.ru/vacancy/117398331?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117473797", Company: "ООО\u00a0Триада", Title: "Golang developer", Link: "https://hh.ru/vacancy/117473797?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117478320", Company: "ООО\u00a0Диджитал Форс", Title: "Backend-разработчик", Link: "https://hh.ru/vacancy/117478320?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117464371", Company: "ООО\u00a0Зигмунд Онлайн", Title: "Go Backend Developer / Go бэкенд-разрабочтик", Link: "https://hh.ru/vacancy/117464371?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117528695", Company: "Delaweb", Title: "Golang-разработчик", Link: "https://hh.ru/vacancy/117528695?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117280251", Company: "ООО\u00a0Содействие", Title: "Backend ishlab chiqaruvchisi (IT)", Link: "https://hh.ru/vacancy/117280251?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117251443", Company: "Orion soft", Title: "Golang-разработчик", Link: "https://hh.ru/vacancy/117251443?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117243051", Company: "Rasulova Nisso", Title: "Backend-разработчик", Link: "https://hh.ru/vacancy/117243051?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117532518", Company: "ООО\u00a0РБ", Title: "Middle Backend-developer", Link: "https://hh.ru/vacancy/117532518?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "117486322", Company: "CubeLab", Title: "Middle Backend-разработчик на Go", Link: "https://hh.ru/vacancy/117486322?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"},
			{Id: "116873691", Company: "ЭР-Телеком", Title: "Java-разработчик", Link: "https://hh.ru/vacancy/116873691?query=%22go%22&hhtmFrom=vacancy_search_list", Service: "hh"}}

		habrResults := []parser.VacancyData{
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

		expectedResults := append(habrResults, hhResults...)

		params := []parser.ParserParams{
			{Service: "hh", Url: server.URL + "/search/vacancy?text=%22go%22&salary=&professional_role=96&items_on_page=20"},
			{Service: "habr", Url: server.URL + "/vacancies?page=4&type=all"},
		}

		actualResults := parser.Start(params)

		assert.Equal(t, len(expectedResults), len(actualResults), "Results should match the mocked content")

		for _, el := range actualResults {
			assert.Contains(t, expectedResults, el, "Element not found")
		}
	})
}
