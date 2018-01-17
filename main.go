package main

import (
    "fmt"
    "os"
    "github.com/aws/aws-sdk-go/service/lambda"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws"
)


func main() {
    // get function region and set default of us-east-1
    region := os.Getenv("PLUGIN_FUNCTION_REGION")

    if (region == "") {
      region = "us-east-1"
    }

    svc := lambda.New(session.New(&aws.Config{
        Region: aws.String(region),
    }))

    input := &lambda.UpdateFunctionCodeInput{
        FunctionName:    aws.String(os.Getenv("PLUGIN_FUNCTION_NAME")),
        Publish:         aws.Bool(true),
        S3Bucket:        aws.String(os.Getenv("PLUGIN_S3_BUCKET")),
        S3Key:           aws.String(os.Getenv("PLUGIN_FILE_NAME")),
    }

    result, err := svc.UpdateFunctionCode(input)
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
                case lambda.ErrCodeServiceException:
                    fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
                case lambda.ErrCodeResourceNotFoundException:
                    fmt.Println(lambda.ErrCodeResourceNotFoundException, aerr.Error())
                case lambda.ErrCodeInvalidParameterValueException:
                    fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
                case lambda.ErrCodeTooManyRequestsException:
                    fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
                case lambda.ErrCodeCodeStorageExceededException:
                    fmt.Println(lambda.ErrCodeCodeStorageExceededException, aerr.Error())
                default:
                    fmt.Println(aerr.Error())
            }
        } else {
            // Print the error, cast err to awserr.Error to get the Code and
            // Message from an error.
            fmt.Println(err.Error())
        }
        return
    }

    fmt.Println(result)
}
