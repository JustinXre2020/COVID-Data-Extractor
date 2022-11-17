package model

import "pro/app/models"

func AdvanceQueryOneSQL() ([]models.AdvanceQueryOneInput, error) {
	query := "select count(p.PatientId), p.Region 
			  from Patients as p natural join Vaccinations as v 
			  where v.Fully_vaccinated=1 group by p.Region limit 15"
	var result = make([]models.AdvanceQueryOneInput, 0)
	advanceQuery := models.AdvanceQueryOneInput{}
	rows, err := Db.Query(query)
	for rows.Next() {
		err = rows.Scan(&advanceQuery.Count, &advanceQuery.Region)
		checkErr(err)
		result = append(result, advanceQuery)
	}
	return result, nil
}
