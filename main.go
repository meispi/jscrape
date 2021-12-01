package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sync"
)

var reslist []string

func finds3(jobs chan string, wg *sync.WaitGroup, re *regexp.Regexp){
	defer wg.Done()
	for url := range jobs {
		res, err := http.Get(url)
		if err != nil {
			continue
		}
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			continue
		}

		matches := re.FindAllString(string(data),-1)
		for _, i := range matches {
			reslist = append(reslist, i)
		}
	}
}

func findsubs(jobs chan string, wg *sync.WaitGroup, re *regexp.Regexp){
	defer wg.Done()
	for url := range jobs {
		res, err := http.Get(url)
		if err != nil {
			continue
		}
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			continue
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
	s3 := flag.Bool("s3",false,"looks for *.s3.amazonaws.com only (default false)")
	flag.Parse()

	if *word != "" && *s3 == false {

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
	} else if *s3 {
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
		
		re := regexp.MustCompile(`[A-Za-z0-9\.\-]*\.s3\.amazonaws\.com`)

		for i := 0; i < *con; i++ {
			wg.Add(1)
			go finds3(jobs, wg, re)
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