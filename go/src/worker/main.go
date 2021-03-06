package main

import (
    "bytes"
    "cs50vn/virustracker/crawler/worker/apprepository"
    "cs50vn/virustracker/crawler/worker/apprepository/model"
    "cs50vn/virustracker/crawler/worker/utils"
    "database/sql"
    "encoding/json"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/jasonlvhit/gocron"
    _ "github.com/mattn/go-sqlite3"
    "io/ioutil"
    "net/http"
    "strconv"
    "strings"
    "time"
)

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

func LoadMappingName() {
    //Clear and recreate new map
    apprepository.MappingCountryList = make(map[string]string)

    fmt.Println("LoadMappingName()")
    prepareSelect := utils.SQL_GET_ALL_MAPPING_NAME

    prepareSt, err := apprepository.DbConnection.Prepare(prepareSelect)
    if err != nil {
        fmt.Println(err.Error())
    }

    rows, _ := prepareSt.Query()

    for rows.Next() {
        var leftName string
        var rightName string

        err = rows.Scan(&leftName, &rightName)
        if err != nil {
            fmt.Println(err.Error())
        } else {
            apprepository.MappingCountryList[strings.TrimSpace(leftName)] = rightName
        }
    }

    fmt.Println(apprepository.MappingCountryList)
}

func ProcessJob() {
    LoadMappingName()
    apprepository.CountryList = make([]*model.Item, 0)
    //
    //Call url to process
    //res, err := http.Get("https://worldometers.info/coronavirus")
    res, err := http.Get("https://" + apprepository.Config.CrawUrl)
    if err != nil {
        fmt.Println(err.Error())
    }
    defer res.Body.Close()

    //Parse data
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        fmt.Println(err.Error())
    }
    now := time.Now().Unix()

    count := 0
    doc.Find("table#main_table_countries_today tbody tr").Each(func(i int, s *goquery.Selection) {
        nodes := s.Find("td")
        firstNode := nodes.Eq(1)

        fmt.Println(firstNode.Text())

        //fmt.Println(node.Text())
        if value, ok := apprepository.MappingCountryList[strings.TrimSpace(firstNode.Text())]; ok {
            count++
            leftName := nodes.Eq(1).Text()
            rightName := value
            totalCases, err := strconv.ParseInt(strings.TrimSpace(strings.ReplaceAll(nodes.Eq(2).Text(), ",", "")), 10, 64)
            if err != nil {
                totalCases = 0
            }
            totalDeaths, err := strconv.ParseInt(strings.TrimSpace(strings.ReplaceAll(nodes.Eq(4).Text(), ",", "")), 10, 64)
            if err != nil {
                totalDeaths = 0
            }
            totalRecovered, err := strconv.ParseInt(strings.TrimSpace(strings.ReplaceAll(nodes.Eq(6).Text(), ",", "")), 10, 64)
            if err != nil {
                totalRecovered = 0
            }
            seriousCases, err := strconv.ParseInt(strings.TrimSpace(strings.ReplaceAll(nodes.Eq(8).Text(), ",", "")), 10, 64)
            if err != nil {
                seriousCases = 0
            }
            totalCasesPer1Pop, err := strconv.ParseFloat(strings.TrimSpace(strings.ReplaceAll(nodes.Eq(9).Text(), ",", "")), 32)
            if err != nil {
                totalCasesPer1Pop = 0.0
            }
            totalCasesPer1Pop = float64(int(totalCasesPer1Pop * 100)) / 100

            totalDeathsPer1Pop, err := strconv.ParseFloat(strings.TrimSpace(strings.ReplaceAll(nodes.Eq(10).Text(), ",", "")), 32)
            if err != nil {
                totalDeathsPer1Pop = 0.0
            }
            totalDeathsPer1Pop = float64(int(totalDeathsPer1Pop * 100)) / 100

            totalTests, err := strconv.ParseInt(strings.TrimSpace(strings.ReplaceAll(nodes.Eq(11).Text(), ",", "")), 10, 64)
            if err != nil {
                totalTests = 0
            }
            testsPer1Pop, err := strconv.ParseFloat(strings.TrimSpace(strings.ReplaceAll(nodes.Eq(12).Text(), ",", "")), 32)
            if err != nil {
                testsPer1Pop = 0.0
            }
            testsPer1Pop = float64(int(testsPer1Pop * 100)) / 100

            timestamp := now

            apprepository.CountryList = append(apprepository.CountryList, model.MakeItem(leftName, rightName, totalCases, totalDeaths, totalRecovered, seriousCases, totalCasesPer1Pop, totalDeathsPer1Pop, totalTests, testsPer1Pop, timestamp))

        }
    })

    fmt.Println("Total nodes: ", count)

    for _, item := range apprepository.CountryList {
        fmt.Println(item)
    }

    //Write data to db
    tx, err := apprepository.DbConnection.Begin()
    if err != nil {
        fmt.Println(err.Error())
    }
    preparedInsert := utils.SQL_INSERT_TRACKING_ITEM

    preparedSt, err := apprepository.DbConnection.Prepare(preparedInsert)
    if err != nil {
        fmt.Println(err.Error())
    }

    for _, item := range apprepository.CountryList {
        _, err := preparedSt.Exec(item.LeftName, item.RightName, item.TotalCases, item.TotalDeaths, item.TotalRecovered, item.SeriousCases, item.TotalCasesPer1Pop, item.TotalDeathsPer1Pop, item.TotalTests, item.TestsPer1Pop, item.Timestamp)
        if err != nil {
            fmt.Println(err.Error())
        }
    }

    tx.Commit()

}

func SendUpdate() {
    requestBody, err := json.Marshal(apprepository.CountryList)
    if err != nil {
        fmt.Println(err.Error())
    }

    timeout := time.Duration(5 * time.Second)
    client := http.Client{
        Timeout: timeout,
    }
    fmt.Println(string(requestBody))
    request, err := http.NewRequest("POST", "https://" + apprepository.Config.HookUrl, bytes.NewBuffer(requestBody))
    request.Header.Set("Content-type", "application/json")
    if err != nil {
        fmt.Println(err.Error())
		return
    }

    response, err := client.Do(request)
    if err != nil {
        fmt.Println(err.Error())
		return
    }

    defer response.Body.Close()

    result, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err.Error())
		return
    }

    fmt.Println(string(result))

}

func HandleJob() {
    fmt.Println("I am running task.")
    ProcessJob()
    fmt.Println("Send job")
	SendUpdate()
}

func InitApp() {
    LoadConfig()
    InitDb()

    gocron.Every(apprepository.Config.Times).Seconds().Do(HandleJob)
}

func main() {
    InitApp()
    //HandleJob()

    // Start all the pending jobs
    <-gocron.Start()
}
