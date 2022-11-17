package model

import "pro/app/models"

func AdvanceQueryTwoSQL() ([]models.AdvanceQueryTwoInput, error) {
	query := "select a1.Region, a1.Smoking_prevalence  from (select Region,Smoking_prevalence from Health where Smoking_prevalence > (select avg(Smoking_prevalence)from Health where Smoking_prevalence <> 0))as a1 NATURAL JOIN  ( select Region,max(Current_ventilator_patients) as ans from Hospitalizations  group by Region  )as b1 limit 15"
	var result = make([]models.AdvanceQueryTwoInput, 0)
	advanceQuery := models.AdvanceQueryTwoInput{}
	rows, err := Db.Query(query)
	for rows.Next() {
		err = rows.Scan(&advanceQuery.Region, &advanceQuery.Smoking_prevalence)
		checkErr(err)
		result = append(result, advanceQuery)
	}
	return result, nil
}
