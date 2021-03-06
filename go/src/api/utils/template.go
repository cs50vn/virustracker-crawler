package utils

var SQL_GET_NEW_TRACKING_ITEM = "select left_name, right_name, total_cases, total_deaths, total_recovered, serious_cases, total_cases_per_1pop, total_deaths_per_1pop, total_tests, tests_per_1pop, timestamp from TRACKING_ITEM where timestamp = (select max(timestamp) from TRACKING_ITEM) order by total_cases desc"
var SQL_INSERT_MAPPING_NAME_ITEM = "insert into MAPPING_NAME(left_name, right_name) values(?, ?)"
var SQL_DELETE_MAPPING_NAME_TABLE = "delete from MAPPING_NAME"


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