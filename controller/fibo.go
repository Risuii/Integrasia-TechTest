package controller

import (
	"encoding/json"
	"integrasia/helpers"
	"net/http"
)

func Fibo(w http.ResponseWriter, r *http.Request) {
	var numb Fibonaci
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&numb); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
	}

	var result []int

	n := numb.StartNumb
	n1 := numb.StartNumb
	n2 := 2

	point := numb.ResponseNumb

	for i := 0; i < point; i++ {

		result = append(result, n)

		n = n1 + n2
		n2 = n1 // n1 akan menjadi angka sebelumnya
		n1 = n  // n akan menjadi angka sekarang

	}

	response := result
	helpers.ResponseJSON(w, http.StatusOK, response)
}
