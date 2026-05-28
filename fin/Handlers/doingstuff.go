package handlers

import (
	r "Financiering/Repositories"
	"net/http"
)

func DoingStuff(wr http.ResponseWriter, rq *http.Request) {
	r.GetDossiers()
}
