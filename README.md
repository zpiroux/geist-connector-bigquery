# GEIST BigQuery Connector
<div>

[![Go Report Card](https://goreportcard.com/badge/github.com/zpiroux/geist-connector-bigquery)](https://goreportcard.com/report/github.com/zpiroux/geist-connector-bigquery)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=zpiroux_geist-connector-bigquery&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=zpiroux_geist-connector-bigquery)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=zpiroux_geist-connector-bigquery&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=zpiroux_geist-connector-bigquery)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=zpiroux_geist-connector-bigquery&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=zpiroux_geist-connector-bigquery)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=zpiroux_geist-connector-bigquery&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=zpiroux_geist-connector-bigquery)

</div>

Geist BigQuery Connector enables BigQuery as a sink type in stream specs when using Geist.

Note that this connector previously resided in the [geist-connector-gcp](https://github.com/zpiroux/geist-connector-gcp) repo, but was moved out to its own for improved maintainability and ease of use.

## Usage
See [GEIST core repo](https://github.com/zpiroux/geist) for general information.

Install with:
```sh
go get github.com/zpiroux/geist-connector-bigquery
```

### Geist Integration

Register connector prior to starting up Geist with (error handling omitted):
```go
import (
	"github.com/zpiroux/geist"
	"github.com/zpiroux/geist-connector-bigquery/gbigquery"
)

...
geistConfig := geist.NewConfig()

bqConfig := &gbigquery.Config{ /* add config */ }

err = geistConfig.RegisterExtractorType(gbigquery.NewExtractorFactory(bqConfig))
err = geistConfig.RegisterLoaderType(gbigquery.NewLoaderFactory(bqConfig))

g, err := geist.New(ctx, geistConfig)
...
```

### Stream Spec Configuration

For now, see stream spec examples [here](https://github.com/zpiroux/geist-connector-bigquery/gbigquery/test/specs/kafkasrc-bigquerysink-fooevents.json) and [here](https://github.com/zpiroux/geist-connector-bigquery/gbigquery/gbigquery_test.go).

## Contact
info @ zpiroux . com

## License
Geist BigQuery Connector source code is available under the MIT License.
