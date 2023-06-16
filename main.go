package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gocolly/colly"
)

var counter = 0

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("git.ir"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		title := e.Attr("title")
		if title == "دانلود Download" {
			file, err := os.OpenFile("url.txt", os.O_APPEND|os.O_WRONLY, 0664)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			if _, err = file.WriteString(e.Attr("href") + "\n"); err != nil {
				fmt.Println(err)
			}
			counter++
			fmt.Printf("Getting url of video %d \n", counter)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://en.git.ir/udemy-the-ultimate-guide-to-starting-a-drop-servicing-business/")

	cmd := exec.Command("sh ./dl.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
