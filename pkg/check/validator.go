package check

import (
	"errors"
	"regexp"
	"rent-car/config"
	"time"
)

func ValidateCarYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}
	return nil
}

func ValidateGmailCustomer(e string) bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,3}$`)
    return emailRegex.MatchString(e)
}


func ValidatePhoneNumberOfCustomer(phone string) bool {
	if 12 < len(phone) && len(phone) <= 13{
		phoneregex:= regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
		return phoneregex.MatchString(phone)
	}
	return false
}

func ValidatePassword(password string) error {
  lowercaseRegex := `[a-z]`
  hasLowercase, _ := regexp.MatchString(lowercaseRegex, password)
  uppercaseRegex := `[A-Z]`
  hasUppercase, _ := regexp.MatchString(uppercaseRegex, password)
  digitRegex := `[0-9]`
  hasDigit, _ := regexp.MatchString(digitRegex, password)
  symbolRegex := `[!@#$%^&*()-_+=~\[\]{}|\\:;"'<>,.?\/]`
  hasSymbol, _ := regexp.MatchString(symbolRegex, password)

  if hasLowercase && hasUppercase && hasDigit && hasSymbol && len(password) >= 8 {
    return nil
  }

  return errors.New("password does not meet the criteria")
  }


  func ValidatingOrderStatusForAuth(status string) error {
		for _, s := range config.ORDER_STATUS {
			if s == status {
				return nil 
			}
			
		}
		return errors.New("error Valid order status")
	}


  func ValidateDateOfFormatForOrder(dateStr string) error {
    datePattern := `^\d{4}-\d{2}-\d{2}$`
    dateRegex := regexp.MustCompile(datePattern)
    if !dateRegex.MatchString(dateStr) {
      return errors.New("invalid date format")
    }
    return nil
  }