package leetcode

import (
	"fmt"
	"strings"
)

func ReformatDate(date string) string {
	sdate := strings.Split(date, " ")
	fdate := make([]string, 3)
	f := NewFormator()
	fdate[0] = f.Date(sdate[0])
	fdate[1] = f.Month(sdate[1])
	fdate[2] = sdate[2]

	return strings.Join(fdate, "-")
}

type Formater struct {
	DataMap  map[string]string
	MonthMap map[string]string
}

func NewFormator() Formater {
	dm := map[string]string{
		"1st":  "01",
		"2nd":  "02",
		"3rd":  "03",
		"21st": "21",
		"22nd": "22",
		"23rd": "23",
		"31st": "31",
	}
	for i := 4; i < 21; i++ {
		dm[fmt.Sprintf("%dth", i)] = fmt.Sprintf("%02d", i)
	}
	for i := 24; i < 31; i++ {
		dm[fmt.Sprintf("%dth", i)] = fmt.Sprintf("%02d", i)
	}

	mm := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	return Formater{dm, mm}
}

func (f Formater) Date(s string) string {
	return f.DataMap[s]
}

func (f Formater) Month(s string) string {
	return f.MonthMap[s]
}
