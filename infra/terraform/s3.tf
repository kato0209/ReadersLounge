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

resource "aws_s3_bucket_public_access_block" "images" {
  bucket                  = aws_s3_bucket.images.id
  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

resource "aws_s3_bucket_policy" "images" {
  bucket = aws_s3_bucket.images.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::620958051842:user/terraform"
        }
        Action = [
          "s3:GetObject",
          "s3:PutObject",
          "s3:DeleteObject",
          "s3:PutBucketAcl"
        ]
        Resource = [
          "${aws_s3_bucket.images.arn}/*",
          "${aws_s3_bucket.images.arn}"
        ]
      },
    ]
  })
}
