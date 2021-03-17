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

func getProjectStateContentData() *tests.OrderedTests {
	dataSet := &tests.OrderedTests{
		OrderedList: make(tests.OrderedTestList, 0),
		TestDataSet: make(tests.DataSet),
	}

	testCase := "test_notRunning"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input: NotRunning,
		Expected: &StateContent{
			text:  NotRunning,
			badge: "badge-warning",
		},
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_paused"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input: Paused,
		Expected: &StateContent{
			text:  Paused,
			badge: "badge-primary",
		},
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_running"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input: Running,
		Expected: &StateContent{
			text:  Running,
			badge: "badge-success",
		},
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_stopped"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input: Stopped,
		Expected: &StateContent{
			text:  Stopped,
			badge: "badge-danger",
		},
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_unreachable"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input: Unreachable,
		Expected: &StateContent{
			text:  Unreachable,
			badge: "badge-danger",
		},
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_invalid_type"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input: "Invalid",
		Expected: &StateContent{
			text:  "Invalid",
			badge: "badge-secondary",
		},
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	return dataSet
}

func TestGetProjectStateContent(t *testing.T) {
	// Create test data
	dataSet := getProjectStateContentData()

	// Run tests
	for _, testCaseString := range dataSet.OrderedList {
		testCaseString := testCaseString
		t.Run(testCaseString, func(t *testing.T) {
			testCase := dataSet.TestDataSet[testCaseString]
			output := getProjectStateContent(dataSet.TestDataSet[testCaseString].Input.(string))
			tests.CheckResult(testCase.Expected, output, nil, nil, testCaseString, t)
		})
	}
}

func setLocationStringData() *tests.OrderedTests {
	dataSet := &tests.OrderedTests{
		OrderedList: make(tests.OrderedTestList, 0),
		TestDataSet: make(tests.DataSet),
	}

	testCase := "test_bothDefined"
	input := make(map[string]interface{})
	input["city"] = "New York"
	input["country"] = "US"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    input,
		Expected: "New York, US",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_cityOnly"
	input = make(map[string]interface{})
	input["city"] = "New York"
	input["country"] = ""
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    input,
		Expected: "New York",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_countryOnly"
	input = make(map[string]interface{})
	input["city"] = ""
	input["country"] = "US"
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    input,
		Expected: "US",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	testCase = "test_notDefined"
	input = make(map[string]interface{})
	input["city"] = ""
	input["country"] = ""
	dataSet.TestDataSet[testCase] = &tests.Data{
		Input:    input,
		Expected: "Not specified",
	}
	dataSet.OrderedList = append(dataSet.OrderedList, testCase)

	return dataSet
}

func TestSetLocationString(t *testing.T) {
	// Create test data
	dataSet := setLocationStringData()

	// Run tests
	for _, testCaseString := range dataSet.OrderedList {
		testCaseString := testCaseString
		t.Run(testCaseString, func(t *testing.T) {
			testCase := dataSet.TestDataSet[testCaseString]
			input := testCase.Input.(map[string]interface{})
			output := setLocationString(input["country"].(string), input["city"].(string))
			tests.CheckResult(testCase.Expected, output, nil, nil, testCaseString, t)
		})
	}
}
