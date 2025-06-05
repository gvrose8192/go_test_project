# go_test_project
Experimental golang project to collect the daily treasury rates from the US government website.

The database initialization and code to populate the database with records since 2015 is complete. Use the init.sql file to initialize your 'treasury_rate' database and then run the program with the "db-init" flag.

The daily update is available via flag "-db-update"

The daily-treasury-rates.sql is now deprecated since the database can be initialized via "-db-init".

The init.sql initializer file is derived from CSV download from the US Treasury.

https://home.treasury.gov/resource-center/data-chart-center/interest-rates/TextView?type=daily_treasury_yield_curve&field_tdr_date_value_month=202505

For CSV format - which is what my daily treasury rate tracking program will want - use this URL:
wget https://home.treasury.gov/resource-center/data-chart-center/interest-rates/daily-treasury-rates.csv/all/202505?type=daily_treasury_yield_curve&field_tdr_date_value_month=202505&page&_format=csv

The URL for year queries is a bit different:
"https://home.treasury.gov/resource-center/data-chart-center/interest-rates/daily-treasury-rates.csv/%s/all?type=daily_treasury_yield_curve&field_tdr_date_value=%s&page&_format=csv"

See the addNewRatesFromBaseTime() function for details.
