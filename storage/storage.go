package storage

import (
	"github.com/s4kibs4mi/rapunzel-blog/models"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"gopkg.in/mgo.v2/bson"
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
	FindByID(ID bson.ObjectId) *models.User
	FindByUsername(username string) *models.User
	FindByEmail(username string) *models.User
	FindAll() []models.User
	FindAllByQuery(query protos.Query) []models.User
}

type SessionStorage interface {
	Init() bool
	SaveSession(session *models.Session) bool
	DeleteSession(session *models.Session) bool
	FindSessionByAccessToken(ID string) *models.Session
}

type PostStorage interface {
	Init() bool
	SavePost(post *models.Post) bool
	FindPostsByQuery(query []*protos.Query) []*models.Post
	FindPostByID(postID string) *models.Post
}

var mongodbStorage *MongodbStorage

func NewUserStorage() UserStorage {
	if mongodbStorage == nil {
		mongodbStorage = &MongodbStorage{}
	}
	return mongodbStorage
}

func NewSessionStorage() SessionStorage {
	if mongodbStorage == nil {
		mongodbStorage = &MongodbStorage{}
	}
	return mongodbStorage
}

func NewPostStorage() PostStorage {
	if mongodbStorage == nil {
		mongodbStorage = &MongodbStorage{}
	}
	return mongodbStorage
}
