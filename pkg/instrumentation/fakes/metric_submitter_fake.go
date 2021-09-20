package fakes

import "github.com/Azure/Tivan-Libs/pkg/instrumentation"

type FakeMetricSubmitter struct {
	MetricValuePairs []MetricValuePair
}

type MetricValuePair struct {
	Metric instrumentation.Metric
	Value  int
}

func NewFakeMetricSubmitter() instrumentation.MetricSubmitter {
	return &FakeMetricSubmitter{
		MetricValuePairs: make([]MetricValuePair, 0),
	}
}

func (metricSubmitter *FakeMetricSubmitter) SendMetric(value int, metric instrumentation.Metric) {

	currentMetricValuePair := metricSubmitter.MetricValuePairs
	currentMetricValuePair = append(currentMetricValuePair, MetricValuePair{
		Metric: metric,
		Value:  value,
	})
	metricSubmitter.MetricValuePairs = currentMetricValuePair
}

func (metricSubmitter *FakeMetricSubmitter) SendMetricToNamespace(value int, metric instrumentation.Metric, accountName, namespaceName string) {

	currentMetricValuePair := metricSubmitter.MetricValuePairs
	currentMetricValuePair = append(currentMetricValuePair, MetricValuePair{
		Metric: metric,
		Value:  value,
	})
	metricSubmitter.MetricValuePairs = currentMetricValuePair
}
