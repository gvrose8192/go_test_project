# go_test_project
Experimental golang project to collect the daily treasury rates from the US government website.

Next steps:
Add a function to run daily, collect the latest CSV format download and add any new records not present in my database.

The daily-treasury-rates.sql initializer file is derived from CSV download from the US Treasury.
The URL below is for May 2025 when this project began.  Modifying the field_tdr_date_value_month
can get a different month's report.

https://home.treasury.gov/resource-center/data-chart-center/interest-rates/TextView?type=daily_treasury_yield_curve&field_tdr_date_value_month=202505

For CSV format - which is what my daily treasury rate tracking program will want - use this URL:
wget https://home.treasury.gov/resource-center/data-chart-center/interest-rates/daily-treasury-rates.csv/all/202505?type=daily_treasury_yield_curve&field_tdr_date_value_month=202505&page&_format=csv
