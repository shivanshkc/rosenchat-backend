package logger

import (
	"context"
	"fmt"
	"rosenchat/src/configs"
	"rosenchat/src/utils/ctxutils"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/option"

	"cloud.google.com/go/logging"
)

var conf = configs.Get()

// implGCP implements ILogger using Google's cloud logging.
type implGCP struct {
	client *logging.Client
}

func (i *implGCP) Debugf(ctx context.Context, format string, a ...interface{}) {
	i.Logf(ctx, logging.Debug, format, a...)
}

func (i *implGCP) Infof(ctx context.Context, format string, a ...interface{}) {
	i.Logf(ctx, logging.Info, format, a...)
}

func (i *implGCP) Warnf(ctx context.Context, format string, a ...interface{}) {
	i.Logf(ctx, logging.Warning, format, a...)
}

func (i *implGCP) Errorf(ctx context.Context, format string, a ...interface{}) {
	i.Logf(ctx, logging.Error, format, a...)
}

func (i *implGCP) Close(ctx context.Context) {
	// If cloud logging is not enabled, we don't attempt disconnection.
	if !conf.GoogleCloudLogger.CloudEnabled || i.client == nil {
		return
	}

	if err := i.client.Close(); err != nil {
		fmt.Printf("Failed to close the GCP logger: %+v\n", err)
	}
}

func (i *implGCP) Logf(ctx context.Context, sev logging.Severity, format string, a ...interface{}) {
	// Getting additional logging info (package, file etc)
	packageInfo, file, line := getCallerDetails(3)

	// Logging locally on stdout.
	localFormat := fmt.Sprintf("%s %d %s: %s\n", file, line, packageInfo, format)
	fmt.Printf(localFormat, a...)

	// If cloud logging is not enabled, we don't go further.
	if !conf.GoogleCloudLogger.CloudEnabled {
		return
	}

	// Getting additional logging info (request/log ID)
	reqCtx := ctxutils.GetRequestInfo(ctx)
	logID := uuid.NewString()
	if reqCtx != nil {
		logID = reqCtx.ID
	}

	// Logging key-value pairs.
	labels := map[string]string{
		"package": packageInfo,
		"file":    file,
		"line":    fmt.Sprintf("%d", line),
		"logID":   logID,
	}

	// Logging on the cloud.
	logger := i.client.Logger(conf.Application.Name)
	logger.Log(logging.Entry{
		Timestamp: time.Now(),
		Severity:  sev,
		Payload:   fmt.Sprintf(format, a...),
		Labels:    labels,
		InsertID:  uuid.NewString(),
	})
}

func (i *implGCP) init() {
	// If cloud logging is not enabled, we don't connect the client.
	if !conf.GoogleCloudLogger.CloudEnabled {
		return
	}

	credFileOption := option.WithCredentialsFile(conf.GoogleCloudLogger.KeyFile)
	client, err := logging.NewClient(context.Background(), conf.GoogleCloudLogger.ProjectID, credFileOption)
	if err != nil {
		panic("failed to obtain GCP logging client: " + err.Error())
	}

	i.client = client
}
