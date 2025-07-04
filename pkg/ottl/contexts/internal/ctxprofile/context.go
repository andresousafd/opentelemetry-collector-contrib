// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxprofile // import "github.com/andresousafd/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxprofile"
import (
	"go.opentelemetry.io/collector/pdata/pprofile"
)

const (
	Name   = "profile"
	DocRef = "https://github.com/andresousafd/opentelemetry-collector-contrib/tree/main/pkg/ottl/contexts/ottlprofile"
)

type Context interface {
	GetProfile() pprofile.Profile
}
