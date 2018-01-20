package storage

import (
	"github.com/s4kibs4mi/rapunzel-blog/models"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
)

/**
 * := Coded with love by Sakib Sami on 20/1/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

type UserStorage interface {
	Init() bool
	Save(user models.User) bool
	Update(user models.User) bool
	Delete(user models.User) bool
	Count() int
	FindByID(ID string) *models.User
	FindByUsername(username string) *models.User
	FindByEmail(username string) *models.User
	FindAll() []models.User
	FindAllByQuery(query protos.Query) []models.User
}

var mongodbStorage *MongodbStorage

func NewUserStorage() UserStorage {
	if mongodbStorage == nil {
		mongodbStorage = &MongodbStorage{}
	}
	return mongodbStorage
}
