package handlers

import (
	r "Financiering/Repositories"
	"net/http"
	// "strconv"
)

func HomePageHandler(wr http.ResponseWriter, rq *http.Request) {
	LoadTemplate(wr, "Templates/Home.gohtml", r.GetAllDossiers())
}

func AddorRemovePageHandler(wr http.ResponseWriter, rq *http.Request) {
	LoadTemplate(wr, "Templates/AddorRemove.gohtml", nil)
}

// func DossierPageHandler(wr http.ResponseWriter, rq *http.Request) {
// 	DossierID, err := strconv.Atoi(rq.PathValue("dossierID"))
// 	if err != nil {
// 		fmt.Println("DossierPageHandler", err)
// 		return
// 	}
// 	LoadTemplate(wr, "Templates/DetailDossier.gohtml", r.GetDossierbyID(DossierID))
// }

// func ViewBudgetPage(wr http.ResponseWriter, rq *http.Request) {
// 	ClientID, err := strconv.Atoi(rq.PathValue("ClientID"))
// 	if err != nil {
// 		fmt.Println("ViewBudgetPage: ", err)
// 		HomePageHandler(wr, rq)
// 		return
// 	}
// 	LoadTemplate(wr, "Templates/ViewClientBudget.gohtml", r.GetClientBudget(ClientID))
// }
