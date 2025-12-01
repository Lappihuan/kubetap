// Copyright 2020 Soluble Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// +build !race

package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSpinner(t *testing.T) {
	// Disable pterm output during tests
	oldEnv := os.Getenv("CI")
	defer func() {
		if oldEnv != "" {
			os.Setenv("CI", oldEnv)
		} else {
			os.Unsetenv("CI")
		}
	}()
	os.Setenv("CI", "true")

	// Test spinner creation and basic operations
	t.Run("spinner_lifecycle", func(t *testing.T) {
		spinner := NewSpinner("Testing...")
		require.NotNil(t, spinner)
		require.NotNil(t, spinner.spinner)
		require.NotNil(t, spinner.done)

		// Give spinner goroutine time to start
		time.Sleep(20 * time.Millisecond)

		// Test update
		spinner.Update("Updated message")
		time.Sleep(20 * time.Millisecond)

		// Test stop
		spinner.Stop("Done!")
		time.Sleep(100 * time.Millisecond)

		// Verify stop was called
		require.True(t, spinner.stopped)
		require.Nil(t, spinner.spinner)
	})

	t.Run("spinner_fail", func(t *testing.T) {
		spinner := NewSpinner("Testing...")
		require.NotNil(t, spinner)

		// Give spinner goroutine time to start
		time.Sleep(20 * time.Millisecond)

		// Test fail
		spinner.Fail("Failed!")
		time.Sleep(100 * time.Millisecond)

		// Verify fail was called
		require.True(t, spinner.stopped)
		require.Nil(t, spinner.spinner)
	})

	t.Run("spinner_idempotent_stop", func(t *testing.T) {
		spinner := NewSpinner("Testing...")

		// Give spinner goroutine time to start
		time.Sleep(20 * time.Millisecond)

		// Call stop multiple times - should not panic
		spinner.Stop("First stop")
		time.Sleep(50 * time.Millisecond)

		spinner.Stop("Second stop")
		time.Sleep(50 * time.Millisecond)

		require.True(t, spinner.stopped)
	})
}

