package app

import (
	"cs50vn/virustracker/crawler/api/apprepository"
	"cs50vn/virustracker/crawler/api/apprepository/model"
	"cs50vn/virustracker/crawler/api/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func isValidDoc(doc []model.Item) bool {
	for _, item := range doc {
		if item.LeftName == "" || item.RightName == "" {
			return false
		}
	}
	return true
}

func RecreateMetadata(doc []model.Item) {
	//Delete current table data
	tx, err := apprepository.DbConnection.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	preparedInsert := utils.SQL_DELETE_MAPPING_NAME_TABLE

	preparedSt, err := apprepository.DbConnection.Prepare(preparedInsert)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = preparedSt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}

	//Insert new data
	preparedInsert = utils.SQL_INSERT_MAPPING_NAME_ITEM
	preparedSt, err = apprepository.DbConnection.Prepare(preparedInsert)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, item := range doc {
		_, err = preparedSt.Exec(item.LeftName, item.RightName)
	}

	tx.Commit()

}

///////////////////////////////////////////////////////////////////////////////
//Handler

func CreateNewMetaDataHandler(params ...string) string {
	var resultCode = 200
	var data = params[0]

	if len(data) <= 0 {
		resultCode = 400
		data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid request")
	} else {
		var arr []model.Item
		json.Unmarshal([]byte(data), &arr)

		if len(arr) == 0 {
			resultCode = 400
			data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid request")
		} else {
			if isValidDoc(arr) {
				//Recreate new metadata
				RecreateMetadata(arr)

				fmt.Println(arr)
				data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "OK")
			} else {
				resultCode = 400
				data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid request")
			}
		}
	}

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

/////////////////////////////////////////////////////////////////////////////////
////Rest endpoint - account profile section

func CreateNewMetaData(c *gin.Context) {

	dataJson, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var data = CreateNewMetaDataHandler(string(dataJson))
	c.String(200, data)
}