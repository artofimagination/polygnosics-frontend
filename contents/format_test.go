package contents

import (
	"polygnosics-frontend/tests"
	"testing"
)

func convertToCheckboxValueData() *tests.OrderedTests {
	dataSet := &tests.OrderedTests{
		OrderedList: make(tests.OrderedTestList, 0),
		TestDataSet: make(tests.DataSet),
	}

	testCase := "test_unchecked"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    "unchecked",
		Expected: "",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_checked"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    "checked",
		Expected: "checked",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_invalid"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    "test",
		Expected: "",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	return dataSet
}

func TestConvertToCheckboxValue(t *testing.T) {
	// Create test data
	dataSet := convertToCheckboxValueData()

	// Run tests
	for _, testCaseString := range dataSet.OrderedList {
		testCaseString := testCaseString
		t.Run(testCaseString, func(t *testing.T) {
			testCase := dataSet.TestDataSet[testCaseString]
			output := convertToCheckboxValue(testCase.Input.(string))
			tests.CheckResult(testCase.Expected, output, nil, nil, testCaseString, t)
		})
	}
}

func convertCheckedToYesNoData() *tests.OrderedTests {
	dataSet := &tests.OrderedTests{
		OrderedList: make(tests.OrderedTestList, 0),
		TestDataSet: make(tests.DataSet),
	}

	testCase := "test_unchecked"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    "unchecked",
		Expected: "No",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_checked"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    "checked",
		Expected: "Yes",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_invalid"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    "test",
		Expected: "No",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	return dataSet
}

func TestConvertCheckedToYesNo(t *testing.T) {
	// Create test data
	dataSet := convertCheckedToYesNoData()

	// Run tests
	for _, testCaseString := range dataSet.OrderedList {
		testCaseString := testCaseString
		t.Run(testCaseString, func(t *testing.T) {
			testCase := dataSet.TestDataSet[testCaseString]
			output := convertCheckedToYesNo(testCase.Input.(string))
			tests.CheckResult(testCase.Expected, output, nil, nil, testCaseString, t)
		})
	}
}

func generatePriceStringData() *tests.OrderedTests {
	dataSet := &tests.OrderedTests{
		OrderedList: make(tests.OrderedTestList, 0),
		TestDataSet: make(tests.DataSet),
	}

	testCase := "test_free"
	input := make(map[string]interface{})
	input["type"] = PaymentTypeFree
	input["price"] = ""
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    input,
		Expected: PaymentTypeFree,
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_singlePayment"
	input = make(map[string]interface{})
	input["type"] = PaymentTypeSingle
	input["price"] = "5"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    input,
		Expected: "5 NZD",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_subscription"
	input = make(map[string]interface{})
	input["type"] = PaymentTypeSub
	input["price"] = "5"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    input,
		Expected: "5 NZD/month",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_invalid_type"
	input = make(map[string]interface{})
	input["type"] = "TestType"
	input["price"] = ""
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    input,
		Expected: InvalidPaymentType,
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	return dataSet
}

func TestGeneratePriceString(t *testing.T) {
	// Create test data
	dataSet := generatePriceStringData()

	// Run tests
	for _, testCaseString := range dataSet.OrderedList {
		testCaseString := testCaseString
		t.Run(testCaseString, func(t *testing.T) {
			testCase := dataSet.TestDataSet[testCaseString]
			input := testCase.Input.(map[string]interface{})
			output := generatePriceString(input["type"].(string), input["price"].(string))
			tests.CheckResult(testCase.Expected, output, nil, nil, testCaseString, t)
		})
	}
}
