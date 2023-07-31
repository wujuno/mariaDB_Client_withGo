package db

import "mariadbClient/model"

func SelectIdolInfo() ([]model.Member, error){
	members := []model.Member{}

	q := `
		SELECT *
		FROM member
	`
	db, err := Init()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		member := model.Member{}
		err := rows.Scan(&member.Id, &member.Name, &member.Number, &member.Addr, &member.Phone1, &member.Phone2,  &member.Height, &member.DebutDate)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}