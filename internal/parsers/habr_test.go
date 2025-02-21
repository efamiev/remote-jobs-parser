package parsers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func readHTML() string {
	file, err := os.Open("/Users/famiev/projects/remote-jobs-parser/test/habr-page.html")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed reading data from file: %s", err)
	}

	content := string(data)
	return content
}

func TestParseHabr(t *testing.T) {
	html := readHTML()
	
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if r.URL.Path == "/search/vacancy" {
		// 	fmt.Fprintln(w, html)
		// 	return
		// }
		fmt.Fprintln(w, html)
		return 
	}))
	defer server.Close()

	// t.Run("Add cookie device_magritte_breakpoint to request", func(t *testing.T) {
	// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		cookie, _ := r.Cookie("device_magritte_breakpoint")
	// 		assert.Equal(t, "xxl", cookie.Value)
	// 	}))
	// 	defer server.Close()
	//
	// 	Start(server.URL+"/search/vacancy?text=%22go%22&salary=&professional_role=96&items_on_page=20", 1)
	// })

	t.Run("Returns jobs names", func(t *testing.T) {
		expectedResults := []string{"Backend-разработчик (Golang и микросервисы)", "Senior Frontend-разработчик (JS, React)", "Junior Python Developer", "Junior backend-разработчик", "Senior/Middle+ Backend Developer", "Junior Python Developer", "Middle Golang разработчик", "Front-end разработчик (React.js) junior+", "Go Developer", "Backend-разработчик (Middle) в команду открытых данных", "Golang developer", "Backend-разработчик", "Go Backend Developer / Go бэкенд-разрабочтик", "Golang-разработчик", "Backend ishlab chiqaruvchisi (IT)", "Golang-разработчик", "Backend-разработчик", "Middle Backend-developer", "Middle Backend-разработчик на Go", "Java-разработчик"}
		
		results := make(chan []string, 30)
		
		client := &http.Client{}
		ParseHarb(results, client, server.URL+"/search/vacancy?text=%22go%22&salary=&professional_role=96&items_on_page=20")

		actualResults := []string{}

		for x := range results {
			actualResults = append(actualResults, x...)
		}
		close(results)

		assert.Equal(t, expectedResults, actualResults, "Results should match the mocked content")
	})
}

