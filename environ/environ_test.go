// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package environ

import (
	"testing"

	"github.com/drone/drone-go/drone"

	"github.com/google/go-cmp/cmp"
)

func TestRepo(t *testing.T) {
	v := &drone.Repo{
		ID:         1,
		UID:        "2",
		UserID:     3,
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
		SCM:        "git",
		HTTPURL:    "https://github.com/octocat/hello-world.git",
		SSHURL:     "git@github.com:octocat/hello-world",
		Link:       "https://github.com/octocat/hello-world",
		Branch:     "master",
		Private:    true,
		Visibility: "internal",
	}
	a := Repo(v)
	b := map[string]string{
		"GITFOX_REPO":            "octocat/hello-world",
		"GITFOX_REPO_SCM":        "git",
		"GITFOX_REPO_OWNER":      "octocat",
		"GITFOX_REPO_NAMESPACE":  "octocat",
		"GITFOX_REPO_NAME":       "hello-world",
		"GITFOX_REPO_LINK":       "https://github.com/octocat/hello-world",
		"GITFOX_REPO_BRANCH":     "master",
		"GITFOX_REMOTE_URL":      "https://github.com/octocat/hello-world.git",
		"GITFOX_GIT_HTTP_URL":    "https://github.com/octocat/hello-world.git",
		"GITFOX_GIT_SSH_URL":     "git@github.com:octocat/hello-world",
		"GITFOX_REPO_VISIBILITY": "internal",
		"GITFOX_REPO_PRIVATE":    "true",

		"CI_REMOTE_URL":   "https://github.com/octocat/hello-world.git",
		"CI_REPO":         "octocat/hello-world",
		"CI_REPO_LINK":    "https://github.com/octocat/hello-world",
		"CI_REPO_NAME":    "octocat/hello-world",
		"CI_REPO_PRIVATE": "true",
		"CI_REPO_REMOTE":  "https://github.com/octocat/hello-world.git",
	}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Fail()
		t.Log(diff)
	}
}

func TestBuild(t *testing.T) {
	v := &drone.Build{
		Trigger:      "root",
		Source:       "develop",
		Target:       "master",
		After:        "762941318ee16e59dabbacb1b4049eec22f0d303",
		Before:       "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		Ref:          "refs/pull/32/head",
		Link:         "https://github.com/octocat/Hello-World/commit/762941318ee16e59dabbacb1b4049eec22f0d303",
		Title:        "feat: update README",
		Message:      "updated README",
		Author:       "octocat",
		AuthorAvatar: "https://avatars0.githubusercontent.com/u/583231",
		AuthorEmail:  "octocat@github.com",
		AuthorName:   "The Octocat",
		Number:       1,
		Parent:       2,
		Event:        "pull_request",
		Action:       "opened",
		Deploy:       "prod",
		DeployID:     235634642,
		Debug:        true,
		Status:       drone.StatusFailing,
		Created:      1561421740,
		Started:      1561421746,
		Finished:     1561421753,
		Stages: []*drone.Stage{
			{
				Name:   "backend",
				Number: 1,
				Status: drone.StatusPassing,
			},
			{
				Name:   "frontend",
				Number: 2,
				Status: drone.StatusFailing,
			},
		},
	}

	a := Build(v)
	b := map[string]string{
		"GITFOX_BRANCH":               "master",
		"GITFOX_BUILD_NUMBER":         "1",
		"GITFOX_BUILD_PARENT":         "2",
		"GITFOX_BUILD_STATUS":         "failure",
		"GITFOX_BUILD_EVENT":          "pull_request",
		"GITFOX_BUILD_DEBUG":          "true",
		"GITFOX_BUILD_ACTION":         "opened",
		"GITFOX_BUILD_CREATED":        "1561421740",
		"GITFOX_BUILD_STARTED":        "1561421746",
		"GITFOX_BUILD_FINISHED":       "1561421753",
		"GITFOX_COMMIT":               "762941318ee16e59dabbacb1b4049eec22f0d303",
		"GITFOX_COMMIT_BEFORE":        "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		"GITFOX_COMMIT_AFTER":         "762941318ee16e59dabbacb1b4049eec22f0d303",
		"GITFOX_COMMIT_BRANCH":        "master",
		"GITFOX_COMMIT_LINK":          "https://github.com/octocat/Hello-World/commit/762941318ee16e59dabbacb1b4049eec22f0d303",
		"GITFOX_COMMIT_MESSAGE":       "updated README",
		"GITFOX_COMMIT_REF":           "refs/pull/32/head",
		"GITFOX_COMMIT_AUTHOR":        "octocat",
		"GITFOX_COMMIT_AUTHOR_AVATAR": "https://avatars0.githubusercontent.com/u/583231",
		"GITFOX_COMMIT_AUTHOR_EMAIL":  "octocat@github.com",
		"GITFOX_COMMIT_AUTHOR_NAME":   "The Octocat",
		"GITFOX_COMMIT_SHA":           "762941318ee16e59dabbacb1b4049eec22f0d303",
		"GITFOX_DEPLOY_TO":            "prod",
		"GITFOX_DEPLOY_ID":            "235634642",
		"GITFOX_FAILED_STAGES":        "frontend",
		"GITFOX_PULL_REQUEST":         "32",
		"GITFOX_PULL_REQUEST_TITLE":   "feat: update README",
		"GITFOX_SOURCE_BRANCH":        "develop",
		"GITFOX_TARGET_BRANCH":        "master",
		"GITFOX_BUILD_TRIGGER":        "root",

		"CI_BUILD_CREATED":        "1561421740",
		"CI_BUILD_EVENT":          "pull_request",
		"CI_BUILD_FINISHED":       "1561421753",
		"CI_BUILD_LINK":           "https://github.com/octocat/Hello-World/commit/762941318ee16e59dabbacb1b4049eec22f0d303",
		"CI_BUILD_NUMBER":         "1",
		"CI_BUILD_STARTED":        "1561421746",
		"CI_BUILD_STATUS":         "failure",
		"CI_BUILD_TARGET":         "prod",
		"CI_COMMIT_AUTHOR":        "octocat",
		"CI_COMMIT_AUTHOR_AVATAR": "https://avatars0.githubusercontent.com/u/583231",
		"CI_COMMIT_AUTHOR_EMAIL":  "octocat@github.com",
		"CI_COMMIT_AUTHOR_NAME":   "The Octocat",
		"CI_COMMIT_BRANCH":        "master",
		"CI_COMMIT_MESSAGE":       "updated README",
		"CI_COMMIT_REF":           "refs/pull/32/head",
		"CI_COMMIT_SHA":           "762941318ee16e59dabbacb1b4049eec22f0d303",
		"CI_PARENT_BUILD_NUMBER":  "2",
	}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Fail()
		t.Log(diff)
	}

	v.Started = 0
	v.Finished = 0
	a = Build(v)
	if a["GITFOX_BUILD_STARTED"] == "0" {
		t.Errorf("Expect non-zero started time")
	}
	if a["GITFOX_BUILD_FINISHED"] == "0" {
		t.Errorf("Expect non-zero stopped time")
	}

	v.Ref = "refs/tags/v1.2.3"
	a = Build(v)
	if a["GITFOX_TAG"] != "v1.2.3" {
		t.Errorf("Expect tag extraced from ref")
	}
	if a["GITFOX_SEMVER"] != "1.2.3" {
		t.Errorf("Expect semver from ref")
	}
	if a["GITFOX_SEMVER_MAJOR"] != "1" {
		t.Errorf("Expect semver major")
	}
	if a["GITFOX_SEMVER_MINOR"] != "2" {
		t.Errorf("Expect semver minor")
	}
	if a["GITFOX_SEMVER_PATCH"] != "3" {
		t.Errorf("Expect semver patch")
	}
}

func TestSystem(t *testing.T) {
	v := &drone.System{
		Proto:   "http",
		Host:    "gitfox.company.com",
		Link:    "http://gitfox.company.com",
		Version: "v1.0.0",
	}
	a := System(v)
	b := map[string]string{
		"CI":                    "true",
		"GITFOX":                 "true",
		"GITFOX_SYSTEM_HOST":     "gitfox.company.com",
		"GITFOX_SYSTEM_HOSTNAME": "gitfox.company.com",
		"GITFOX_SYSTEM_PROTO":    "http",
		"GITFOX_SYSTEM_VERSION":  "v1.0.0",
	}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Fail()
		t.Log(diff)
	}
}

func TestStep(t *testing.T) {
	v := &drone.Step{
		Name:   "clone",
		Number: 1,
	}
	a := Step(v)
	b := map[string]string{
		"GITFOX_STEP_NAME":   "clone",
		"GITFOX_STEP_NUMBER": "1",
	}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Fail()
		t.Log(diff)
	}
}

func TestStage(t *testing.T) {
	v := &drone.Stage{
		Kind:      "pipeline",
		Type:      "docker",
		Name:      "deploy",
		Number:    1,
		Machine:   "laptop",
		OS:        "linux",
		Arch:      "arm64",
		Variant:   "7",
		Version:   2004,
		Status:    drone.StatusFailing,
		Started:   1561421746,
		Stopped:   1561421753,
		DependsOn: []string{"backend", "frontend"},
		Steps: []*drone.Step{
			{
				Name:   "clone",
				Number: 1,
				Status: drone.StatusPassing,
			},
			{
				Name:   "test",
				Number: 2,
				Status: drone.StatusFailing,
			},
		},
	}

	a := Stage(v)
	b := map[string]string{
		"GITFOX_STAGE_KIND":       "pipeline",
		"GITFOX_STAGE_TYPE":       "docker",
		"GITFOX_STAGE_NAME":       "deploy",
		"GITFOX_STAGE_NUMBER":     "1",
		"GITFOX_STAGE_MACHINE":    "laptop",
		"GITFOX_STAGE_OS":         "linux",
		"GITFOX_STAGE_ARCH":       "arm64",
		"GITFOX_STAGE_VARIANT":    "7",
		"GITFOX_STAGE_VERSION":    "2004",
		"GITFOX_STAGE_STATUS":     "failure",
		"GITFOX_STAGE_STARTED":    "1561421746",
		"GITFOX_STAGE_FINISHED":   "1561421753",
		"GITFOX_STAGE_DEPENDS_ON": "backend,frontend",
		"GITFOX_FAILED_STEPS":     "test",
		"GITFOX_CARD_PATH":        "/dev/stdout",
	}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Fail()
		t.Log(diff)
	}

	v.Started = 0
	v.Stopped = 0
	a = Stage(v)
	if a["GITFOX_STAGE_STARTED"] == "0" {
		t.Errorf("Expect non-zero started time")
	}
	if a["GITFOX_STAGE_FINISHED"] == "0" {
		t.Errorf("Expect non-zero stopped time")
	}
}

func TestLink(t *testing.T) {
	sys := &drone.System{
		Proto: "http",
		Host:  "drone.company.com",
	}
	build := &drone.Build{Number: 42}
	repo := &drone.Repo{Slug: "octocat/hello-world"}
	a := Link(repo, build, sys)
	b := map[string]string{
		"GITFOX_BUILD_LINK": "http://drone.company.com/octocat/hello-world/42",
	}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Fail()
		t.Log(diff)
	}
}

func TestSlice(t *testing.T) {
	v := map[string]string{
		"CI":    "true",
		"GITFOX": "true",
	}
	a := Slice(v)
	b := []string{"CI=true", "GITFOX=true"}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Fail()
		t.Log(diff)
	}
}

func TestCombine(t *testing.T) {
	v1 := map[string]string{
		"CI":    "true",
		"GITFOX": "true",
	}
	v2 := map[string]string{
		"CI":                    "false",
		"GITFOX_SYSTEM_HOST":     "gitfox.company.com",
		"GITFOX_SYSTEM_HOSTNAME": "gitfox.company.com",
		"GITFOX_SYSTEM_PROTO":    "http",
		"GITFOX_SYSTEM_VERSION":  "v1.0.0",
	}
	a := Combine(v1, v2)
	b := map[string]string{
		"CI":                    "false",
		"GITFOX":                 "true",
		"GITFOX_SYSTEM_HOST":     "gitfox.company.com",
		"GITFOX_SYSTEM_HOSTNAME": "gitfox.company.com",
		"GITFOX_SYSTEM_PROTO":    "http",
		"GITFOX_SYSTEM_VERSION":  "v1.0.0",
	}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Fail()
		t.Log(diff)
	}
}

func Test_isStageFailing(t *testing.T) {
	s := &drone.Stage{
		Status: drone.StatusPassing,
		Steps: []*drone.Step{
			{
				Status: drone.StatusPassing,
			},
			{
				ErrIgnore: true,
				Status:    drone.StatusFailing,
			},
		},
	}
	if isStageFailing(s) {
		t.Errorf("Expect stage not failing if ErrIgnore")
	}

	s = &drone.Stage{
		Status: drone.StatusFailing,
	}
	if isStageFailing(s) == false {
		t.Errorf("Expect stage failing")
	}

	s = &drone.Stage{
		Status: drone.StatusRunning,
		Steps: []*drone.Step{
			{
				Status: drone.StatusPassing,
			},
			{
				ErrIgnore: false,
				Status:    drone.StatusFailing,
			},
		},
	}
	if isStageFailing(s) == false {
		t.Errorf("Expect stage failing if step failing")
	}
}

func Test_isBuildFailing(t *testing.T) {
	v := &drone.Build{
		Status: drone.StatusPassing,
	}
	if isBuildFailing(v) == true {
		t.Errorf("Expect build passing")
	}

	v.Status = drone.StatusFailing
	if isBuildFailing(v) == false {
		t.Errorf("Expect build failing")
	}

	v.Status = drone.StatusRunning
	v.Stages = []*drone.Stage{
		{Status: drone.StatusPassing},
		{Status: drone.StatusFailing},
	}
	if isBuildFailing(v) == false {
		t.Errorf("Expect build failing if stage failing")
	}

	v.Stages = []*drone.Stage{
		{Status: drone.StatusPassing},
		{Status: drone.StatusRunning},
		{Status: drone.StatusPending},
	}
	if isBuildFailing(v) == true {
		t.Errorf("Expect build passing if all stages passing")
	}
}

func Test_failedSteps(t *testing.T) {
	s := &drone.Stage{
		Status: drone.StatusRunning,
		Steps: []*drone.Step{
			{
				Name:   "clone",
				Status: drone.StatusPassing,
			},
			{
				Name:   "test",
				Status: drone.StatusFailing,
			},
			{
				Name:   "integration",
				Status: drone.StatusFailing,
			},
			{
				Name:      "experimental",
				ErrIgnore: true,
				Status:    drone.StatusFailing,
			},
		},
	}
	a, b := failedSteps(s), []string{"test", "integration"}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Fail()
		t.Log(diff)
	}
}

func Test_failedStages(t *testing.T) {
	v := &drone.Build{
		Status: drone.StatusRunning,
		Stages: []*drone.Stage{
			{
				Name:   "step_blocked",
				Status: drone.StatusBlocked,
			},
			{
				Name:   "step_declined",
				Status: drone.StatusDeclined,
			},
			{
				Name:   "step_error",
				Status: drone.StatusError,
			},
			{
				Name:   "step_failing",
				Status: drone.StatusFailing,
			},
			{
				Name:   "step_killed",
				Status: drone.StatusKilled,
			},
			{
				Name:   "step_passing",
				Status: drone.StatusPassing,
			},
			{
				Name:   "step_pending",
				Status: drone.StatusPending,
			},
			{
				Name:   "step_running",
				Status: drone.StatusRunning,
			},
			{
				Name:   "step_skipped",
				Status: drone.StatusSkipped,
			},
			{
				Name:   "step_waiting",
				Status: drone.StatusWaiting,
			},
		},
	}
	a, b := failedStages(v), []string{"step_error", "step_failing", "step_killed"}
	if diff := cmp.Diff(a, b); diff != "" {
		t.Fail()
		t.Log(diff)
	}
}
