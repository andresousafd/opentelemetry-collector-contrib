// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxspan // import "github.com/andresousafd/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxspan"

import "go.opentelemetry.io/collector/pdata/ptrace"

const (
	Name   = "span"
	DocRef = "https://github.com/andresousafd/opentelemetry-collector-contrib/tree/main/pkg/ottl/contexts/ottlspan"
)

type Context interface {
	GetSpan() ptrace.Span
}
