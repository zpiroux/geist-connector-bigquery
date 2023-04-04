# GEIST BigQuery Connector

Geist BigQuery Connector enables BigQuery as a sink type in stream specs when using Geist.
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
