package apprepository

import (
	"cs50vn/virustracker/crawler/api/apprepository/model"
	"database/sql"
)

//App
var ConfigName = "config.json"
var Config model.AppConfig
var DbConnection *sql.DB

var CountryList []*model.Item = make([]*model.Item, 0)

