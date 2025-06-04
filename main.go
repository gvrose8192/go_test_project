package main

import (
	"database/sql"
	"net/http"
	"flag"
	"fmt"
	"log"
	"strconv"
	"bufio"
	"time"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

type Rate struct {
	ID		int64
	Date		string
	one_month	float32
	one_5month	float32
	two_month	float32
	three_month	float32
	four_month	float32
	six_month	float32
	one_year	float32
	two_year	float32
	three_year	float32
	five_year	float32
	seven_year	float32
	ten_year	float32
	twenty_year	float32
	thirty_year	float32
}

// Pull all records from the database
func getAllRates() ([]Rate, error) {
	var rates []Rate

	rows, err := db.Query("SELECT * FROM rate")
	if err != nil {
		return nil, fmt.Errorf("rateByDate %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var rate Rate
		if err := rows.Scan(&rate.ID, &rate.Date, &rate.one_month, &rate.one_5month, &rate.two_month, &rate.three_month, &rate.four_month, &rate.six_month, &rate.one_year, &rate.two_year, &rate.three_year, &rate.five_year, &rate.seven_year, &rate.ten_year, &rate.twenty_year, &rate.thirty_year); err != nil {
			return nil, fmt.Errorf("rateByDate %v", err)
		}
		rates = append(rates, rate)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rateByDate %v", err)
	}
	return rates, nil
}

func rateGetAll(c *gin.Context) {
	var rates []Rate

	// Get everything - queries can pick and choose later
	rates, err := getAllRates()
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, rates)
}

func rateByDate(date string) ([]Rate, error) {
	var rates []Rate

	rows, err := db.Query("SELECT * FROM rate WHERE date like ?", date)
	if err != nil {
		return nil, fmt.Errorf("rateByDate %q: %v", date, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var rate Rate
		if err := rows.Scan(&rate.ID, &rate.Date, &rate.one_month, &rate.one_5month, &rate.two_month, &rate.three_month, &rate.four_month, &rate.six_month, &rate.one_year, &rate.two_year, &rate.three_year, &rate.five_year, &rate.seven_year, &rate.ten_year, &rate.twenty_year, &rate.thirty_year); err != nil {
			return nil, fmt.Errorf("rateByDate %q: %v", date, err)
		}
		rates = append(rates, rate)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rateByDate %q: %v", date, err)
	}
	return rates, nil
}

// rateByID queries for the daily rate with the specified ID.
func rateByID(id int64) (Rate, error) {
	// A rate to hold data from the returned row.
	var rate Rate

	row := db.QueryRow("SELECT * FROM rate WHERE id = ?", id)
	if err := row.Scan(&rate.ID, &rate.Date, &rate.one_month, &rate.one_5month, &rate.two_month, &rate.three_month, &rate.four_month, &rate.six_month, &rate.one_year, &rate.two_year, &rate.three_year, &rate.five_year, &rate.seven_year, &rate.ten_year, &rate.twenty_year, &rate.thirty_year); err != nil {
		if err == sql.ErrNoRows {
			return rate, fmt.Errorf("rateById %d: no such album", id)
		}
		return rate, fmt.Errorf("rateById %d: %v", id, err)
	}
	return rate, nil
}

func getRateByID(c *gin.Context) {
	string_id := c.Param("id")
	var rate Rate

	// Get the id
	id, err := strconv.ParseInt(string_id, 10, 64)
	if err != nil {
		fmt.Println("Error converting string to int64:", err)
		return
	}

	rate, err2 := rateByID(id)
	if err2 != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, rate)
}

func convertStringToFloat(number_string string) (float32, error) {
	retVal64, err := strconv.ParseFloat(number_string, 64)
	if err != nil {
		return 0.0, err
	}
	retVal32 := float32(retVal64)
	return retVal32, nil
}

func reverseStringArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func insertNewRate(rateString []string) error {
	newRate := Rate{ID: 0, Date: rateString[0]}
	fmt.Printf("New Record Date is %s\n", newRate.Date)

	var err error = nil
	newRate.one_month, err = convertStringToFloat(rateString[1])
	if err != nil {
		return err
	}

	newRate.one_5month, err = convertStringToFloat(rateString[2])
	if err != nil {
		return err
	}

	newRate.two_month, err = convertStringToFloat(rateString[3])
	if err != nil {
		return err
	}

	newRate.three_month, err = convertStringToFloat(rateString[4])
	if err != nil {
		return err
	}

	newRate.four_month, err = convertStringToFloat(rateString[5])
	if err != nil {
		return err
	}

	newRate.six_month, err = convertStringToFloat(rateString[6])
	if err != nil {
		return err
	}

	newRate.one_year, err = convertStringToFloat(rateString[7])
	if err != nil {
		return err
	}

	newRate.two_year, err = convertStringToFloat(rateString[8])
	if err != nil {
		return err
	}

	newRate.three_year, err = convertStringToFloat(rateString[9])
	if err != nil {
		return err
	}

	newRate.five_year, err = convertStringToFloat(rateString[10])
	if err != nil {
		return err
	}

	newRate.seven_year, err = convertStringToFloat(rateString[11])
	if err != nil {
		return err
	}

	newRate.ten_year, err = convertStringToFloat(rateString[12])
	if err != nil {
		return err
	}

	newRate.twenty_year, err = convertStringToFloat(rateString[13])
	if err != nil {
		return err
	}

	newRate.thirty_year, err = convertStringToFloat(rateString[14])
	if err != nil {
		return err
	}


	fmt.Printf("one month: %f\tone 1/2 month: %f\ttwo month: %f\tthree month: %f\t four_month: %f\n", newRate.one_month, newRate.one_5month, newRate.two_month, newRate.three_month, newRate.four_month)
	fmt.Printf("six month: %f\tone year: %f\ttwo year: %f\tthree year: %f\tfive year: %f\n", newRate.six_month, newRate.one_year, newRate.two_year, newRate.three_year, newRate.five_year)
	fmt.Printf("seven year: %f\tten year: %f\ttwenty year: %f\tthirty year: %f\n", newRate.seven_year, newRate.ten_year, newRate.twenty_year, newRate.thirty_year)

	result, dberr := db.Exec("INSERT INTO rate(date, one_month, one_5month, two_month, three_month, four_month, six_month, one_year, two_year, three_year, five_year, seven_year, ten_year, twenty_year, thirty_year) 					 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", newRate.Date, newRate.one_month, newRate.one_5month, newRate.two_month, newRate.three_month, newRate.four_month, newRate.six_month, newRate.one_year, newRate.two_year, newRate.three_year, newRate.five_year, newRate.seven_year, newRate.ten_year, newRate.twenty_year, newRate.thirty_year)
	if dberr != nil {
		fmt.Printf("insert rate: %v\n", dberr)
		return dberr
	}
	fmt.Println("Insert finished")

	rowsAffected, dberr := result.RowsAffected()
	if dberr != nil {
		fmt.Printf("insert rate: %v\n", dberr)
		return dberr
	}
	fmt.Printf("Rows Affected: %d\n", rowsAffected)

	id, dberr := result.LastInsertId()
	if dberr != nil {
		fmt.Printf("get last id: %v\n", dberr)
		return dberr
	}

	fmt.Printf("Created record with id %d\n", id)

	return nil
}

func addNewRates(rateArray []string) error {
	for i := 0; i < len(rateArray); i = i + 1 {
		csvElements := strings.Split(rateArray[i], ",")
		var rates []Rate
		var rate Rate

		fmt.Println(csvElements)
		fmt.Printf("Checking for record with Date: %s\n", csvElements[0])

		rows, err := db.Query("SELECT * FROM rate WHERE date like ?",  csvElements[0])
		if err != nil {
		 	return err
		}
		for rows.Next() {
			if err := rows.Scan(&rate.ID, &rate.Date, &rate.one_month, &rate.one_5month, &rate.two_month, &rate.three_month, &rate.four_month, &rate.six_month, &rate.one_year, &rate.two_year, &rate.three_year, &rate.five_year, &rate.seven_year, &rate.ten_year, &rate.twenty_year, &rate.thirty_year); err != nil {
				return err
			}
			rates = append(rates, rate)
		}
		rows.Close()
		if len(rates) == 0 {
			insertNewRate(csvElements)
		}
	}

	return nil
}

func updateRateTable() {
	newLines := []string{}
	currentTime := time.Now()
	f := currentTime.Format("200601")

	s := fmt.Sprintf("https://home.treasury.gov/resource-center/data-chart-center/interest-rates/daily-treasury-rates.csv/all/%s?type=daily_treasury_yield_curve&field_tdr_date_value_month=%s&page&_format=csv", f, f)

	resp, err := http.Get(s)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		newLines = append(newLines, line)
	}

	if scan_err := scanner.Err(); scan_err != nil {
		fmt.Println("Error reading request body:", scan_err)
		return
	}


	// The Treasury delivers the CSV lines from highest to lowest date
	// Reverse that and also drop the first line which is the CVS header fields
	// Our database 'rate' table already defines those.
	reversedLines := reverseStringArray(newLines)
	newRateLines := reversedLines[:len(reversedLines) - 1]
	addNewRates(newRateLines)
}

func main() {
	// Let's get started
	username := flag.String("username", "your name", "a string")
	password := flag.String("password", "your password", "a string")
	db_update := flag.Bool("db-update", false, "Update database")

	flag.Parse()

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = *username
	cfg.Passwd = *password
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "treasury_rate"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	if *db_update {
		updateRateTable()
	}

	router := gin.Default()
	router.GET("/rates", rateGetAll)
	router.GET("/rates/:id", getRateByID)
	router.Run("localhost:8080")
}
