// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml_test

import (
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

func TestCT_RPrOriginalConstructor(t *testing.T) {
	v := wordprocessingml.NewCT_RPrOriginal()
	if v == nil {
		t.Errorf("wordprocessingml.NewCT_RPrOriginal must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wordprocessingml.CT_RPrOriginal should validate: %s", err)
	}
}
