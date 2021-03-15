package contents

import (
	"encoding/json"
	"fmt"
)

const (
	CheckBoxUnChecked = "unchecked"
	CheckBoxChecked   = "checked"
)

func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
		return
	}
	fmt.Println("Failed to pretty print data")
}

type StateContent struct {
	text  string
	badge string
}

func convertToCheckboxValue(input string) string {
	if input == CheckBoxUnChecked {
		return ""
	}
	return input
}

func convertCheckedToYesNo(input string) string {
	if input == CheckBoxUnChecked {
		return "No"
	}
	return "Yes"
}

func generatePriceString(paymentType string, amount string) string {
	if paymentType == PaymentTypeFree {
		return paymentType
	} else if paymentType == PaymentTypeSub {
		return fmt.Sprintf("%s NZD/month", amount)
	} else {
		return fmt.Sprintf("%s NZD", amount)
	}
}

// GetProjectStateContent returns UI color of the project state based on the state value.
func getProjectStateContent(stateString string) *StateContent {
	state := &StateContent{
		text: stateString,
	}
	switch stateString {
	case NotRunning:
		state.badge = "badge-warning" // orange
	case Paused:
		state.badge = "badge-primary" // lightblue
	case Running:
		state.badge = "badge-success" // green
	case Stopped:
		state.badge = "badge-danger" // red
	case Unreachable:
		state.badge = "badge-danger" // red
	default:
		state.badge = "badge-secondary" // lightgray
	}
	return state
}

func setLocationString(country string, city string) string {
	if country == "" && city == "" {
		return "Not specified"
	} else if country != "" && city == "" {
		return country
	} else if country == "" && city != "" {
		return city
	}
	return fmt.Sprintf("%s, %s", city, country)
}
