// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxdatapoint // import "github.com/andresousafd/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxdatapoint"

const (
	Name   = "datapoint"
	DocRef = "https://github.com/andresousafd/opentelemetry-collector-contrib/tree/main/pkg/ottl/contexts/ottldatapoint"
)

type Context interface {
	GetDataPoint() any
}
