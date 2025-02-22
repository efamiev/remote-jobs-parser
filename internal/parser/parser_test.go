package parser

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"remote-jobs-parser/test"
)

func TestStart(t *testing.T) {
	hhHTML := test.ReadHTML("hh-page.html")
	habrHTML := test.ReadHTML("habr-page.html")

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
		hhResults := []string{"Backend-разработчик (Golang и микросервисы)", "Senior Frontend-разработчик (JS, React)", "Junior Python Developer", "Junior backend-разработчик", "Senior/Middle+ Backend Developer", "Junior Python Developer", "Middle Golang разработчик", "Front-end разработчик (React.js) junior+", "Go Developer", "Backend-разработчик (Middle) в команду открытых данных", "Golang developer", "Backend-разработчик", "Go Backend Developer / Go бэкенд-разрабочтик", "Golang-разработчик", "Backend ishlab chiqaruvchisi (IT)", "Golang-разработчик", "Backend-разработчик", "Middle Backend-developer", "Middle Backend-разработчик на Go", "Java-разработчик"}

		habrResults := []string{"Golang-разработчик (Разработчик облачного оркестратора)", "Специалист по статическому анализу кода (Svace, AppSec)", "Ведущий системный инженер Linux", "Go Developer (SIEM KUMA)", "Tech Lead Golang в проект Deckhouse", "Ведущий Go разработчик (Платформа разработки)", "Ведущий SRE инженер (Платформа разработки)", "Стажёр Flutter-разработчик", "Lead Fullstack Developer (удаленно)", "Ruby / Go Developer", "Product Owner (slots)", "Senior PHP Программист", "Тимлид группы анализа клиентских данных и веб-аналитики/Senior Data Scientist", "Senior Go-разработчик", "Бизнес-аналитик (разработка, автоматизация)", "User acquisition manager", "Senior Golang Engineer", "Intern Golang Developer [Развитие инфраструктуры]", "Системный аналитик", "Site Reliability Engineer (SRE)", "Бизнес-аналитик", "Старший бизнес-аналитик", "Старший системный аналитик", "Go-pазработчик", " Middle/Senior Go Developer [Голосовая экосистема]"}
		expectedResults := append(habrResults, hhResults...)

		params := []ParserParams{
			{Service: "hh", Url: "/search/vacancy?text=%22go%22&salary=&professional_role=96&items_on_page=20"},
			{Service: "habr", Url: "/vacancies?page=4&type=all"},
		}

		actualResults := Start(params)

		assert.Equal(t, len(expectedResults), len(actualResults), "Results should match the mocked content")
		
		for i, el := range actualResults {
			assert.Containsf(t, expectedResults[i], el, "Element not found")
		}
	})
}
