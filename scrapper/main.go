package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var BASEURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

type extractedJobs struct {
	id          string
	title       string
	companyName string
	desc        string
	location    string
}

func main() {
	totalPages := getPages()
	var jobs []extractedJobs
	c := make(chan []extractedJobs)

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Process is finshed ", len(jobs))
}

func writeJobs(jobs []extractedJobs) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	utf8bom := []byte{0xEF, 0xBB, 0xBF}
	file.Write(utf8bom)
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Company", "Desc", "Location"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.companyName, job.location, job.desc}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int, mainC chan<- []extractedJobs) {
	var jobs []extractedJobs
	c := make(chan extractedJobs)

	pageUrl := BASEURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting : ", pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkStatusCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	jobCard := doc.Find(".tapItem")

	jobCard.Each(func(i int, s *goquery.Selection) {
		go extractJobs(s, c)
	})

	for i := 0; i < jobCard.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	defer res.Body.Close()

	mainC <- jobs
}

func extractJobs(s *goquery.Selection, c chan<- extractedJobs) {
	id, _ := s.Attr("data-jk")
	title := cleanString(s.Find("h2>span").Text())
	companyName := cleanString(s.Find(".companyName").Text())
	location := cleanString(s.Find(".companyLocation").Text())
	desc := cleanString(s.Find(".job-snippet").Text())
	c <- extractedJobs{
		id:          id,
		location:    location,
		title:       title,
		desc:        desc,
		companyName: companyName,
	}
}

func cleanString(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

func getPages() int {
	pages := 0

	res, err := http.Get(BASEURL)

	checkErr(err)
	checkStatusCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	defer res.Body.Close()
	return pages
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed with status code : ", res.StatusCode)
	}
}
