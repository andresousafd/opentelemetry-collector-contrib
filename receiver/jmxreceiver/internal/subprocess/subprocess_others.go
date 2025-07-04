// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !linux

package subprocess // import "github.com/andresousafd/opentelemetry-collector-contrib/receiver/jmxreceiver/internal/subprocess"

import (
	"os/exec"
)

func applyOSSpecificCmdModifications(_ *exec.Cmd) {}
