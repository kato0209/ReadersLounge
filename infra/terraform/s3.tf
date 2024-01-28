##########
#   S3   #
##########
resource "aws_s3_bucket" "images" {
  bucket = "readerslounge-s3-bucket-for-images"

  tags = {
    Name = "readerslounge-s3-bucket-for-images"
  }
}

resource "aws_s3_bucket_cors_configuration" "images" {
  bucket = aws_s3_bucket.images.id

  cors_rule {
    allowed_origins = ["*"]
    allowed_methods = ["GET"]
    allowed_headers = ["*"]
  }
}

resource "aws_s3_bucket_acl" "images" {
  bucket = aws_s3_bucket.images.id
  acl    = "public-read"
}
