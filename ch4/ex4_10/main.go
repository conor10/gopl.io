// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var oneMonthOrLess = make([]*github.Issue, 0, 50)
	var oneYearOrLess = make([]*github.Issue, 0, 50)
	var oneYearOrMore = make([]*github.Issue, 0, 50)

	now := time.Now()
	oneMonth := now.AddDate(0, -1, 0)
	oneYear := now.AddDate(-1, 0, 0)

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		created := item.CreatedAt

		if created.After(oneMonth) {
			oneMonthOrLess = append(oneMonthOrLess, item)
		} else if created.After(oneYear) {
			oneYearOrLess = append(oneYearOrLess, item)
		} else {
			oneYearOrMore = append(oneYearOrMore, item)
		}
	}

	fmt.Printf("Items created in the last month:\n")
	print(oneMonthOrLess)

	fmt.Printf("\nItems created in the last year:\n")
	print(oneYearOrLess)

	fmt.Printf("\nItems created over a year ago:\n")
	print(oneYearOrMore)
}

func print(issues []*github.Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format(time.RFC822), item.User.Login, item.Title)
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
