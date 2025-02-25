build-ProcessCommentAnalysis:
	GOOS=linux CGO_ENABLE=0 go build -o lambda/ProcessCommentAnalysis/bootstrap lambda/ProcessCommentAnalysis/main.go
	cp lambda/ProcessCommentAnalysis/bootstrap $(ARTIFACTS_DIR)/.

build-TransformationFunction:
	GOOS=linux CGO_ENABLE=0 go build -o lambda/TransformationFunction/bootstrap lambda/TransformationFunction/main.go
	cp lambda/TransformationFunction/bootstrap $(ARTIFACTS_DIR)/.