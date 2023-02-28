package main

import(
"github.com/dineshsprabu/concurrent-web-crawler"
"os"
alert  "github.com/gookit/color"
"github.com/martinlindhe/notify"
"fmt"
"bufio"
"bytes"
"strconv"
"flag"
"github.com/raphamorim/go-rainbow"
"github.com/dimiro1/banner"
"github.com/mattn/go-colorable"
"github.com/logrusorgru/aurora"
)

func usage() {

	
    fmt.Fprintf(os.Stderr, "usage: ./webcrawler [file-url-list] [max-req] [delay-req] -save [SAVE-PATH-FOLDER]\n")
    flag.PrintDefaults()
    os.Exit(2)
}

func main(){
	// Creating a web crawler object with configurations.


	flag.Usage = usage
    	flag.Parse()

    	args := flag.Args()
    	
	if len(args) < 4 {

		who := "ELMO9AWIMDZ"

		notify.Alert("WEBCRAWLER", "SETUP AND USAGE",who, "/root/bug.png")
								

		isEnabled := true
  		isColorEnabled := true
		logo := `
{{.AnsiColor.BrightBlue}}████████████████████████████████████████████████████████████████████████████████
{{.AnsiColor.BrightBlue}}████████████████████████████████████████████████████████████████████████████████
{{.AnsiColor.Red}}██████████████████{{.AnsiColor.BrightBlue}}████████████████{{.AnsiColor.Black}}██████████████████████████████{{.AnsiColor.BrightBlue}}████████████████
{{.AnsiColor.Red}}████████████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██████████████
{{.AnsiColor.BrightRed}}████{{.AnsiColor.Red}}██████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Magenta}}██████████████████████{{.AnsiColor.White}}██████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████████████
{{.AnsiColor.BrightRed}}██████████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Magenta}}████████████████{{.AnsiColor.Black}}████{{.AnsiColor.Magenta}}██████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██{{.AnsiColor.Black}}████{{.AnsiColor.BrightBlue}}██████
{{.AnsiColor.BrightRed}}██████████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.Magenta}}██████{{.AnsiColor.White}}██{{.AnsiColor.Black}}████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████
{{.AnsiColor.BrightYellow}}██████████████████{{.AnsiColor.BrightRed}}████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Magenta}}██████{{.AnsiColor.White}}██{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████
{{.AnsiColor.BrightYellow}}██████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightYellow}}██████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}████████{{.AnsiColor.White}}████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████
{{.AnsiColor.BrightYellow}}████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Black}}██{{.AnsiColor.BrightYellow}}████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████
{{.AnsiColor.BrightGreen}}██████████████████{{.AnsiColor.BrightYellow}}██{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Black}}████████{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██
{{.AnsiColor.BrightGreen}}██████████████████████{{.AnsiColor.White}}████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.BrightYellow}}██{{.AnsiColor.White}}██████████{{.AnsiColor.BrightYellow}}██{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██
{{.AnsiColor.BrightGreen}}██████████████████████{{.AnsiColor.Black}}████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Black}}████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██
{{.AnsiColor.Blue}}██████████████████{{.AnsiColor.BrightGreen}}████████{{.AnsiColor.Black}}██████{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}████{{.AnsiColor.White}}████████████████{{.AnsiColor.Magenta}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██
{{.AnsiColor.Blue}}██████████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}████████████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████
{{.AnsiColor.BrightBlue}}██████████████████{{.AnsiColor.Blue}}████{{.AnsiColor.Blue}}██████{{.AnsiColor.Black}}████{{.AnsiColor.White}}██████{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████████████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██████
{{.AnsiColor.BrightBlue}}██████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Black}}████{{.AnsiColor.White}}████████████████████{{.AnsiColor.Black}}██████████████████{{.AnsiColor.BrightBlue}}████████
{{.AnsiColor.BrightBlue}}████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}████████████████████████████████{{.AnsiColor.White}}██{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████████████
{{.AnsiColor.BrightBlue}}████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.BrightBlue}}████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Black}}████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████████████
{{.AnsiColor.BrightBlue}}████████████████████████{{.AnsiColor.Black}}██████{{.AnsiColor.BrightBlue}}████{{.AnsiColor.Black}}██████{{.AnsiColor.BrightBlue}}████████████{{.AnsiColor.Black}}██████{{.AnsiColor.BrightBlue}}████{{.AnsiColor.Black}}██████{{.AnsiColor.BrightBlue}}████████████
████████████████████████████████████████████████████████████████████████████████
{{ .AnsiColor.Default }} {{.AnsiColor.BrightBlue}}███ BY ELMO9AWIMDZ DZ-HACK-LAB {{ .AnsiColor.Default }}`

		banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString(logo))

        	fmt.Println("SEE USAGE ? ⚠️  ");
		usage();
        	os.Exit(1);
    	}


	max, _ := strconv.Atoi(os.Args[2])
			

	delay, _ := strconv.Atoi(os.Args[3])
			

	myCrawler := web.Crawler{ 
			MaxConcurrencyLimit: max, 
			StoragePath: "crawler/storage",  
			CrawlDelay: delay,
		}

	// List of URLS to be crawled as a string array.
	filePath := os.Args[1]
        readFile, err := os.Open(filePath)
  
        if err != nil {
                fmt.Println(err)
        }
        fileScanner := bufio.NewScanner(readFile)
        fileScanner.Split(bufio.ScanLines)
        var fileLines []string
  
        for fileScanner.Scan() {
                fileLines = append(fileLines, fileScanner.Text())
        }
 
        readFile.Close()

	save := flag.String("save", "report", "folder name to save on ")

	for _, line := range fileLines {


                fmt.Println(" 🤖 "+line)


		notify.Alert("WEBCRAWLER", "CRAWL ",line, "/root/bug.png")
				

                alert.Error.Println("⚠️  ATTACK ON => 🤖 "+line)
		fmt.Println(rainbow.Bold(rainbow.Hex("#8E44AD", "CRAWLER WORK ON : "+line)))
		alert.Error.Println("⚠️  SAVE ON => 🤖 "+*save)
		fmt.Printf("[%s] %s\n", aurora.Green("PASSED").String(),line)

		urls := []string{line}
		// Starting the crawler by passing the list of URLs.
		myCrawler.Start(urls)
		
		notify.Alert("WEBCRAWLER", "DONE FOR ",line, "/root/bug.png")
				
	}
	
	alert.Error.Println("⚠️  SAVE ON => 🤖 "+*save)
			

	notify.Alert("WEBCRAWLER", "SAVE ON ",*save, "/root/bug.png")
			

	os.Rename("crawler",*save)
}
