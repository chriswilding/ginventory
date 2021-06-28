data "aws_iam_policy_document" "ginventory_api_iam_policy_document" {
  statement {
    actions = [
      "sts:AssumeRole"
    ]

    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ginventory_api_iam_role" {
  name               = "ginventory-api-role"
  assume_role_policy = data.aws_iam_policy_document.ginventory_api_iam_policy_document.json
}

resource "aws_iam_role_policy_attachment" "ginventory_api_iam_role_policy_attachment" {
  role       = aws_iam_role.ginventory_api_iam_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

data "archive_file" "dummy" {
  type        = "zip"
  output_path = "${path.module}/ginventory-api.zip"

  source {
    content  = "dummy"
    filename = "dummy.txt"
  }
}
