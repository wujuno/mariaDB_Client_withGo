package httpserver

import "net/http"



func idolListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//db에서 GET
	//JSON 형식의 데이터 응답
}