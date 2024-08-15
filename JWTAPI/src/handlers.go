package main
import (
	"net/http"
	"encoding/json"
	"jwtapi/models"
)

func handleAuth(w http.ResponseWriter, r *http.Request){
	var creds models.Credentials

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&creds)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	is_correct := creds.CheckCreds()
	if is_correct{ 
		jwtString, err := creds.GetJwt()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": jwtString})

	}else{
		json.NewEncoder(w).Encode(map[string]string{"error":"Good luck mate"})
		w.WriteHeader(http.StatusUnauthorized)
	}
}
