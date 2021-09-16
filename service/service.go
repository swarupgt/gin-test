package service

import (
	"database/sql"
	"realdummy/logger"
	"realdummy/requestdata"
	"strings"

	_ "github.com/lib/pq"
)

type RequestService interface {
	RespondText(data requestdata.RequestData) requestdata.SentData
	RespondLoc(data requestdata.LocationData) requestdata.SentLocData
}

type reqService struct {
	dat  requestdata.RequestData
	rsDB *sql.DB
}

func New() RequestService {

	logger.Logger("RequestService created")

	connStr := "user=swarupgt dbname=dummydb password=weewoo host=localhost sslmode=disable"
	//database opening
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	//database pinging
	err2 := DB.Ping()
	if err2 != nil {
		panic(err)
	}

	//log
	logger.Logger("Database dummydb connected")

	return &reqService{
		rsDB: DB,
	}
}

func (rs *reqService) RespondText(data requestdata.RequestData) requestdata.SentData {
	logger.Logger("RespondText() called")

	source_no := data.Payload.Source

	//query to get brandname from table, store it into BrandName
	var BrandName string
	exQuery := "SELECT bname FROM brand_data WHERE baseno = '" + source_no + "' ;"
	err3 := rs.rsDB.QueryRow(exQuery).Scan(&BrandName)
	if err3 != nil {
		BrandName = "Outsider"
	}

	//log
	logger.Logger("dummydb query: " + exQuery)

	rs.dat = data

	var sd requestdata.SentData
	sd.Message = Decision(rs.dat.Payload.Payload.Text, BrandName)

	return sd
}

func (rs *reqService) RespondLoc(data requestdata.LocationData) requestdata.SentLocData {
	logger.Logger("RespondLoc() called")

	source_no := data.Payload.Source

	//query to get brandname from table, store it into BrandName
	var BrandName string
	exQuery := "SELECT bname FROM brand_data WHERE baseno = '" + source_no + "' ;"
	err3 := rs.rsDB.QueryRow(exQuery).Scan(&BrandName)
	if err3 != nil {
		BrandName = "Outsider"
	}
	logger.Logger("dummydb query: " + exQuery)
	return NearestLoc(data)
}

//internal functions
func Decision(mess string, bname string) string {
	logger.Logger("Decision() called")

	if strings.EqualFold(mess, "hi") || strings.EqualFold(mess, "hello") {
		return "Welcome to " + bname + ", How can we help you?"
	}

	if strings.Contains(mess, "nearest") || strings.Contains(mess, "closest") || strings.Contains(mess, "near me") {
		return "Please send your location, we will respond with the nearest " + bname + " store's location."
	}

	return "Sorry, we didn't understand :("
}

func NearestLoc(ll requestdata.LocationData) requestdata.SentLocData {

	logger.Logger("NearestLoc() called")
	//brand name can be used to get brand specific lat long calcs
	//can change implementation for nearest location to heart's desire

	var temp requestdata.SentLocData
	temp.Latitude = ll.Payload.Location.Latitude
	temp.Longitude = ll.Payload.Location.Longitude

	return temp
}
