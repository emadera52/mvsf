/*
 * Copyright 2017 Stephen D. Wood. All rights reserved.
 * Use of this source code is governed by the MIT License.
 * For details see the LICENSE file.
 */

package datactrl

import (
	// Standard library
	"strconv"
	"time"
	
	// Internal
	"mvsf/models"
)

// Types and method definitions used to add methods to models.CmntData

type roCmntData struct {
	ModelCD models.CmntData
}

// This is possible thanks to powers granted to an empty struct 
type roCD struct{}

// roComment allows us to effectively add methods to models.CmntData. All data
// and "local" methods, if any, defined for models.CmntData are accessible.
type roComment struct {
	*roCmntData
	*roCD
}

// Other packages can use Rcd to access methods added to models.CmntData
var	Rcd = roComment {
		roCmntData: &roCmntData{},
		roCD: &roCD{},
	}
	
// Remote methods called by "meth" test
func (cd *roComment) NewComment(ncd models.CmntData) *models.CmntData {

	cd.roCmntData.ModelCD = ncd
	cd.ModelCD.DateCreated = time.Now()
	
	return &cd.ModelCD
}

func (cd *roComment) DefaultComment(cnt int) *models.CmntData {

	dcd := models.DefaultCD
	dcd.Slug = "demo-comment-" + strconv.Itoa(cnt + 1)
	
	return cd.NewComment(dcd)
}


// Remote functions called by "func" test

func NewComment(cd models.CmntData) *models.CmntData {

	cd.DateCreated = time.Now()

	return &cd
}

func DefaultComment(cnt int) *models.CmntData {

	cd := models.DefaultCD 
	cd.Slug = "demo-comment-" + strconv.Itoa(cnt + 1)

	return NewComment(cd)
}

