// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"

	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
)

const (
	apiKey = "AIzaSyAYiKiTfWtT1PWPLWKlqHMPfQtV74XApKk"
	cx     = "partner-pub-8773012213497498:fi4b0xyyr1e"
	query  = "elezioni europee"
)

func customSearchMain() {
	fmt.Println("inside customSearchMain")
	client := &http.Client{Transport: &transport.APIKey{Key: apiKey}}

	svc, err := customsearch.New(client)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := svc.Cse.Siterestrict.List(query).Cx(cx).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("before for loop resp.Items", resp.Items)
	for i, result := range resp.Items {
		fmt.Printf("#%d: %s\n", i+1, result.Title)
		fmt.Printf("\t%s\n", result.Snippet)
	}
}

func main() {
	customSearchMain()
}