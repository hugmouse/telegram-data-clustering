/*
 * There is a lot of panic in this code. It was intentional.
 */
package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/karrick/godirwalk"
    "html"
    "io/ioutil"
    "log"
    "os"
    "strings"
)

const (
    dirname            = ""
    dirnameDestination = ""
    optVerbose         = true
)

func main() {
    err := godirwalk.Walk(dirname, &godirwalk.Options{
        Callback: func(osPathname string, de *godirwalk.Dirent) error {

            if de.IsDir() {
                log.Println("Working in directory -> ", osPathname)
            } else {

                // xyz.html -> xyz.md
                NewFilenameNameOnly := strings.TrimSuffix(de.Name(), ".html") + ".md"

                file, err := os.Open(osPathname)

                doc, err := goquery.NewDocumentFromReader(file)
                if err != nil {
                    log.Fatal(err)
                }

                // Finding title
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

                // originalPath/original_file_name.html -> destPath/original_file_name.md
                err = ioutil.WriteFile(dirnameDestination+"git-"+NewFilenameNameOnly, []byte(HugoFormatting), 0644)
                if err != nil {
                    panic(err)
                }

                err = file.Close()
                if err != nil {
                    panic(err)
                }

            }
            return nil
        },
        ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
            if optVerbose {
                _, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
            }

            // For the purposes of this example, a simple SkipNode will suffice,
            // although in reality perhaps additional logic might be called for.
            return godirwalk.SkipNode
        },
        Unsorted: true, // set true for faster yet non-deterministic enumeration (see godoc)
    })
    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
        os.Exit(1)
    }
}
