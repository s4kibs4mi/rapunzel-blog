package conn

import (
	"github.com/spf13/viper"
	"fmt"
	"gopkg.in/mgo.v2"
)

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

var mSession *mgo.Session
var mDatabase *mgo.Database
var mConnectError error

func NewMongodbConnection() bool {
	mSession, mConnectError = mgo.Dial(viper.GetString("databases.mongodb.uri"))
	if mConnectError != nil {
		fmt.Printf("Couldn't connect to database [ %s/%s ]\n", viper.GetString("databases.mongodb.uri"),
			viper.GetString("databases.mongodb.name"))
		return false
	}
	mSession.SetMode(mgo.Monotonic, true)
	mDatabase = mSession.DB(viper.GetString("databases.mongodb.name"))
	return true
}

func GetMongoDB() *mgo.Database {
	return mDatabase
}

func GetUserCollection() *mgo.Collection {
	return GetMongoDB().C(viper.GetString("databases.mongodb.collections.user"))
}

func GetSessionCollection() *mgo.Collection {
	return GetMongoDB().C(viper.GetString("databases.mongodb.collections.session"))
}

func GetPostCollection() *mgo.Collection {
	return GetMongoDB().C(viper.GetString("databases.mongodb.collections.post"))
}

func GetCommentCollection() *mgo.Collection {
	return GetMongoDB().C(viper.GetString("databases.mongodb.collections.comment"))
}
