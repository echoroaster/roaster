package logging // import "github.com/echoroaster/roaster/pkg/logging"

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/sirupsen/logrus"
)

var _ logrus.Hook = &cloudWatchHook{}

func init() {
	if os.Getenv("AWS_CLOUDWATCHLOGS_GROUP_NAME") != "" {
		hook, err := newCloudWatchHook()
		if err != nil {
			RootLogger.WithFields(logrus.Fields{
				"component": "logger",
				"driver":    "cloudwatch",
			}).Error(err)
		} else {
			RootLogger.AddHook(hook)
			RootLogger.Info("Logger add cloudwatch destination")
		}
	}
}

func newCloudWatchHook() (*cloudWatchHook, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	return &cloudWatchHook{
		Service:       cloudwatchlogs.New(sess),
		LogGroupName:  os.Getenv("AWS_CLOUDWATCHLOGS_GROUP_NAME"),
		LogStreamName: os.Getenv("AWS_CLOUDWATCHLOGS_STREAM_NAME"),
	}, nil
}

type cloudWatchHook struct {
	Service       *cloudwatchlogs.CloudWatchLogs
	LogGroupName  string
	LogStreamName string
	sequenceToken *string
	m             sync.Mutex
}

func (h *cloudWatchHook) Init(ctx context.Context) (err error) {
	_, err = h.Service.CreateLogStreamWithContext(ctx, &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  &h.LogGroupName,
		LogStreamName: &h.LogStreamName,
	})
	return
}

func (h *cloudWatchHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (h *cloudWatchHook) Fire(entry *logrus.Entry) (err error) {
	type logMessage struct {
		Message string                 `json:"message"`
		Level   string                 `json:"level"`
		Fields  map[string]interface{} `json:"fields"`
		Time    string                 `json:"time"`
	}

	msg, err := json.Marshal(&logMessage{
		Message: entry.Message,
		Level:   entry.Level.String(),
		Fields:  entry.Data,
		Time:    entry.Time.Format(time.RFC3339),
	})
	if err != nil {
		return err
	}

	event := &cloudwatchlogs.InputLogEvent{
		Timestamp: aws.Int64(entry.Time.UnixNano() / int64(time.Millisecond)),
		Message:   aws.String(string(msg)),
	}
	resp, err := h.Service.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
		LogEvents: []*cloudwatchlogs.InputLogEvent{
			event,
		},
		LogGroupName:  &h.LogGroupName,
		LogStreamName: &h.LogStreamName,
		SequenceToken: h.sequenceToken,
	})
	if err != nil {
		return
	}

	h.sequenceToken = resp.NextSequenceToken

	return
}
