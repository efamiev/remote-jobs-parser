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
	html := test.ReadHTML("hh-page.html")

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

		Start(server.URL+"/search/vacancy?text=%22go%22&salary=&professional_role=96&items_on_page=20", 1)
	})

	t.Run("Returns jobs names", func(t *testing.T) {
		expectedResults := []string{"Backend-разработчик (Golang и микросервисы)", "Senior Frontend-разработчик (JS, React)", "Junior Python Developer", "Junior backend-разработчик", "Senior/Middle+ Backend Developer", "Junior Python Developer", "Middle Golang разработчик", "Front-end разработчик (React.js) junior+", "Go Developer", "Backend-разработчик (Middle) в команду открытых данных", "Golang developer", "Backend-разработчик", "Go Backend Developer / Go бэкенд-разрабочтик", "Golang-разработчик", "Backend ishlab chiqaruvchisi (IT)", "Golang-разработчик", "Backend-разработчик", "Middle Backend-developer", "Middle Backend-разработчик на Go", "Java-разработчик"}

		actualResults := Start(server.URL+"/search/vacancy?text=%22go%22&salary=&professional_role=96&items_on_page=20", 1)
		assert.Equal(t, expectedResults, actualResults, "Results should match the mocked content")
	})
}
