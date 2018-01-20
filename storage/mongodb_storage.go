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

type MongodbStorage struct {
}

func (db *MongodbStorage) Init() bool {
	return false
}

func (db *MongodbStorage) Save(user models.User) bool {
	return false
}

func (db *MongodbStorage) Update(user models.User) bool {
	return false
}

func (db *MongodbStorage) Delete(user models.User) bool {
	return false
}

func (db *MongodbStorage) FindByID(ID string) *models.User {
	return nil
}

func (db *MongodbStorage) FindByUsername(username string) *models.User {
	return nil
}

func (db *MongodbStorage) FindByEmail(username string) *models.User {
	return nil
}

func (db *MongodbStorage) FindAll() []models.User {
	return []models.User{}
}

func (db *MongodbStorage) FindAllByQuery(query protos.Query) []models.User {
	return []models.User{}
}
