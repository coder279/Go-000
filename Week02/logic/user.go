package logic

import (
	"Week02/dao"
	"Week02/models"
	"database/sql"
	"github.com/pkg/errors"
)

func GetUserById(id int)(u models.User,err error){
	if u,err :=  dao.GetUserById(id); errors.Is(err,sql.ErrNoRows){
		return u, errors.Wrapf(err, "User not found")
	}
	return u ,nil
}