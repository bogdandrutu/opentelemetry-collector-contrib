// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter/internal/metadata"

	"go.opentelemetry.io/collector/component/componenttest"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := componenttest.NewTelemetry()
	tb, err := metadata.NewTelemetryBuilder(testTel.NewTelemetrySettings())
	require.NoError(t, err)
	defer tb.Shutdown()
	tb.ExporterPrometheusremotewriteConsumers.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteFailedTranslations.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteSentBatches.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteTranslatedTimeSeries.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWalBytesRead.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWalBytesWritten.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWalLag.Record(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWalReadLatency.Record(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWalReads.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWalReadsFailures.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWalWriteLatency.Record(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWalWrites.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWalWritesFailures.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWrittenExemplars.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWrittenHistograms.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteWrittenSamples.Add(context.Background(), 1)
	AssertEqualExporterPrometheusremotewriteConsumers(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteFailedTranslations(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteSentBatches(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteTranslatedTimeSeries(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWalBytesRead(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWalBytesWritten(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWalLag(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWalReadLatency(t, testTel,
		[]metricdata.HistogramDataPoint[int64]{{}}, metricdatatest.IgnoreValue(),
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWalReads(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWalReadsFailures(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWalWriteLatency(t, testTel,
		[]metricdata.HistogramDataPoint[int64]{{}}, metricdatatest.IgnoreValue(),
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWalWrites(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWalWritesFailures(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWrittenExemplars(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWrittenHistograms(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualExporterPrometheusremotewriteWrittenSamples(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())

	require.NoError(t, testTel.Shutdown(context.Background()))
}
