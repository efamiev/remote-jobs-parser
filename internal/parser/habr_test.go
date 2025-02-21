package parser

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"remote-jobs-parser/test"
)

func TestParseHabr(t *testing.T) {
	html := test.ReadHTML("habr-page.html")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, html)
		return
	}))
	defer server.Close()

	t.Run("Returns jobs names", func(t *testing.T) {
		expectedResults := []string{"Golang-разработчик (Разработчик облачного оркестратора)", "Специалист по статическому анализу кода (Svace, AppSec)", "Ведущий системный инженер Linux", "Go Developer (SIEM KUMA)", "Tech Lead Golang в проект Deckhouse", "Ведущий Go разработчик (Платформа разработки)", "Ведущий SRE инженер (Платформа разработки)", "Стажёр Flutter-разработчик", "Lead Fullstack Developer (удаленно)", "Ruby / Go Developer", "Product Owner (slots)", "Senior PHP Программист", "Тимлид группы анализа клиентских данных и веб-аналитики/Senior Data Scientist", "Senior Go-разработчик", "Бизнес-аналитик (разработка, автоматизация)", "User acquisition manager", "Senior Golang Engineer", "Intern Golang Developer [Развитие инфраструктуры]", "Системный аналитик", "Site Reliability Engineer (SRE)", "Бизнес-аналитик", "Старший бизнес-аналитик", "Старший системный аналитик", "Go-pазработчик", " Middle/Senior Go Developer [Голосовая экосистема]"}

		results := make(chan []string, 2)

		client := &http.Client{}
		ParseHarb(results, client, server.URL+"/vacancies?q=go&sort=date&type=all")

		actualResults := []string{}

		for x := range results {
			actualResults = append(actualResults, x...)
		}

		assert.Equal(t, expectedResults, actualResults, "Results should match the mocked content")
	})
}
