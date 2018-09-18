/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package amalgomated

import (
	"context"

	"github.com/pkg/errors"

	"github.com/sniperkit/snk.fork.palantir-godel-dep-plugin/generated_src/internal/github.com/golang/dep/gps"
)

// TODO solve failures can be really creative - we need to be similarly creative
// in handling them and informing the user appropriately
func handleAllTheFailuresOfTheWorld(err error) error {
	switch errors.Cause(err) {
	case context.Canceled, context.DeadlineExceeded, gps.ErrSourceManagerIsReleased:
		return nil
	}

	return errors.Wrap(err, "Solving failure")
}
