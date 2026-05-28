package repositories

import (
	m "Financiering/Models"
	"fmt"

	"gorm.io/gorm"
)

func GetAllDossiers() []m.FinanceDossier {
	var Dossiers []m.FinanceDossier
	// result := GetDossiers()
	// result.Scan()
	return Dossiers
}

func GetDossiers() *gorm.DB {
	db := Database_Get()
	innerJoins := db.Table("m.Finance_Dossiers").Joins("left join budgets on budgets.id = m.Finance_Dossiers.budgets.id").Find(&m.FinanceDossier{})
	fmt.Println(innerJoins)

	return innerJoins
}

// shouldn't even be called if there is no budget
func ProcessPayment(limit float64, used float64, ID int) *gorm.DB {
	db := Database_Get()
	result := db.Table("budgets").Update("limit", limit).Update("used", used).Where("id = ?", ID)
	return result
}
