package dao

import (
	"Week02/models"
	"database/sql"
	"fmt"
)

func GetUserById(id int)(u *models.User,err error){
	u = new(models.User)
	sqlStr := "select id, name from user where id=?"
	err = db.QueryRow(sqlStr, id).Scan(u.Id, u.Name)
	fmt.Println(id)
	return u,sql.ErrNoRows


}
