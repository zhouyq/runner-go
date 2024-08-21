// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package environ

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInvalidSemver(t *testing.T) {
	a := versions("this is an invalid version")
	b := map[string]string{"GITFOX_SEMVER_ERROR": "this is an invalid version is not in dotted-tri format"}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Errorf("Unexpected semver variables")
		t.Log(diff)
	}
}

func TestSemver(t *testing.T) {
	a := versions("v1.2.3-alpha+001")
	b := map[string]string{
		"GITFOX_SEMVER":            "1.2.3-alpha+001",
		"GITFOX_SEMVER_MAJOR":      "1",
		"GITFOX_SEMVER_MINOR":      "2",
		"GITFOX_SEMVER_PATCH":      "3",
		"GITFOX_SEMVER_SHORT":      "1.2.3",
		"GITFOX_SEMVER_PRERELEASE": "alpha",
		"GITFOX_SEMVER_BUILD":      "001",
	}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Errorf("Unexpected semver variables")
		t.Log(diff)
	}
}
