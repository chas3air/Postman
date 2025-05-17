package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"postman/internal/lib/httperrors"
	"time"
)

type App struct {
	log *slog.Logger
}

func New(log *slog.Logger) *App {
	return &App{
		log: log,
	}
}

func (a *App) Run() {
	var method string
	var url string
	scanner := bufio.NewScanner(os.Stdin)
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	for {
		fmt.Println("Enter method")
		_ = scanner.Scan()
		method = scanner.Text()

		fmt.Println("Enter url")
		_ = scanner.Scan()
		url = scanner.Text()

		switch method {
		case "GET":
			fmt.Println("GET request")
			resp, err := client.Get(url)
			if err != nil {
				fmt.Println("Error sending request")
				fmt.Println("Error:", err.Error())

				fmt.Println("Press Enter to exit...")
				bufio.NewReader(os.Stdin).ReadString('\n')
			}

			if resp.StatusCode >= 399 {
				httperrors.ErrorHandler(resp.StatusCode)

				fmt.Println("Press Enter to exit...")
				bufio.NewReader(os.Stdin).ReadString('\n')
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body")
				fmt.Println("Error:", err.Error())

				fmt.Println("Press Enter to exit...")
				bufio.NewReader(os.Stdin).ReadString('\n')
			}

			var jsonData interface{}
			if err := json.Unmarshal(body, &jsonData); err != nil {
				fmt.Println("Error parsing JSON:", err)
			} else {
				formattedJSON, err := json.MarshalIndent(jsonData, "", "  ")
				if err != nil {
					fmt.Println("Error formatting JSON:", err)
				} else {
					fmt.Println(string(formattedJSON))
				}
			}

			fmt.Println("Press Enter to exit...")
			bufio.NewReader(os.Stdin).ReadString('\n')

		case "POST":
			fmt.Println("POST request")
		case "PUT":
			fmt.Println("PUT request")
		case "DELETE":
			fmt.Println("DELETE request")
		default:
			fmt.Println("Method not allowed")
		}
	}
}
