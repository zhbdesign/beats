// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build integration

package container_registry

import (
	"testing"

	"github.com/elastic/beats/v7/x-pack/metricbeat/module/azure/test"

	"github.com/stretchr/testify/assert"

	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"

	// Register input module and metricset
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/azure/monitor"
)

func TestFetchMetricset(t *testing.T) {
	config, err := test.GetConfig("container_registry")
	if err != nil {
		t.Skip("Skipping TestFetch: " + err.Error())
	}
	metricSet := mbtest.NewReportingMetricSetV2Error(t, config)
	events, errs := mbtest.ReportingFetchV2Error(metricSet)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)
}

func TestData(t *testing.T) {
	config, err := test.GetConfig("container_registry")
	if err != nil {
		t.Skip("Skipping TestFetch: " + err.Error())
	}
	metricSet := mbtest.NewFetcher(t, config)
	metricSet.WriteEvents(t, "/")
}
