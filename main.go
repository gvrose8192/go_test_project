package main

import (
	"database/sql"
	"net/http"
	"flag"
	"fmt"
	"log"
	"strconv"

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

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
//func addAlbum(alb Album) (int64, error) {
//	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
//	if err != nil {
//		return 0, fmt.Errorf("addAlbum: %v", err)
//	}
//	id, err := result.LastInsertId()
//	if err != nil {
//		return 0, fmt.Errorf("addAlbum: %v", err)
//	}
//	return id, nil
//}

func main() {

	username := flag.String("username", "your name", "a string")
	password := flag.String("password", "your password", "a string")

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
	fmt.Println("Connected!")

	rates, err := rateByDate("2025-22-05")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Rates found: %v\n", rates)

	// Hard-code ID 3 here to test the query.
	rate, err := rateByID(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Rate found: %v\n", rate)


	router := gin.Default()
	router.GET("/rates", rateGetAll)
	router.GET("/rates/:id", getRateByID)
	router.Run("localhost:8080")

//	alb, err := albumByID(2)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Album found: %v\n", alb)
//
//	albID, err := addAlbum(Album{
//		Title:  "The Modern Sound of Betty Carter",
//		Artist: "Betty Carter",
//		Price:  49.99,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("ID of added album: %v\n", albID)
}
