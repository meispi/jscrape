package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
)

var reslist []string

func findsubs(jobs chan string, wg *sync.WaitGroup, re *regexp.Regexp){
	defer wg.Done()
	for url := range jobs {
		res, err := http.Get(url)
		if err != nil {
			log.Println("Failed! : ", err)
		}
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Print("Error : ", err)
		}

		matches := re.FindAllString(string(data),-1)
		for _, i := range matches {
			reslist = append(reslist, i)
		}
	}
}

func main() {
	con := flag.Int("c", 8, "number of threads")
	word := flag.String("w","","Enter your word (e.g. uber)")
	all := flag.Bool("all",false,"looks for *target* (and not just *.target.*) (default false)")
	flag.Parse()

	if *word != "" {

		jobs := make(chan string)
		resmap := make(map[string]int)

		go func () {
			sc := bufio.NewScanner(os.Stdin)
			for sc.Scan() {
				jobs <- sc.Text()
			}
			close(jobs)
		}()

		wg := &sync.WaitGroup{}

		if *all {
			re := regexp.MustCompile(`[A-Za-z0-9\-\.]*`+*word+`[A-Za-z0-9\.\-]*\.(com|net|network|io|org)[A-Za-z0-9\-\.]*`)

			for i := 0; i < *con; i++ {
				wg.Add(1)
				go findsubs(jobs, wg, re)
			}
		} else {
			re := regexp.MustCompile(`[A-Za-z0-9\-\.]*\.`+*word+`\.(com|net|network|io|org)[A-Za-z0-9\-\.]*`)

			for i := 0; i < *con; i++ {
				wg.Add(1)
				go findsubs(jobs, wg, re)
			}
		}
		
		wg.Wait()

		for _, subs := range reslist {
			resmap[subs] = 1
		}
		for str := range resmap {
			fmt.Println(str)
		}
	} else {
		flag.Usage()
	}
}