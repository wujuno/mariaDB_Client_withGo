package db

import (
	"database/sql"
	"mariadbClient/model"
)

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
		var phone1 sql.NullString
		var phone2 sql.NullString
		
		err := rows.Scan(
			&member.Id,
			&member.Name,
			&member.Number,
			&member.Addr,
			&phone1,
			&phone2,
			&member.Height,
			&member.DebutDate,
		)
		if err != nil {
			return nil, err
		}

		if phone1.Valid {
			member.Phone1 = phone1.String
		} else {
			member.Phone1 = ""
		}

		if phone2.Valid {
			member.Phone2 = phone2.String
		} else {
			member.Phone2 = ""
		}
		members = append(members, member)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}