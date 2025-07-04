// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlqueryreceiver // import "github.com/andresousafd/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"

import (
	"database/sql"

	"go.opentelemetry.io/collector/receiver"

	"github.com/andresousafd/opentelemetry-collector-contrib/internal/sqlquery"
	"github.com/andresousafd/opentelemetry-collector-contrib/receiver/sqlqueryreceiver/internal/metadata"
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithLogs(createLogsReceiverFunc(sql.Open, sqlquery.NewDbClient), metadata.LogsStability),
		receiver.WithMetrics(createMetricsReceiverFunc(sql.Open, sqlquery.NewDbClient), metadata.MetricsStability),
	)
}
