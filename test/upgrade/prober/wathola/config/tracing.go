/*
Copyright 2022 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"fmt"

	"go.uber.org/zap"
	"knative.dev/pkg/tracing"
	tracingconfig "knative.dev/pkg/tracing/config"
)

var Tracer *tracing.OpenCensusTracer

func SetupTracing() {
	config, err := tracingconfig.JSONToTracingConfig(Instance.TracingConfig)
	if err != nil {
		Log.Warn("Tracing configuration is invalid, using the no-op default", zap.Error(err))
	}
	if Tracer, err = SetupStaticPublishing(Log, "", config); err != nil {
		Log.Fatal("Error setting up trace publishing", zap.Error(err))
	}
}

func SetupStaticPublishing(logger *zap.SugaredLogger, serviceName string, cfg *tracingconfig.Config) (*tracing.OpenCensusTracer, error) {
	oct := tracing.NewOpenCensusTracer(tracing.WithExporter(serviceName, logger))
	if err := oct.ApplyConfig(cfg); err != nil {
		return nil, fmt.Errorf("unable to set OpenCensusTracing config: %w", err)
	}
	return oct, nil
}
