# SAM-AWSComprehend-Pipeline

<p align="center">
  <img src="./infra.svg" alt="Required User Permissions" />
</p>

## Overview

This project leverages AWS Serverless Application Model (SAM) to build an automated sentiment analysis pipeline using AWS Comprehend. It enables seamless text analysis to determine sentiment polarity (positive, negative, neutral, or mixed).

For a detailed explanation of the implementation, visit my blog: [Sentiment Analysis Pipeline](https://jfolgado.com/posts/sentimentalanalises/).

## Features
- **Serverless Deployment**: Utilizes AWS SAM for easy deployment and scalability.

- **AWS Comprehend Integration**: Analyzes text sentiment with high accuracy.

- **Scalable and Cost-Effective**: Pay-as-you-go pricing with AWS Lambda.

- **Automated Processing**: Handles real-time or batch sentiment analysis.


## Prerequisites

Ensure you have the following installed:

- [AWS CLI](https://aws.amazon.com/cli/)

- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)

- Go 2.22

## Deployment Steps

1. Clone the repository
```
git clone https://github.com/your-repo/sam-awscomprehend-pipeline.git
cd sam-awscomprehend-pipeline
```
2. Build the project
```
sam build -u
```

3. Deploy the Application:
```
sam deploy -g
```
Follow the prompts to configure AWS credentials and stack settings.

## Usage

Once deployed, the pipeline will process input text and return sentiment analysis results via AWS Lambda and AWS Comprehend. You can invoke the Lambda function manually or integrate it with an API Gateway for real-time analysis.

## Architecture
This project follows a serverless architecture using:

- **API Gatway**

- **AWS Lambda** for processing text data.

- **AWS Comprehend** for sentiment analysis.

- **AWS Firehose** for real time streaming.

- **Amazon S3** for storing text data.

## Contributing
Contributions are welcome! Feel free to fork the repository and submit a pull request