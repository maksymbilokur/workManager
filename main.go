package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/workmanager/entity"

	"github.com/workmanager/parser"

	_ "github.com/workmanager/parser"
)

//const Limit = 100
const SupportedReports = "members, projects"

func main() {
	var fileNameFlag = flag.String("file", "", "name of the file with data")
	var formatFlag = flag.String("format", "txt", "output format(txt/json)")
	var timeRangeFlag = flag.String("range", "", "time range for retorts. Format: dd:mm:yyyy-dd:mm:yyyy, start -inclusive, end exclusive")

	flag.Parse()

	var timeRange []time.Time
	var err error
	if *timeRangeFlag != "" {
		timeRange, err = getTimeRange(*timeRangeFlag)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		timeRange = append(timeRange, time.Time{}, time.Time{})
	}

	_, repProject, repMember := parser.Reader(*fileNameFlag, parser.FieldNumber, timeRange[0], timeRange[1])

	reportTypes := flag.Args()
	if len(reportTypes) == 0 {
		fmt.Println("no report types in args")
		return
	}

	for _, r := range reportTypes {
		switch r {
		case "members":
			print(repMember, "members", *formatFlag)
		case "projects":
			print(repProject, "projects", *formatFlag)
		default:
			fmt.Println("no such type of report:", r)
			fmt.Println("Supported reports list:", SupportedReports)
		}
	}
}

func print(msg interface{}, reportType string, fmtType string) error {
	fmt.Println("Report type: ", reportType, "=================")
	switch fmtType {
	case "txt":
		fmt.Println(msg)

	/*case "txtln":
	for _, v := range msg {
		fmt.Println(v)
	}
	*/
	case "json":
		b, err := json.MarshalIndent(msg, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	default:
		return errors.New("no such formatting type")
	}

	return nil
}

func getTimeRange(sTimeRange string) ([]time.Time, error) {
	var newRange []time.Time

	sliceTimeRange := strings.Split(sTimeRange, "-")

	if len(sliceTimeRange) == 2 {
		for _, t := range sliceTimeRange {
			tt, err := entity.StringToDate(t)
			if err != nil {
				return newRange, err
			}
			newRange = append(newRange, tt)
		}
		if newRange[0].After(newRange[1]) {
			return newRange, errors.New("end date before start date")
		}
		return newRange, nil
	}

	// add feature to get all month with 1 identifier
	return newRange, errors.New("can't parse date range from flag")
}
