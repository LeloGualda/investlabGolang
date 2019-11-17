package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func returnStruct(w http.ResponseWriter, r *http.Request, face interface{}) {
	json, _ := json.Marshal(face)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
