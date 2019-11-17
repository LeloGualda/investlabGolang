package handles

import (
	"net/http"

	"../alphavantage"
)

func adminSerachAPI(w http.ResponseWriter, r *http.Request, keyword string) {
	returnStruct(w, r, alphavantage.GetAPi(keyword, false))
}
