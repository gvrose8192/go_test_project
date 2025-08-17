package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type Rate struct {
	ID          int64
	Date        string
	one_month   float32
	one_5month  float32
	two_month   float32
	three_month float32
	four_month  float32
	six_month   float32
	one_year    float32
	two_year    float32
	three_year  float32
	five_year   float32
	seven_year  float32
	ten_year    float32
	twenty_year float32
	thirty_year float32
}

var db *sql.DB

// Pull all records from the database
func getAllRates() ([]Rate, error) {
	var rates []Rate

	rows, err := db.Query("SELECT * FROM rate")
	if err != nil {
		return nil, fmt.Errorf("getAllRates %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var rate Rate
		if err := rows.Scan(&rate.ID, &rate.Date, &rate.one_month, &rate.one_5month, &rate.two_month, &rate.three_month, &rate.four_month, &rate.six_month, &rate.one_year, &rate.two_year, &rate.three_year, &rate.five_year, &rate.seven_year, &rate.ten_year, &rate.twenty_year, &rate.thirty_year); err != nil {
			return nil, fmt.Errorf("getAllRates %v", err)
		}
		rates = append(rates, rate)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllRates %v", err)
	}
	return rates, nil
}

func rateGetAll(c *gin.Context) {
	var rates []Rate
	rateStrings := []string{}

	// Get everything - queries can pick and choose later
	rates, err := getAllRates()
	if err != nil {
		return
	}

	for i := 0; i < len(rates); i = i + 1 {
		s := fmt.Sprintf("%d %s %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f",  rates[i].ID, rates[i].Date, rates[i].one_month, rates[i].one_5month, rates[i].two_month,  rates[i].three_month, rates[i].four_month, rates[i].six_month, rates[i].one_year, rates[i].two_year, rates[i].three_year, rates[i].five_year, rates[i].seven_year, rates[i].ten_year, rates[i].twenty_year, rates[i].thirty_year)
		rateStrings = append(rateStrings, s)
	}
	c.IndentedJSON(http.StatusOK, rateStrings)
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

func getRateByDate(c *gin.Context) {
	date_string := c.Param("date")
	var rates []Rate

	fmt.Printf("Date : %s\n", date_string)
	fixed_date_string := strings.Replace(date_string, "-", "/", 2)

	rates, err := rateByDate(fixed_date_string)
	if err != nil {
		return
	}

	i := 0
	s := fmt.Sprintf("{%d %s %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f}",  rates[i].ID, rates[i].Date, rates[i].one_month, rates[i].one_5month, rates[i].two_month,  rates[i].three_month, rates[i].four_month, rates[i].six_month, rates[i].one_year, rates[i].two_year, rates[i].three_year, rates[i].five_year, rates[i].seven_year, rates[i].ten_year, rates[i].twenty_year, rates[i].thirty_year)
	c.IndentedJSON(http.StatusOK, s)
}

// rateByID queries for the daily rate with the specified ID.
func rateByID(id int64) (string, error) {
	// A rate to hold data from the returned row.
	var rate Rate

	row := db.QueryRow("SELECT * FROM rate WHERE id = ?", id)
	if err := row.Scan(&rate.ID, &rate.Date, &rate.one_month, &rate.one_5month, &rate.two_month, &rate.three_month, &rate.four_month, &rate.six_month, &rate.one_year, &rate.two_year, &rate.three_year, &rate.five_year, &rate.seven_year, &rate.ten_year, &rate.twenty_year, &rate.thirty_year); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("rateById %d: no such album", id)
		}
		return "", fmt.Errorf("rateById %d: %v", id, err)
	}
	s := fmt.Sprintf("{%d %s %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f}",  rate.ID, rate.Date, rate.one_month, rate.one_5month, rate.two_month,  rate.three_month, rate.four_month, rate.six_month, rate.one_year, rate.two_year, rate.three_year, rate.five_year, rate.seven_year, rate.ten_year, rate.twenty_year, rate.thirty_year)
	return s, nil
}

func getRateByID(c *gin.Context) {
	string_id := c.Param("id")

	// Get the id
	id, err := strconv.ParseInt(string_id, 10, 64)
	if err != nil {
		fmt.Printf("Error converting string to int64\n")
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

func insertNewRate(rateString []string, columns []string) error {
	var err error = nil
	newRate := Rate{ID: 0, Date: rateString[0]}
	//	fmt.Printf("New Record Date is %s\n", newRate.Date)

	for i := 1; i < len(columns); i = i + 1 {
		// Sometimes the government changes horses in the middle of the dang stream
		// See year 2018 for example that this conditional handles
		if len(rateString[i]) < 2 {
			continue
		}
		switch columns[i] {
		case "\"1 Mo\"":
			newRate.one_month, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error one month %v\n", err)
				return err
			}
		case "\"1.5 Month\"":
			newRate.one_5month, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error one_5month %v\n", err)
				return err
			}
		case "\"2 Mo\"":
			newRate.two_month, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error two month %v\n", err)
				return err
			}
		case "\"3 Mo\"":
			newRate.three_month, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error three month%v\n", err)
				return err
			}
		case "\"4 Mo\"":
			newRate.four_month, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error four month %v\n", err)
				return err
			}
		case "\"6 Mo\"":
			newRate.six_month, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error six month %v\n", err)
				return err
			}
		case "\"1 Yr\"":
			newRate.one_year, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error one year %v\n", err)
				return err
			}
		case "\"2 Yr\"":
			newRate.two_year, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error two year %v\n", err)
				return err
			}
		case "\"3 Yr\"":
			newRate.three_year, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error three year %v\n", err)
				return err
			}
		case "\"5 Yr\"":
			newRate.five_year, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error five year %v\n", err)
				return err
			}
		case "\"7 Yr\"":
			newRate.seven_year, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error seven year %v\n", err)
				return err
			}
		case "\"10 Yr\"":
			newRate.ten_year, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error ten year %v\n", err)
				return err
			}
		case "\"20 Yr\"":
			newRate.twenty_year, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error twenty year %v\n", err)
				return err
			}
		case "\"30 Yr\"":
			newRate.thirty_year, err = convertStringToFloat(rateString[i])
			if err != nil {
				fmt.Printf("Conversion error thirty year %v\n", err)
				return err
			}
		}

	}

	//	fmt.Printf("one month: %f\tone 1/2 month: %f\ttwo month: %f\tthree month: %f\t four_month: %f\n", newRate.one_month, newRate.one_5month, newRate.two_month, newRate.three_month, newRate.four_month)
	//	fmt.Printf("six month: %f\tone year: %f\ttwo year: %f\tthree year: %f\tfive year: %f\n", newRate.six_month, newRate.one_year, newRate.two_year, newRate.three_year, newRate.five_year)
	//	fmt.Printf("seven year: %f\tten year: %f\ttwenty year: %f\tthirty year: %f\n", newRate.seven_year, newRate.ten_year, newRate.twenty_year, newRate.thirty_year)

	result, dberr := db.Exec("INSERT INTO rate(date, one_month, one_5month, two_month, three_month, four_month, six_month, one_year, two_year, three_year, five_year, seven_year, ten_year, twenty_year, thirty_year) 					 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", newRate.Date, newRate.one_month, newRate.one_5month, newRate.two_month, newRate.three_month, newRate.four_month, newRate.six_month, newRate.one_year, newRate.two_year, newRate.three_year, newRate.five_year, newRate.seven_year, newRate.ten_year, newRate.twenty_year, newRate.thirty_year)
	if dberr != nil {
		fmt.Printf("insert rate: %v\n", dberr)
		return dberr
	}
	//	fmt.Println("Insert finished")

	rowsAffected, dberr := result.RowsAffected()
	if dberr != nil {
		fmt.Printf("insert rate: %v\n", dberr)
		return dberr
	}
	if rowsAffected != 1 {
		fmt.Println("Rows affected should be one but isn't, instead it is %d\n", rowsAffected)
	}
	//	fmt.Printf("Rows Affected: %d\n", rowsAffected)

	id, dberr := result.LastInsertId()
	if dberr != nil {
		fmt.Printf("get last id: %v\n", dberr)
		return dberr
	}

	fmt.Printf("Created record from date %s with id %d\n", newRate.Date, id)

	return nil
}

func addNewRates(rateArray []string, columns []string) error {
	for i := 0; i < len(rateArray); i = i + 1 {
		csvElements := strings.Split(rateArray[i], ",")
		var rates []Rate
		var rate Rate

		//		fmt.Println(csvElements)
		//		fmt.Printf("Checking for record with Date: %s\n", csvElements[0])

		rows, err := db.Query("SELECT * FROM rate WHERE date like ?", csvElements[0])
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
			insertNewRate(csvElements, columns)
		}
	}

	return nil
}

// Base time will be 2015
func addNewRatesFromBaseTime() {
	var years = [11]string{"2015", "2016", "2017", "2018", "2019", "2020", "2021", "2022", "2023", "2024", "2025"}
	for i := 0; i < len(years); i = i + 1 {
		newLines := []string{}
		s := fmt.Sprintf("https://home.treasury.gov/resource-center/data-chart-center/interest-rates/daily-treasury-rates.csv/%s/all?type=daily_treasury_yield_curve&field_tdr_date_value=%s&page&_format=csv", years[i], years[i])
		//		fmt.Println(s)
		resp, err := http.Get(s)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			//			fmt.Println(line)
			newLines = append(newLines, line)
		}

		if scan_err := scanner.Err(); scan_err != nil {
			fmt.Println("Error reading request body:", scan_err)
			return
		}

		resp.Body.Close()
		// The Treasury delivers the CSV lines from highest to lowest date
		// Reverse them and save the first line, which is the CVS header fields
		// for use in the record insertion, where the header fields change
		// from year to year and sometimes even in the middle of the year.
		csvColumns := strings.Split(newLines[0], ",")
		reversedLines := reverseStringArray(newLines)
		newRateLines := reversedLines[:len(reversedLines)-1]
		//		fmt.Println(csvColumns)
		//		fmt.Println(newRateLines[0])
		addNewRates(newRateLines, csvColumns)
	}
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
	// Reverse them and save the first line, which is the CVS header fields
	// for use in the record insertion, where the header fields change
	// from year to year and sometimes even in the middle of the year.
	csvColumns := strings.Split(newLines[0], ",")
	reversedLines := reverseStringArray(newLines)
	newRateLines := reversedLines[:len(reversedLines)-1]
	addNewRates(newRateLines, csvColumns)
}

func updateRates(c *gin.Context) {
	updateRateTable()
}

func main() {
	// Let's get started
	username := flag.String("username", "your name", "a string")
	password := flag.String("password", "your password", "a string")
	db_update := flag.Bool("db-update", false, "Update database")
	db_init := flag.Bool("db-init", false, "Pull in all records since 2015")

	flag.Parse()

	host_port := os.Getenv("SQL_HOST_PORT")

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = *username
	cfg.Passwd = *password
	cfg.Net = "tcp"
	cfg.Addr = host_port
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

	if *db_init {
		addNewRatesFromBaseTime()
	}

	router := gin.Default()
	// curl http://<ip address>:<port>/rates
	router.GET("/rates", rateGetAll)
	// curl http://<ip address>:<port>/rates/<id>
	router.GET("/rates/:id", getRateByID)
	// curl http://<ip address>:<port>/rateDate/<mm-dd-yyyy>
	router.GET("/rateDate/:date", getRateByDate)
	// curl -X POST http://<ip address>:<port>/updateRates
	router.POST("/updateRates", updateRates)
	router.Run("0.0.0.0:8080")
}
