package parser

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"remote-jobs-parser/test"

	"github.com/stretchr/testify/assert"
)

func TestParseHH(t *testing.T) {
	html := test.ReadHTML("hh-page.html")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, html)
		return
	}))
	defer server.Close()

	t.Run("Returns jobs names", func(t *testing.T) {
		expectedResults := []string{"Backend-разработчик (Golang и микросервисы)", "Senior Frontend-разработчик (JS, React)", "Junior Python Developer", "Junior backend-разработчик", "Senior/Middle+ Backend Developer", "Junior Python Developer", "Middle Golang разработчик", "Front-end разработчик (React.js) junior+", "Go Developer", "Backend-разработчик (Middle) в команду открытых данных", "Golang developer", "Backend-разработчик", "Go Backend Developer / Go бэкенд-разрабочтик", "Golang-разработчик", "Backend ishlab chiqaruvchisi (IT)", "Golang-разработчик", "Backend-разработчик", "Middle Backend-developer", "Middle Backend-разработчик на Go", "Java-разработчик"}

		results := make(chan []string, 2)

		client := &http.Client{}
		ParseHH(results, client, server.URL+"/vacancies?q=go&sort=date&type=all")

		actualResults := []string{}

		for x := range results {
			actualResults = append(actualResults, x...)
		}

		assert.Equal(t, expectedResults, actualResults, "Results should match the mocked content")
	})
}
