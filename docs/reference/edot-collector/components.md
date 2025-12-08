---
navigation_title: Components
description: List of components included in the EDOT Collector, categorized as Core or Extended.
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

# Components included in the EDOT Collector

The {{edot}} (EDOT) Collector includes embedded Collector components from the [OTel Collector Core](https://github.com/open-telemetry/opentelemetry-collector),
[OTel Collector Contrib](https://github.com/open-telemetry/opentelemetry-collector-contrib) and the [Elastic Collector Components](https://github.com/elastic/opentelemetry-collector-components) repositories.

The components included in the EDOT Collector are categorized into **[Core]** and **[Extended]** components. The following table describes the current components included in the EDOT Collector, their source, and support status.

% The following table is automatically generated from the EDOT Collector source code.
% Automation is handled by /docs/scripts/update-docs/update-components-docs.py, which
% reads the go.mod file, the components.yml file, and then generates the table.
% Note that while this runs on `main`, the table is updated based on the latest released version.

% start:edot-collector-components-table
## List of components

These components are included in EDOT Collector version 9.2.2.

| Component | GitHub Repo | Support status | Version |
|:---|:---|:---|:---|
|***Extensions***||||
| [extensiontest](https://github.com/open-telemetry/opentelemetry-collector/tree/main/extension/extensiontest) | [OTel Core Repo](https://github.com/open-telemetry/opentelemetry-collector) | [Extended] | v0.139.0 |

% end:edot-collector-components-table

[Extended]: opentelemetry://reference/compatibility/nomenclature.md#extended-components
[Core]: opentelemetry://reference/compatibility/nomenclature.md#core-components
