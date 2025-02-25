package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

type TransformedAnalise struct {
	Entity      string `json:"entity"`
	Type        string `json:"type"`
	Sentimental string `json:"sentimental"`
}

func Handler(ctx context.Context, event events.KinesisFirehoseEvent) (events.KinesisFirehoseResponse, error) {
	var transformedRecords []events.KinesisFirehoseResponseRecord

	for _, record := range event.Records {
		var analyseResult comprehend.DetectTargetedSentimentOutput

		err := json.Unmarshal(record.Data, &analyseResult)
		if err != nil {
			log.Printf("Error decoding JSON: %v", err)
			transformedRecords = append(transformedRecords, events.KinesisFirehoseResponseRecord{
				RecordID: record.RecordID,
				Result:   events.KinesisFirehoseTransformedStateDropped,
				Data:     []byte{},
			})
			continue
		}

		var sb strings.Builder

		for _, e := range analyseResult.Entities {
			for _, m := range e.Mentions {
				if m.Text != nil && m.Type != nil && m.MentionSentiment.Sentiment != nil {
					analise := TransformedAnalise{
						Entity:      *m.Text,
						Type:        *m.Type,
						Sentimental: *m.MentionSentiment.Sentiment,
					}

					analiseJSON, err := json.Marshal(analise)
					if err != nil {
						log.Printf("Error encoding transformed data: %v", err)
						continue
					}

					sb.Write(analiseJSON)
					sb.WriteString("\n")
				}
			}
		}

		transformedData := []byte(sb.String())

		transformedRecords = append(transformedRecords, events.KinesisFirehoseResponseRecord{
			RecordID: record.RecordID,
			Result:   events.KinesisFirehoseTransformedStateOk,
			Data:     transformedData,
		})
	}

	return events.KinesisFirehoseResponse{Records: transformedRecords}, nil
}

func main() {
	lambda.Start(Handler)
}
