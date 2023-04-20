package gbigquery

import (
	"context"
	"errors"
	"sync"

	"cloud.google.com/go/bigquery"
	"github.com/zpiroux/geist/entity"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

const sinkTypeId = "bigquery"

type Config struct {

	// ProjectId (required) specifies GCP project ID for this deployment.
	ProjectId string

	// Client (optional) enables fully customized clients to be used, e.g., for unit
	// testing.
	Client BigQueryClient

	// Creds (optional) can be used to override the default authentication method (using
	// GOOGLE_APPLICATION_CREDENTIALS) by providing externally created credentials.
	Creds *google.Credentials
}

// bigQueryMetadataMutex reduces the amount of unneeded requests for certain stream setup operations.
// If a stream is configured to operate with more than one concurrent instance (ops.streamsPerPod > 1),
// certain operations might be attempted by more than one of its stream entity instances (e.g. a stream's
// BQ loaders creating tables if requested in its spec).
// The mutex scope is per pod, but this is good enough in this case.
var bigQueryMetadataMutex sync.Mutex

type LoaderFactory struct {
	client         *bigquery.Client
	providedClient BigQueryClient
}

// NewLoaderFactory creates a new BigQuery loader connector.
func NewLoaderFactory(ctx context.Context, config Config) (*LoaderFactory, error) {
	var (
		err error
		lf  LoaderFactory
	)

	if config.ProjectId == "" {
		return nil, errors.New("no project id set")
	}

	if !isNil(config.Client) {
		lf.providedClient = config.Client
		return &lf, nil
	}

	if config.Creds == nil {
		lf.client, err = bigquery.NewClient(ctx, config.ProjectId)
	} else {
		lf.client, err = bigquery.NewClient(ctx, config.ProjectId, option.WithCredentials(config.Creds))
	}

	return &lf, err
}

func (lf *LoaderFactory) SinkId() string {
	return sinkTypeId
}

func (lf *LoaderFactory) NewLoader(ctx context.Context, c entity.Config) (entity.Loader, error) {
	var client BigQueryClient

	if isNil(lf.providedClient) {
		client = NewBigQueryClient(c, lf.client)
	} else {
		client = lf.providedClient
	}

	return newLoader(ctx, c, client, &bigQueryMetadataMutex)
}

func (lf *LoaderFactory) NewSinkExtractor(ctx context.Context, c entity.Config) (entity.Extractor, error) {
	return nil, nil
}

func (lf *LoaderFactory) Close(ctx context.Context) error {
	if lf.client != nil {
		if err := lf.client.Close(); err != nil {
			return err
		}
	}
	return nil
}
