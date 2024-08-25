package main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)


type Request struct {
	Data []string `json:"data"`
}


type Response struct {
	IsSuccess                bool     `json:"is_success"`
	UserID                   string   `json:"user_id"`
	Email                    string   `json:"email"`
	RollNumber               string   `json:"roll_number"`
	Numbers                  []string `json:"numbers"`
	Alphabets                []string `json:"alphabets"`
	HighestLowercaseAlphabet string   `json:"highest_lowercase_alphabet"`
}


type HardcodedResponse struct {
	OperationCode int `json:"operation_code"`
}


func processRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")


	if r.Method == http.MethodGet {

		hardcodedResponse := HardcodedResponse{
			OperationCode: 1,
		}
		jsonResponse, err := json.Marshal(hardcodedResponse)
		if err != nil {
			http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
	} else if r.Method == http.MethodPost {

		var requestData Request
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}


		response := Response{
			IsSuccess:  true,
			UserID:     "nikhil_jagyasi_03092003",   
			Email:      "jagyasi.nikhil2021@vitstudent.ac.in", 
			RollNumber: "21BBS0086",           
		}


		var numbers []string
		var alphabets []string
		highestLowercase := ""


		for _, item := range requestData.Data {
			if _, err := strconv.Atoi(item); err == nil {
	
				numbers = append(numbers, item)
			} else if isAlphabet(item) {
		
				alphabets = append(alphabets, item)
				
				if item == strings.ToLower(item) && (highestLowercase == "" || item > highestLowercase) {
					highestLowercase = item
				}
			}
		}


		response.Numbers = numbers
		response.Alphabets = alphabets
		response.HighestLowercaseAlphabet = highestLowercase


		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
	} else {

		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}


func isAlphabet(s string) bool {
	s = strings.ToLower(s)
	return s >= "a" && s <= "z"
}


func main() {
	

	fs := http.FileServer(http.Dir("."))
	http.Handle("/bfhl.html", fs)
	http.HandleFunc("/bfhl", processRequest) 


	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}