package app

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"postman/internal/lib/httperrors"
	"strings"
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
		fmt.Print("Enter method: ")
		_ = scanner.Scan()
		method = scanner.Text()

		fmt.Print("Enter url: ")
		_ = scanner.Scan()
		url = scanner.Text()

		switch strings.ToUpper(method) {
		case "GET":
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
			fmt.Println("Enter request body")
			_ = scanner.Scan()
			req_body := scanner.Text()
			resp, err := client.Post(url, "application/json", bytes.NewBuffer([]byte(req_body)))
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

		case "PUT":
			fmt.Println("Enter request body")
			_ = scanner.Scan()
			req_body := scanner.Text()
			req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(req_body)))
			resp, err := client.Do(req)
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

		case "DELETE":
			fmt.Println("DELETE request")

			req, _ := http.NewRequest(http.MethodDelete, url, nil)

			resp, err := client.Do(req)
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

		default:
			fmt.Println("Method not allowed")

			fmt.Println("Press Enter to exit...")
			bufio.NewReader(os.Stdin).ReadString('\n')
		}
	}
}
