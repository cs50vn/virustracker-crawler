package utils

var SQL_GET_ALL_MAPPING_NAME = "select left_name, right_name from MAPPING_NAME"
var SQL_INSERT_TRACKING_ITEM = "insert into TRACKING_ITEM(left_name, right_name, total_cases, total_deaths, total_recovered, serious_cases, total_cases_per_1pop, total_deaths_per_1pop, total_tests, tests_per_1pop, timestamp) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

var RESULT_TEMPLATE = `{
    "statusCode": %d,
    "data": %s }`
var ARRAY_TEMPLATE = `[%s]`
var OBJECT_TEMPLATE = `{%s}`
var ITEM_TEMPLATE = `"%s":"%s"`
var ITEM2_TEMPLATE = `"%s":%d`
var ITEM3_TEMPLATE = `"%s":%f`
var ITEM4_TEMPLATE = `"%s": [%s]`
var ITEM5_TEMPLATE = `"%s": {%s}`
var ITEM6_TEMPLATE = `"%s": %s`
var ITEM7_TEMPLATE = `"%s"`

//===========================================
var FORCE_UPDATE = "force_update"
var RECOMMEND = "recommend"
var NONE = "none"