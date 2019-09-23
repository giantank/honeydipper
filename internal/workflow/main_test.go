// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.

package workflow

import (
	"os"
	"testing"
	"time"

	"github.com/ghodss/yaml"
	"github.com/golang/mock/gomock"
	"github.com/honeydipper/honeydipper/internal/config"
	"github.com/honeydipper/honeydipper/internal/daemon"
	"github.com/honeydipper/honeydipper/internal/workflow/mock_workflow"
	"github.com/honeydipper/honeydipper/pkg/dipper"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if dipper.Logger == nil {
		//f, _ := os.OpenFile(os.DevNull, os.O_APPEND, 0777)
		f, _ := os.Create("test.log")
		defer f.Close()
		dipper.GetLogger("test service", "DEBUG", f, f)
	}
	os.Exit(m.Run())
}

var mockHelper *mock_workflow.MockSessionStoreHelper
var emptyLabels map[string]string
var store *SessionStore

// every test case has following fields
// - workflow: an workflow object to be executed using StartSession call
// - msg:      an event message to be passed to initiate the workflow session
// - ctx:      the ctx generated by the event used for initiating the workflow
// - asserts:  a func() to be called to assert the success or failure of the test
//             after an initial message is send out from the session
//
// - steps:    during the execution, multiple msgs are passed in and out of the session
//             each step contains following
//
//   * msg:        the message to be passed into the session as using ContinueSession call
//   * sessionID:  the sessionID used to match message with a session
//   * ctx:        the exported ctx from the function call
//   * asserts:    a func() to assert the success of failure of the test after sending
//                 the message
//
// The test flow:
//  1. Load config and ensure validity
//  2. Set up assertion for initial step
//  3. Run `StartSession` with given msg, ctx and workflow to
//  4. Check if function crashes
//  5. Wait and check if all go routines complete
//  6. Validate assertions
//
//  7. Set up assertion for the new step
//  8. Run `ContinueSession` or `ResumeSession` for step
//  9. Check if function crashes
//  10. Wait and check if all go routines complete
//  11. Validate assertions
//
//  12. repeat for all steps

func syntheticTest(t *testing.T, configStr string, testcase map[string]interface{}) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHelper = mock_workflow.NewMockSessionStoreHelper(ctrl)
	store = NewSessionStore(mockHelper)
	defer delete(dipper.IDMapMetadata, &store.sessions)

	testDataSet := &config.DataSet{}
	err := yaml.UnmarshalStrict([]byte(configStr), testDataSet, yaml.DisallowUnknownFields)
	assert.Nil(t, err, "test config")
	testConfig := &config.Config{DataSet: testDataSet}

	var step int
	var teststep map[string]interface{}

	testStartFunc := func() {
		step = -1
		signal := make(chan int, 1)
		go func() {
			store.StartSession(testcase["workflow"].(*config.Workflow), testcase["msg"].(*dipper.Message), testcase["ctx"].(map[string]interface{}))
			daemon.Children.Wait()
			signal <- 1
		}()
		select {
		case <-signal:
		case <-time.After(1 * time.Second):
			panic("timeout due to go routine leak")
		}
	}

	testContinueFunc := func() {
		signal := make(chan int, 1)
		go func() {
			store.ContinueSession(teststep["sessionID"].(string), teststep["msg"].(*dipper.Message), teststep["ctx"].([]map[string]interface{}))
			daemon.Children.Wait()
			signal <- 1
		}()
		select {
		case <-signal:
		case <-time.After(1 * time.Second):
			panic("timeout due to go routine leak")
		}
	}

	testResumeFunc := func() {
		signal := make(chan int, 1)
		go func() {
			store.ResumeSession(teststep["key"].(string), teststep["msg"].(*dipper.Message))
			daemon.Children.Wait()
			signal <- 1
		}()
		select {
		case <-signal:
		case <-time.After(1 * time.Second):
			panic("timeout due to go routine leak")
		}
	}

	mockHelper = mock_workflow.NewMockSessionStoreHelper(ctrl)
	store.Helper = mockHelper
	mockHelper.EXPECT().GetConfig().AnyTimes().Return(testConfig)
	if assertFunc, ok := testcase["asserts"]; ok {
		assertFunc.(func())()
	}
	if shouldPanic, ok := testcase["panic"]; ok && shouldPanic.(bool) {
		assert.Panics(t, testStartFunc, "expecting panic at starting test case")
	} else {
		assert.NotPanics(t, testStartFunc, "expecting not panic at starting test case")
	}

	testSteps := testcase["steps"].([]map[string]interface{})
	for step, teststep = range testSteps {
		mockHelper = mock_workflow.NewMockSessionStoreHelper(ctrl)
		store.Helper = mockHelper
		mockHelper.EXPECT().GetConfig().AnyTimes().Return(testConfig)
		if assertFunc, ok := teststep["asserts"]; ok {
			assertFunc.(func())()
		}
		nextFunc := testContinueFunc
		if resuming, ok := teststep["resuming"]; ok && resuming.(bool) {
			nextFunc = testResumeFunc
		}
		if shouldPanic, ok := teststep["panic"]; ok && shouldPanic.(bool) {
			assert.Panics(t, nextFunc, "expecting panic at step %d", step)
		} else {
			assert.NotPanics(t, nextFunc, "expecting not panic at step %d", step)
		}
	}
	assert.Equal(t, len(testSteps)-1, step, "expecting number of processed steps")
	assert.Equal(t, 0, len(store.sessions), "expecting all session to be completed")
}
