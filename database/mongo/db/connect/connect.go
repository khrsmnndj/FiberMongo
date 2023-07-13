package connect

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/khrsmnndj/FiberMongo/infra/config"
    "github.com/khrsmnndj/FiberMongo/database/mongo/db"
)

func MP() *mongo.Database {
	var dbMp *mongo.Database = db.Connect(config.MONGO_URI_MP, config.MONGO_DB_NAME_MP)
	return dbMp
}

func CMS() *mongo.Database {
	var dbCms *mongo.Database = db.Connect(config.MONGO_URI_CMS, config.MONGO_DB_NAME_CMS)
	return dbCms
}