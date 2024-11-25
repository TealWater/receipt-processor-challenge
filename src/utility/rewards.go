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

func ValidateName(name string) (*int, error) {
	isValid, err := regexp.MatchString("^[\\w\\s\\-&]+$", name)
	if err != nil || !isValid {
		return nil, errors.New("Retailer name on the receipt is invalid. Please submit a name that contains only alpha-numeric characters.\nex: \"Buy Me 50 Coffees-Co\"")
	}

	points := 0
	for _, v := range name {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			points++
		}
	}
	return &points, nil
}

func ValidateTotal(total float64) int {
	if math.Trunc(total) == total {
		return 50
	}

	if math.Mod(total, 0.25) == 0 {
		return 25
	}
	return 0
}

func CountItems(items []model.Item) int {
	return (len(items) / 2) * 5
}

func ValidateItemDescription(items []model.Item) int {
	points := 0
	for _, v := range items {
		trimLen := len(strings.TrimSpace(v.ShortDescription))
		if math.Mod(float64(trimLen), 3) == 0 {
			val, _ := strconv.ParseFloat(v.Price, 64)
			val = val * 0.2

			//if val is xx.00 -> its already at nearest int. if xx.01+ -> neareast int is up 1
			if math.Trunc(val) == val {
				points += int(val)
			} else {
				points += int(val) + 1
			}
		}
	}
	return points
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

func ValidatePurchaseTime(time string) int {
	parsed := strings.Split(time, ":")
	hour, _ := strconv.Atoi(parsed[0])
	min, _ := strconv.Atoi(parsed[1])

	if hour == 14 && min >= 1 || hour > 14 && hour < 16 {
		return 10
	}
	return 0
}
