resource "aws_lambda_function" "ginventory_api_lambda_function" {
  filename      = data.archive_file.dummy.output_path
  function_name = "ginventory-api"
  handler       = "ginventory-api"
  role          = aws_iam_role.ginventory_api_iam_role.arn
  runtime       = "go1.x"
}
