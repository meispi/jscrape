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
		resmap := make(map[string]int)
		for _, i := range matches {
			resmap[i] = 1
		}
		for str := range resmap {
			fmt.Println(str)
		}
	}
}

func main() {
	con := flag.Int("c", 8, "number of threads")
	word := flag.String("w","","Enter your word (e.g. uber)")
	flag.Parse()

	if *word != "" {
		re := regexp.MustCompile(`[A-Za-z0-9\-\.]*.`+*word+`[A-Za-z0-9\.\-]*\.(com|net|network|io|org)`)

		jobs := make(chan string)

		go func () {
			sc := bufio.NewScanner(os.Stdin)
			for sc.Scan() {
				jobs <- sc.Text()
			}
			close(jobs)
		}()

		wg := &sync.WaitGroup{}

		for i := 0; i < *con; i++ {
			wg.Add(1)
			go findsubs(jobs, wg, re)
		}
		wg.Wait()
	} else {
		flag.Usage()
	}
}