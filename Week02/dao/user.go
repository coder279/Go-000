package dao

import (
	"Week02/models"
	"database/sql"
)

func GetUserById(id int)(u models.User,err error){
	sqlStr := "select id, name from user where id=?"
	sql.ErrNoRows = db.QueryRow(sqlStr, id).Scan(&u.Id, &u.Name)
	return


}
