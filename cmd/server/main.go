package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	// Global middleware
	router.Use(middleware.Logger)    // Logs all incoming request to terminal
	router.Use(middleware.Recoverer) // Prevents the server from crashing if the code panics

	// Route 1: Render the main page dashboard
	router.Get("/", func(responseWriter http.ResponseWriter, incomingRequest *http.Request) {
		// Parse the core layout shell
		htmlTemplate, parsingError := template.ParseFiles(
			"internal/templates/layouts/base.html",
			"internal/templates/pages/index.html",
		)

		// Explicitly check if the file paths are wrong or broken

		if parsingError != nil {
			http.Error(responseWriter, parsingError.Error(), http.StatusInternalServerError)
			return
		}

		// Merge the templates and write them back to the user's browser window

		htmlTemplate.Execute(responseWriter, nil)
	})

	// Route 2: The HTMX endpoint that intercepts the weather form submission

	router.Post("/query", func(responseWriter http.ResponseWriter, incomingRequest *http.Request) {
		// Extraxt the value typed into the <input name="station"> HTML field

		requestedStation := incomingRequest.FormValue("station")

		// Fallback safely if the user submitted an empyt input box
		if requestedStation == "" {
			requestedStation = "Unknown Station"
		}

		// Build a raw HTML string snipped and write it directly to the response
		// stream

		fmt.Fprintf(responseWriter, `
    <div class="bg-blue-50 border border-blue-200 p-4 rounded-xl">
        <h3 class="text-lg font-bold text-blue-800">Results for %s</h3>
        <p class="text-blue-600 mt-1">Go backend dynamically injected this without a page reload!</p>
    </div>
	`, requestedStation)

	})

	// Open the network gates and start listening for web traffic
	serverPort := ":8080"
	fmt.Println("Server is running on http://localhost" + serverPort)

	listeningError := http.ListenAndServe(serverPort, router)
	if listeningError != nil {
		fmt.Println("Error starting the network server:", listeningError)
	}

}
