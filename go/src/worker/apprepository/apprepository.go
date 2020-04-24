package apprepository

import (
	"cs50vn/virustracker/crawler/worker/apprepository/model"
	"database/sql"
)

//App
var ConfigName = "config.json"
var Config model.AppConfig
var DbConnection *sql.DB

var MappingCountryList map[string]string = make(map[string]string)
var CountryList []*model.Item = make([]*model.Item, 0)



