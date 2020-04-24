package model

type Item struct {
    LeftName string
    RightName string
    TotalCases int64
    TotalDeaths int64
    TotalRecovered int64
    SeriousCases int64
    TotalCasesPer1Pop float64
    TotalDeathsPer1Pop float64
    TotalTests int64
    TestsPer1Pop float64
    Timestamp int64
}

func MakeItem(leftName string, rightName string, totalCases int64, totalDeaths int64, totalRecovered int64, seriousCases int64, totalCasesPer1Pop float64, totalDeathsPer1Pop float64, totalTests int64, testsPer1Pop float64, timestamp int64) *Item {
    return &Item{LeftName: leftName, RightName: rightName, TotalCases: totalCases, TotalDeaths: totalDeaths, TotalRecovered: totalRecovered, SeriousCases: seriousCases, TotalCasesPer1Pop: totalCasesPer1Pop, TotalDeathsPer1Pop: totalDeathsPer1Pop, TotalTests: totalTests, TestsPer1Pop: testsPer1Pop, Timestamp: timestamp}
}


