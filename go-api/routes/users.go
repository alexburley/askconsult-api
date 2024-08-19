package routes

import (
	"encoding/json"
	"io"
	"net/http"
)

type MyData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func readData[T interface{}](w http.ResponseWriter, r *http.Request, obj T) (T, bool) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return obj, false
	}

	err = json.Unmarshal(body, &obj)
	if err != nil {
		http.Error(w, "Unable to parse JSON", http.StatusBadRequest)
		return obj, false
	}

	return obj, true
}

func Ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		// Handle the error, e.g., return a 500 Internal Server Error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		data, success := readData(w, r, MyData{})
		if !success {
			return
		}
		Ok(w, data)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
