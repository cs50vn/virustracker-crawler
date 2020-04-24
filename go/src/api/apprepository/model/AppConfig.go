package model

type AppConfig struct {
    Adminkey  string `json: "adminkey"`
    Clientkey string `json: "clientkey"`
    Dbname    string `json: "dbname"`
    Port      string `json: "port"`
    CrawUrl   string `json: "crawurl"`
    HookUrl   string `json: "hookurl"`
    Times     []string `json: "times"`
}
