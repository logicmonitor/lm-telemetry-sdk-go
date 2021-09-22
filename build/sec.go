package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	scanFilePath := flag.String("scanFile", "./result.json", "")
	flag.Parse()
	resultBytes, err := ioutil.ReadFile(*scanFilePath)
	if err != nil {
		log.Fatalf("error in accessing securtiy scan result file :%s", err.Error())
	}

	scanReport := &ReportInfo{}
	err = json.Unmarshal(resultBytes, scanReport)
	if err != nil {
		log.Fatalf("error in parsing security scan result file :%s", err.Error())
	}

	if len(scanReport.Issues) == 0 {
		os.Exit(0)
	}
	log.Fatal("security issues found")
}

type CWEType struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Issue struct {
	Severity   string   `json:"severity"`
	Confidence string   `json:"confidence"`
	CWE        *CWEType `json:"cwe"`
	Rule_ID    string   `json:"rule_id"`
	Details    string   `json:"details"`
	File       string   `json:"file"`
	Code       string   `json:"code"`
	Line       string   `json:"line"`
	Column     string   `json:"column"`
}

type Metrics struct {
	NumFiles int `json:"files"`
	NumLines int `json:"lines"`
	NumNosec int `json:"nosec"`
	NumFound int `json:"found"`
}

type Error struct {
	Line   int    `json:"line"`
	Column int    `json:"column"`
	Err    string `json:"error"`
}

type ReportInfo struct {
	Errors       map[string][]Error `json:"Golang errors"`
	Issues       []*Issue
	Stats        *Metrics
	GosecVersion string
}
