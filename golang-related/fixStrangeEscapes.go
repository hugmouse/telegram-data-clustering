/*
 * Fix Strange Escapes
 *
 * And not only escaping "\"
 */
package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "regexp"

    "github.com/karrick/godirwalk"
)

func main() {
    optVerbose := flag.Bool("verbose", false, "Print file system entries.")
    flag.Parse()

    dirname := "."
    if flag.NArg() > 0 {
        dirname = flag.Arg(0)
    }

    err := godirwalk.Walk(dirname, &godirwalk.Options{
        Callback: func(osPathname string, de *godirwalk.Dirent) error {
            file, err := os.Open(osPathname)

            matched, err := regexp.MatchString("\\", "Hello World")
            fmt.Println("Matched:", matched, "Error:", err)

            err = ioutil.WriteFile("C:/Users/User/Documents/SITERELATED/telegram-data-clustering/content/posts/" + de.Name(), []byte(HugoFormatting), 0644)
            if err != nil {
                panic(err)
            }
            return nil
        },
        ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {

            log.Printf("ERROR: %s\n", err)

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