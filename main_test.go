package main

import (
	"csvparser/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCaseType struct {
	resultFile  string
	sourceDir   string
	resultArray []entities.Place
	maxResult   int
	idLimit     int
}

var (
	testCases = []TestCaseType{
		{
			resultArray: []entities.Place{{Id: 1100002, Name: "Hotel 1", Condition: "Standart Room+10 nights+AI", State: "quote", Price: "10000RUB", PriceInt: 10000}, {Id: 1100002, Name: "Hotel 1", Condition: "Standart Room+10 nights+AI", State: "quote", Price: "10000RUB", PriceInt: 10000}, {Id: 1100002, Name: "Hotel 1", Condition: "Standart Room+10 nights+AI", State: "quote", Price: "10000RUB", PriceInt: 10000}, {Id: 1100002, Name: "Hotel 1", Condition: "Standart Room+10 nights+AI", State: "quote", Price: "10000RUB", PriceInt: 10000}, {Id: 1110032, Name: "Hotel 2", Condition: "Standart Room+7 nights+BB", State: "quote", Price: "120000RUB", PriceInt: 120000}, {Id: 1110032, Name: "Hotel 2", Condition: "Standart Room+7 nights+BB", State: "quote", Price: "120000RUB", PriceInt: 120000}, {Id: 1110032, Name: "Hotel 2", Condition: "Standart Room+7 nights+BB", State: "quote", Price: "120000RUB", PriceInt: 120000}, {Id: 1110032, Name: "Hotel 2", Condition: "Standart Room+7 nights+BB", State: "quote", Price: "120000RUB", PriceInt: 120000}, {Id: 1110032, Name: "Hotel 2", Condition: "King Room+7 nights+BB", State: "request", Price: "150000RUB", PriceInt: 150000}, {Id: 1110032, Name: "Hotel 2", Condition: "King Room+7 nights+BB", State: "request", Price: "150000RUB", PriceInt: 150000}, {Id: 1110032, Name: "Hotel 2", Condition: "King Room+7 nights+BB", State: "request", Price: "150000RUB", PriceInt: 150000}},
			maxResult:   10,
			idLimit:     3,
			resultFile:  "./result_test.csv",
			sourceDir:   "./test_data",
		},
		{
			resultArray: []entities.Place{{Id: 1100002, Name: "Hotel 1", Condition: "Standart Room+10 nights+AI", State: "quote", Price: "10000RUB", PriceInt: 10000}, {Id: 1100002, Name: "Hotel 1", Condition: "Standart Room+10 nights+AI", State: "quote", Price: "10000RUB", PriceInt: 10000}, {Id: 1110032, Name: "Hotel 2", Condition: "Standart Room+7 nights+BB", State: "quote", Price: "120000RUB", PriceInt: 120000}, {Id: 1110032, Name: "Hotel 2", Condition: "Standart Room+7 nights+BB", State: "quote", Price: "120000RUB", PriceInt: 120000}, {Id: 1110032, Name: "Hotel 2", Condition: "King Room+7 nights+BB", State: "request", Price: "150000RUB", PriceInt: 150000}, {Id: 1110032, Name: "Hotel 2", Condition: "King Room+7 nights+BB", State: "request", Price: "150000RUB", PriceInt: 150000}},
			maxResult:   5,
			idLimit:     3,
			resultFile:  "./result_test.csv",
			sourceDir:   "./test_data",
		},
	}
)

func TestParse(t *testing.T) {
	for _, testCase := range testCases {
		resultArray := parse(testCase.maxResult, testCase.sourceDir)
		assert.Equal(t, resultArray, testCase.resultArray)
	}
}

func TestFileExist(t *testing.T) {
	for _, testCase := range testCases {
		resultArray := parse(testCase.maxResult, testCase.sourceDir)
		writeResult(resultArray, testCase.idLimit, testCase.resultFile)
		assert.FileExists(t, testCase.resultFile)
	}
}

func TestMaxResult(t *testing.T) {
	for _, testCase := range testCases {
		resultArray := parse(testCase.maxResult, testCase.sourceDir)
		assert.LessOrEqual(t, len(resultArray), testCase.maxResult)
	}
}
