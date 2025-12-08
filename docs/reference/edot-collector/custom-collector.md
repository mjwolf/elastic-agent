---
navigation_title: Custom Collector
description: How to build a custom OpenTelemetry Collector distribution similar to EDOT.
applies_to:
  stack:
  serverless:
    observability:
  product:
    edot_collector: ga
products:
  - id: cloud-serverless
  - id: observability
  - id: edot-collector
---

# Build a custom EDOT-like collector

You can build and configure a [custom Collector](https://opentelemetry.io/docs/collector/custom-collector/) or extend the [OpenTelemetry Collector Contrib ](https://github.com/open-telemetry/opentelemetry-collector-contrib) distribution to collect logs and metrics and send them to Elastic Observability.

For a more seamless experience, use the Elastic Distribution of the OpenTelemetry Collector. Refer to the [configuration](/reference/edot-collector/config/index.md) docs for more information on configuring the EDOT Collector.

## Build a custom collector

To build a custom Collector to collect your telemetry data and send it to Elastic Observability, you need to:

1. Install the OpenTelemetry Collector builder, `ocb`.
1. Create a builder configuration file.
1. Build the Collector.

Refer to the following sections to complete these steps.

### Install the OpenTelemetry Collector builder

Install `ocb` using the command that aligns with your system from the [OpenTelemetry building a custom Collector documentation](https://opentelemetry.io/docs/collector/custom-collector/#step-1---install-the-builder).

:::{important}
Make sure to install the version of OpenTelemetry Collector Builder that matches the desired components' version.
:::

### Create a builder configuration file

Create a builder configuration file,`builder-config.yml`, to define the custom Collector. This file specifies the components, such as extensions, exporters, processors, receivers, and connectors, included in your custom Collector.

The following example, `builder-config.yml`, contains the components needed to send your telemetry data to Elastic Observability. For more information on these components, refer to the [components](/reference/edot-collector/components.md) documentation. Keep or remove components from the example configuration file to fit your needs.

% The following OCB configuration is automatically generated from the EDOT Collector source code.
% Automation is handled by /docs/scripts/update-docs/update-components-docs.py, which
% reads the go.mod file and then generates the OCB configuration.
% Note that while this runs on `main`, the OCB configuration is updated based on the latest released version.

% start:edot-collector-components-ocb
This OCB configuration is for EDOT Collector version 9.2.2.

```yaml
dist:
  otelcol_edot:
    description: "Elastic Distribution of OpenTelemetry Collector"
    output_path: ./dist/otelcol_edot
    builds:
      - name: otelcol_edot
        goos: [linux, darwin, windows]
        goarch: [amd64, arm64]
        output_path: ./dist/otelcol_edot_{{ .OS }}_{{ .Arch }}
        env:
          - CGO_ENABLED=0
          - GOOS={{ .OS }}
          - GOARCH={{ .Arch }}
        ldflags:
          - -s -w
          - -X go.opentelemetry.io/collector/otelcol.buildTimestamp={{ .BuildTimestamp }}
          - -X go.opentelemetry.io/collector/otelcol.version={{ .Version }}

receivers:

processors:

exporters:

connectors:

extensions:
  extensiontest:
    gomod: go.opentelemetry.io/collector/extension/extensiontest v0.139.0

providers:
```
% end:edot-collector-components-ocb

### Build the Collector

Build your custom Collector using the `ocb` tool and the configuration file by running the following command: `builder --config builder-config.yml`.

The command generates a new Collector in the specified output path, `otelcol-dev`. The generated Collector includes the components you specified in the configuration file.

For general information on building a custom Collector, refer to the [OpenTelemetry documentation](https://opentelemetry.io/docs/collector/custom-collector/#step-1---install-the-builder).

