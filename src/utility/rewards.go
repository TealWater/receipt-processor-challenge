package utility

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/TealWater/fetch-rewards/model"
)

func ValidateName(name string) (int, error) {
	isValid, err := regexp.MatchString("^[\\w\\s\\-&]+$", name)
	if err != nil || !isValid {
		return 0, errors.New("Retailer name on the receipt is invalid. Please submit a name that contains only alpha-numeric characters.\nex: \"Buy Me 50 Coffees-Co\"")
	}

	points := 0
	for _, v := range name {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			points++
		}
	}
	return points, nil
}

func ValidateTotal(amnt string) (int, error) {
	subTotal, err := strconv.ParseFloat(amnt, 64)
	total := 0
	if err != nil {
		return 0, errors.New("The total provided on the receipt is malformed")
	}

	if math.Trunc(subTotal) == subTotal {
		total += 50
	}

	if math.Mod(subTotal, 0.25) == 0 {
		total += 25
	}
	return total, nil
}

func CountItems(items []model.Item) int {
	return (len(items) / 2) * 5
}

func ValidateItemDescription(items []model.Item) (int, error) {
	points := 0
	for _, v := range items {
		if len(v.ShortDescription) < 1 {
			return 0, errors.New("Invalid Item Description provided")
		}

		if len(v.Price) < 1 {
			return 0, errors.New("Price tag is malformed")
		}

		trimLen := len(strings.TrimSpace(v.ShortDescription))
		if math.Mod(float64(trimLen), 3) == 0 {
			val, err := strconv.ParseFloat(v.Price, 64)
			if err != nil {
				return 0, errors.New("Invalid Price provided")
			}
			val = val * 0.2

			//if val is xx.00 -> its already at nearest int. if xx.01+ -> neareast int is up 1
			if math.Trunc(val) == val {
				points += int(val)
			} else {
				points += int(val) + 1
			}
		}
	}
	return points, nil
}

func ValidatePurchaseDate(date string) (int, error) {
	times, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, errors.New("Invalid date provided.\neg of valid date: yy-mm-day")
	}

	if times.Day()%2 != 0 {
		return 6, nil
	}
	return 0, nil
}

func ValidatePurchaseTime(time string) (int, error) {
	parsed := strings.Split(time, ":")
	if len(parsed) != 2 {
		return 0, errors.New("Invalid Time provided.")
	}

	hour, err := strconv.Atoi(parsed[0])
	if err != nil {
		return 0, errors.New("Invalid Time provided.")
	}

	min, err := strconv.Atoi(parsed[1])
	if err != nil {
		return 0, errors.New("Invalid Time provided.")
	}

	if hour == 14 && min >= 1 || hour > 14 && hour < 16 {
		return 10, nil
	}
	return 0, nil
}
