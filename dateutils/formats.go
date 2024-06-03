package dateutils

import "time"

// ChangeDateFormatDefault function to change date format based on fixed datetime value in golang
func ChangeDateFormatDefault(inputDate, inputDateLayout, format string) (string, error) {
	t, err := time.Parse(inputDateLayout, inputDate)
	if err != nil {
		return inputDate, err
	}

	return t.Format(format), nil
}

// changeDateFormat TODO - parse dates based on a generally accepted format. for example: yyyy-MM-dd
func changeDateFormat() {

}
