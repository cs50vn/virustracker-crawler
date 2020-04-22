package app

import (
    "cs50vn/virustracker/crawler/api/apprepository"
    "cs50vn/virustracker/crawler/api/utils"
    "database/sql"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "runtime"
    admin_app "cs50vn/virustracker/crawler/api/admin/app"
    "github.com/gin-gonic/gin"
)

/////////////////////////////////////////////////////////////////////////////
//Global vars
func InitApp() {
    LoadConfig()
    InitDb()

    MemoryUsage()
}

func LoadConfig() {
    fmt.Println("Loaded config file: ", apprepository.ConfigName)

    data, err := ioutil.ReadFile(apprepository.ConfigName)
    if err != nil {
        fmt.Print(err)
    }

    err = json.Unmarshal(data, &apprepository.Config)
    if err != nil {
        fmt.Print(err)
    }
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

///////////////////////////////////////////////////////////////////////////////
//Handler

func GetCurrentDataHandler(params ...string) string {
    var resultCode = 200
    var data = "GetCurrentDataHandler"

    return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
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
    r := gin.Default()

    v1 := r.Group("/v1")
    {
        //App
        v1.GET("/app/status", GetCurrentData)

        //*********************************************************************//
        //Admin
        //App
        v1.PATCH("/admin/app/metadata", admin_app.UpdateMetaData)

    }

    r.Run(":" + apprepository.Config.Port)
}