/*
 * walk-fast
 *
 * Walks a file system hierarchy using this library.
 */
package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/karrick/godirwalk"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"html"
)

const (
	dirname    = "C:/Users/User/Documents/ClusterFuck"
	optVerbose = true
)

func main() {
	//converter := md.NewConverter("", true, nil)
	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {

			if de.IsDir() {
				log.Println("Working in directory -> ", osPathname)
			} else {
				//NewFilenameAbsolutePath := strings.TrimSuffix(osPathname, ".html") + ".md"
				NewFilenameNameOnly := strings.TrimSuffix(de.Name(), ".html") + ".md"
				//log.Println("Should make new file: ", NewFilenameAbsolutePath)
				// open file
				//file, err := ioutil.ReadFile(osPathname)
				//if err != nil {
				//    panic(err)
				//}

				file, err := os.Open(osPathname)

				// parsing file
				//markdown, err := converter.ConvertBytes(file)
				//if err != nil {
				//   panic(err)
				//}

				// what the fuck
				doc, err := goquery.NewDocumentFromReader(file)
				if err != nil {
					log.Fatal(err)
				}

				h1 := doc.Find("h1").Text()
				h1 = strings.Replace(h1, `"`, "", -1)
				h1 = strings.Replace(h1, "\"", "", -1)
				h1 = strings.Replace(h1, "\\", "", -1)

				siteOrigin, _ := doc.Find("meta[property='og:url']").Attr("content")
				sites, _ := doc.Find("meta[property='og:site_name']").Attr("content")
				datetime, _ := doc.Find("time").Attr("datetime")

				author := doc.Find("a[rel='author']").Text()
				author = strings.Replace(author, `"`, "", -1)
				author = strings.Replace(author, "\"", "", -1)
				author = strings.Replace(author, "\\", "", -1)

				var text []string
				doc.Find("p").Each(func(i int, s *goquery.Selection) {
					text = append(text, s.Text())
				})

				HugoFormatting := fmt.Sprintf(`---
title: "%s"
itemurl: "%s"
sites: "%s"
date: "%s"
author: "%s"
---

%s
`, h1, siteOrigin, html.EscapeString(sites), datetime, html.EscapeString(author), strings.Join(text, "\n"))

				// original_file_name.html -> original_file_name.md
				err = ioutil.WriteFile("C:/Users/User/Documents/SITERELATED/telegram-data-clustering/content/posts/"+NewFilenameNameOnly, []byte(HugoFormatting), 0644)
				if err != nil {
					panic(err)
				}

				//if optVerbose {
				//   fmt.Printf("%s %s\n", de.ModeType(), osPathname)
				//}

				// trying not to die
				//time.Sleep(time.Millisecond * 200)
				file.Close()

			}
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			if optVerbose {
				fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
			}

			// For the purposes of this example, a simple SkipNode will suffice,
			// although in reality perhaps additional logic might be called for.
			return godirwalk.SkipNode
		},
		Unsorted: true, // set true for faster yet non-deterministic enumeration (see godoc)
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
