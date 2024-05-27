provider "aws" {
    access_key = "test"
    secret_key = "test"
    region     = "eu-central-1"
    
    skip_credentials_validation = true
    skip_metadata_api_check     = true
    skip_requesting_account_id  = true

    endpoints {
        lambda     = "http://localhost:4566"
        iam        = "http://localhost:4566"
        s3         = "http://s3.localhost.localstack.cloud:4566"
    }
}

resource "aws_s3_bucket" "s3_start" {
  bucket = "s3-start"
}

resource "aws_s3_bucket" "s3_finish" {
  bucket = "s3-finish"
}


data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "iam_for_lambda" {
  name               = "iam_for_lambda"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "archive_file" "lambda" {
  type        = "zip"
  source_file = "lambda_function.py"
  output_path = "lambda_function_payload.zip"
}

resource "aws_lambda_function" "test_lambda" {
  filename      = "lambda_function_payload.zip"
  function_name = "lambda_function_name"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "lambda_function.lambda_handler"
  runtime = "python3.12"
}

resource "aws_s3_bucket_lifecycle_configuration" "s3_start_lifecycle" {
  bucket = aws_s3_bucket.s3_start.id

  rule {
    id      = "move_to_s3_finish"
    status = "Enabled"
    prefix = ""

    transition {
      days          = 1
      storage_class = "STANDARD_IA"
    }

    expiration {
      days = 7
    }
  }
  
}
resource "aws_s3_bucket_notification" "s3_start_notification" {
  bucket = aws_s3_bucket.s3_start.id

  lambda_function {
    lambda_function_arn = aws_lambda_function.test_lambda.arn
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = ""
    filter_suffix       = ""
  }
}