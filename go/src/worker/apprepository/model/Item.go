package model

type Item struct {
    LeftName string `json: "leftName"`
    RightName string `json: "rightName"`
    TotalCases int64 `json: "totalCases"`
    TotalDeaths int64 `json: "totalDeaths"`
    TotalRecovered int64 `json: "totalRecovered"`
    SeriousCases int64 `json: "seriousCases"`
    TotalCasesPer1Pop float64 `json: "totalCasesPer1Pop"`
    TotalDeathsPer1Pop float64 `json: "totalDeathsPer1Pop"`
    TotalTests int64 `json: "totalTests"`
    TestsPer1Pop float64 `json: "testsPer1Pop"`
    Timestamp int64 `json: "timestamp"`
}

func MakeItem(leftName string, rightName string, totalCases int64, totalDeaths int64, totalRecovered int64, seriousCases int64, totalCasesPer1Pop float64, totalDeathsPer1Pop float64, totalTests int64, testsPer1Pop float64, timestamp int64) *Item {
    return &Item{LeftName: leftName, RightName: rightName, TotalCases: totalCases, TotalDeaths: totalDeaths, TotalRecovered: totalRecovered, SeriousCases: seriousCases, TotalCasesPer1Pop: totalCasesPer1Pop, TotalDeathsPer1Pop: totalDeathsPer1Pop, TotalTests: totalTests, TestsPer1Pop: testsPer1Pop, Timestamp: timestamp}
}


