DROP TABLE IF EXISTS rate;
CREATE TABLE rate (
	id		INT AUTO_INCREMENT NOT NULL,
	date		VARCHAR(10) NOT NULL,
	one_month	DECIMAL(5,2) NOT NULL,
	one_5month	DECIMAL(5,2) NOT NULL,
	two_month	DECIMAL(5,2) NOT NULL,
	three_month	DECIMAL(5,2) NOT NULL,
	four_month	DECIMAL(5,2) NOT NULL,
	six_month	DECIMAL(5,2) NOT NULL,
	one_year	DECIMAL(5,2) NOT NULL,
	two_year	DECIMAL(5,2) NOT NULL,
	three_year	DECIMAL(5,2) NOT NULL,
	five_year	DECIMAL(5,2) NOT NULL,
	seven_year	DECIMAL(5,2) NOT NULL,
	ten_year	DECIMAL(5,2) NOT NULL,
	twenty_year	DECIMAL(5,2) NOT NULL,
	thirty_year	DECIMAL(5,2) NOT NULL,
	PRIMARY KEY 	(`id`)
);
