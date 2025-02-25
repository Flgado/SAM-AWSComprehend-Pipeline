package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

const DELIVERYSTREAMNAME = "kdf-firehose-78872folgado"

type commentToAnalyse struct {
	Comment string `json:"comment"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var comment commentToAnalyse
	err := json.Unmarshal([]byte(request.Body), &comment)
	if err != nil {
		log.Printf("Failed to decode JSON: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, nil
	}

	mySession := session.Must(session.NewSession())
	comprehendClient := comprehend.New(mySession)

	comprehendPayload := comprehend.DetectTargetedSentimentInput{
		LanguageCode: aws.String("en"),
		Text:         &comment.Comment,
	}

	analyseResult, err := comprehendClient.DetectTargetedSentiment(&comprehendPayload)

	if err != nil {
		log.Printf("Failed to get analise from aws comprehhend: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Invalid",
		}, nil
	}

	analyseResultJSON, err := json.Marshal(analyseResult)
	if err != nil {
		log.Printf("Failed to marshal Comprehend result: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to process sentiment analysis",
		}, nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("Failed to load AWS config: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to load AWS configuration",
		}, nil
	}

	firehoseClient := firehose.NewFromConfig(cfg)

	firehosePayload := firehose.PutRecordInput{
		DeliveryStreamName: aws.String(DELIVERYSTREAMNAME),
		Record: &types.Record{
			Data: analyseResultJSON,
		},
	}

	_, err = firehoseClient.PutRecord(context.TODO(), &firehosePayload)

	if err != nil {
		log.Printf("Failed to put record in Firehouse: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to put record in Firehose",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       analyseResult.String(),
	}, nil
}

func main() {
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return Handler(request)
	})
}
