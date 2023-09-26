CREATE TABLE tblBUSDPrice (
	id SERIAL PRIMARY KEY,
	time timestamp DEFAULT current_timestamp,
	price text
);

CREATE TABLE tblBUSD1MinPercent (
	id SERIAL PRIMARY KEY,
	time timestamp DEFAULT current_timestamp,
    prevTime timestamp,
	symbol text,
	price	real,
	prevPrice real,
	percent real
);

CREATE TABLE tblBUSD5MinPercent (
	id SERIAL PRIMARY KEY,
	time timestamp DEFAULT current_timestamp,
    prevTime timestamp,
	symbol text,
	price	real,
	prevPrice real,
	percent real
);

CREATE TABLE tblBUSD10MinPercent (
	id SERIAL PRIMARY KEY,
	time timestamp DEFAULT current_timestamp,
    prevTime timestamp,
	symbol text,
	price	real,
	prevPrice real,
	percent real
);

CREATE TABLE tblBUSD15MinPercent (
	id SERIAL PRIMARY KEY,
	time timestamp DEFAULT current_timestamp,
    prevTime timestamp,
	symbol text,
	price	real,
	prevPrice real,
	percent real
);

CREATE TABLE tblBUSD30MinPercent (
	id SERIAL PRIMARY KEY,
	time timestamp DEFAULT current_timestamp,
    prevTime timestamp,
	symbol text,
	price	real,
	prevPrice real,
	percent real
);

CREATE TABLE tblBUSD60MinPercent (
	id SERIAL PRIMARY KEY,
	time timestamp DEFAULT current_timestamp,
    prevTime timestamp,
	symbol text,
	price	real,
	prevPrice real,
	percent real
);

CREATE TABLE tblBUSDTopCoinHistory (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    topCoin text
);

CREATE TABLE tblTopCoinCounter (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    topCoinCounter text
)
