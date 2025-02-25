package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"remote-jobs-parser/internal/parser"
	"remote-jobs-parser/tests/helpers"

	"github.com/stretchr/testify/assert"
)

func TestParseHH(t *testing.T) {
	html := helpers.ReadHTML("hh-page.html")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, html)
		return
	}))
	defer server.Close()

	t.Run("Add cookie device_magritte_breakpoint to request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, _ := r.Cookie("device_magritte_breakpoint")
			assert.Equal(t, "xxl", cookie.Value)
		}))
		defer server.Close()

		results := make(chan []parser.VacancyData, 2)

		client := &http.Client{}
		parser.ParseHH(results, client, server.URL+"/vacancies?q=go&sort=date&type=all")
	})

	t.Run("Returns jobs names", func(t *testing.T) {
		expectedResults := []parser.VacancyData{
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

		results := make(chan []parser.VacancyData, 2)

		client := &http.Client{}
		parser.ParseHH(results, client, server.URL+"/vacancies?q=go&sort=date&type=all")

		actualResults := []parser.VacancyData{}

		assert.Equal(t, "116873691", strings.Split(strings.Split("https://hh.ru/vacancy/116873691?query=%22go%22&hhtmFrom=vacancy_search_list", "/")[4], "?")[0])

		for x := range results {
			actualResults = append(actualResults, x...)
		}

		assert.Equal(t, expectedResults, actualResults, "Results should match the mocked content")
	})
}
