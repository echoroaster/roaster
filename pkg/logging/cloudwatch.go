// Copyright 2020 EchoRoaster
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package logging // import "github.com/echoroaster/roaster/pkg/logging"

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func newCloudWatchLoggerHook() (logHook, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	return &cloudWatchLoggerHook{
		Level:         Info,
		Service:       cloudwatchlogs.New(sess),
		LogGroupName:  os.Getenv("AWS_CLOUDWATCHLOGS_GROUP_NAME"),
		LogStreamName: os.Getenv("AWS_CLOUDWATCHLOGS_STREAM_NAME"),
	}, nil
}

type cloudWatchLoggerHook struct {
	Level         Level
	Service       *cloudwatchlogs.CloudWatchLogs
	LogGroupName  string
	LogStreamName string
	sequenceToken *string
	m             sync.Mutex
}

func (h *cloudWatchLoggerHook) Log(entry *logEntry) (err error) {
	if entry.Level < h.Level {
		return nil
	}

	type logMessage struct {
		Message string   `json:"message"`
		Level   string   `json:"level"`
		Source  logFrame `json:"source"`
	}

	msg, err := json.Marshal(&logMessage{
		Message: entry.Message,
		Level:   entry.Level.String(),
		Source:  entry.Frame,
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
