package app

import (
    "cs50vn/virustracker/crawler/api/apprepository"
    "cs50vn/virustracker/crawler/api/apprepository/model"
    "cs50vn/virustracker/crawler/api/utils"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "runtime"
    admin_app "cs50vn/virustracker/crawler/api/admin/app"
    "github.com/gin-gonic/gin"
    "strings"
    "strconv"
)

/////////////////////////////////////////////////////////////////////////////
//Global vars
func InitApp() {
    LoadConfig()
    InitDb()

    MemoryUsage()
}

func LoadConfig() {
    fmt.Println("Loaded config file: ", apprepository.Config.Dbname)

    data, err := ioutil.ReadFile(apprepository.ConfigName)
    if err != nil {
        fmt.Print(err)
    }

    err = json.Unmarshal(data, &apprepository.Config)
    if err != nil {
        fmt.Print(err)
    }

    fmt.Println(apprepository.Config)
}

func InitDb() {
    var err error
    apprepository.DbConnection, err = sql.Open("sqlite3", apprepository.Config.Dbname)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Println("Connected db: " + apprepository.Config.Dbname)

}

func MemoryUsage() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    // For info on each, see: https://golang.org/pkg/runtime/#MemStats
    fmt.Printf("Alloc = %v kB", (m.Alloc / 1024))
    fmt.Printf("\tTotalAlloc = %v kB\n", (m.TotalAlloc / 1024))
}

func LoadData() {
    //Clear and recreate new map
    apprepository.CountryList = make([]*model.Item, 0)

    fmt.Println("LoadMappingName()")
    prepareSelect := utils.SQL_GET_NEW_TRACKING_ITEM

    prepareSt, err := apprepository.DbConnection.Prepare(prepareSelect)
    if err != nil {
        fmt.Println(err.Error())
    }

    rows, _ := prepareSt.Query()

    for rows.Next() {
        var leftName string
        var rightName string
        var totalCases int64
        var totalDeaths int64
        var totalRecovered int64
        var seriousCases int64
        var totalCasesPer1Pop float64
        var totalDeathsPer1Pop float64
        var totalTests int64
        var testsPer1Pop float64
        var timestamp int64

        err = rows.Scan(&leftName, &rightName, &totalCases, &totalDeaths, &totalRecovered, &seriousCases, &totalCasesPer1Pop, &totalDeathsPer1Pop, &totalTests, &testsPer1Pop, &timestamp)
        if err != nil {
            fmt.Println(err.Error())
        } else {
            item := model.MakeItem(leftName, rightName, totalCases, totalDeaths, totalRecovered, seriousCases, totalCasesPer1Pop, totalDeathsPer1Pop, totalTests, testsPer1Pop, timestamp)
            apprepository.CountryList = append(apprepository.CountryList, item)
            fmt.Println(item)
        }
    }
}

///////////////////////////////////////////////////////////////////////////////
//Handler

func GetCurrentDataHandler(params ...string) string {
    LoadData()

    var resultCode = 200
    var data = ""

    for _, value := range apprepository.CountryList {
        data += fmt.Sprintf(utils.OBJECT_TEMPLATE,
            fmt.Sprintf(utils.ITEM_TEMPLATE, "leftName", value.LeftName)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "rightName", value.RightName)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalCases", value.TotalCases)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalDeaths", value.TotalDeaths)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalRecovered", value.TotalRecovered)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "seriousCases", value.SeriousCases)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "totalCasesPer1Pop", strconv.FormatFloat(value.TotalCasesPer1Pop, 'f', 2, 64 ))+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "totalDeathsPer1Pop", strconv.FormatFloat(value.TotalDeathsPer1Pop, 'f', 2, 64 ))+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalTests", value.TotalTests)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "testsPer1Pop", strconv.FormatFloat(value.TestsPer1Pop, 'f', 2, 64 ))+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "timestamp", value.Timestamp)) + ","
    }
    data = strings.TrimSuffix(data, ",")

    return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, fmt.Sprintf(utils.ARRAY_TEMPLATE, data))
}

/////////////////////////////////////////////////////////////////////////////////
////Rest endpoint - App section

func GetCurrentData(c *gin.Context) {
    var versionCode = c.Query("versionCode")

    var data = GetCurrentDataHandler(versionCode)
    c.String(200, data)
}

/////////////////////////////////////////////////////////////////////////////////
//Main func

func Run() {
    InitApp()

    r := gin.Default()

    v1 := r.Group("/v1")
    {
        //App
        v1.GET("/app/status", GetCurrentData)

        //*********************************************************************//
        //Admin
        //App
        v1.POST("/admin/app/metadata", admin_app.CreateNewMetaData)

    }

    r.Run(":" + apprepository.Config.Port)
}