/*
 * Copyright 2017 Stephen D. Wood. All rights reserved.
 * Use of this source code is governed by the MIT License.
 * For details see the LICENSE file.
 */

package models

import (
	// Standard library
	"strconv"
	"time"
	
)

type CmntData struct {
	ID          int
	Slug        string
	Title       string
	PageID      int
	Text		string
	DateCreated time.Time
	CreatedBy   int
}

var DefaultCD = CmntData{ID: 100, Title: "Demo Comment",
	Text: "This is a test comment", PageID: 1000, CreatedBy: 4242}
	

// Local methods called by "method" test
	
func (cd *CmntData) NewComment(ncd CmntData) *CmntData {

	ncd.DateCreated = time.Now()
	
	return &ncd
}

func (cd *CmntData) DefaultComment(cnt int) *CmntData {

	dcd := DefaultCD
	dcd.Slug = "demo-comment-" + strconv.Itoa(cnt + 1)
	
	return cd.NewComment(dcd)
}


// Local functions called by "function" test

func NewComment(cd CmntData) *CmntData {

	cd.DateCreated = time.Now()

	return &cd
}

func DefaultComment(cnt int) *CmntData {

	cd := DefaultCD 
	cd.Slug = "demo-comment-" + strconv.Itoa(cnt + 1)

	return NewComment(cd)
}

