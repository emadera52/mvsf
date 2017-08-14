/*
 * Copyright 2017 Stephen D. Wood. All rights reserved.
 * Use of this source code is governed by the MIT License,
 * Version 2.0. For details see the LICENSE file.
 */

package main

import (
	// Standard library
	"fmt"
	"os"
	"strconv"
	"time"

	// Internal
	"mvsf/models"
	"mvsf/datactrl"
)

// Timer executes one of four commands "limit" times. Each command produces the
// same result. They create models.CmntData based on models.DefaultCD but with
// modifications. 

//NOTE: models.DefaultDC must NOT be modified in the process. After the final
// loop, modified default values and the original models.DefaultCD are shown.

// Added option to repeat Timer 40 times and calc average Durration if "loop"
// is true. Timer does not print details when "loop" is true.

// For "external" tests, definitions are in comments.go in the datactrl pkg.
// For "local" tests, definitions are in models.go in the models pkg. itself
func Timer(cmd string, limit int, loop bool) time.Duration {
	var elapsed time.Duration
	var defCD *models.CmntData

	// Time to execute "limit" loops using external methods
	if cmd == "meth" {
		tstart := time.Now()
		
		// datactrl.Rcd provides access to remote models.CmntData methods
		dcd := &datactrl.Rcd 
		for i := 0; i < limit; i++ {
			defCD = dcd.DefaultComment(i)
		}
		tfinish := time.Now()
		elapsed = tfinish.Sub(tstart)
	}
	

	// Time to execute "limit" loops using external functions
	if cmd == "func" {
		tstart := time.Now()
		for i := 0; i < limit; i++ {
			defCD = datactrl.DefaultComment(i)
		}
		tfinish := time.Now()
		elapsed = tfinish.Sub(tstart)
	}

	
	// Time to execute "limit" loops using local methods
	if cmd == "method" {
		tstart := time.Now()
		dcd := models.CmntData{}
		for i := 0; i < limit; i++ {
			defCD = dcd.DefaultComment(i)
		}
		tfinish := time.Now()
		elapsed = tfinish.Sub(tstart)
	}

	
	// Time to execute "limit" loops using local functions
	if cmd == "function" {
		tstart := time.Now()
		for i := 0; i < limit; i++ {
			defCD = models.DefaultComment(i)
		}
		tfinish := time.Now()
		elapsed = tfinish.Sub(tstart)
	}

	if defCD == nil {
		fmt.Println("Valid commands are 'meth', 'func', 'method' and 'function'")
	} else {
		if !loop {
			fmt.Println(defCD, "\n")
			fmt.Println(models.DefaultCD, "\n")	
		}
	}

	return elapsed
}

func main() {

	cmd := os.Args[1]
	cnt := "100000"
	loop := false
	
	if len(os.Args) > 2 {
		cnt = os.Args[2]
	}
	limit, err := strconv.Atoi(cnt)
	// Convert errors and values < 100,000 to 100,000
	if err != nil || limit < 100000 {
		limit = 100000
	}

	// Any third argument triggers looping
	if len(os.Args) > 3 {
		loop = true
	}

	if loop {
		var elapsed time.Duration 
		var totElapsed time.Duration
		
		for i := 0; i < 40; i++ {
			elapsed = Timer(cmd, limit, loop)
			totElapsed = totElapsed + elapsed
		}
		fmt.Printf("Average Duration: %v\nNative", totElapsed / 40)
	} else {
		fmt.Printf("Duration: %v\nNative", Timer(cmd, limit, loop))
	}
}
