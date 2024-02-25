#============================================================
# DB
#============================================================
variable "DB_USERNAME" {
  type      = string
  sensitive = true
}
variable "DB_PASSWORD" {
  type      = string
  sensitive = true
}
variable "DB_NAME" {
  type      = string
  sensitive = true
}

#============================================================
# API
#============================================================

variable "PGDATABASE" {
  type      = string
  sensitive = true
}
variable "PGUSER" {
  type      = string
  sensitive = true
}
variable "PGPASSWORD" {
  type      = string
  sensitive = true
}
variable "JWT_SECRET" {
  type      = string
  sensitive = true
}
variable "GOOGLE_CLIENT_ID" {
  type      = string
  sensitive = true
}
variable "GOOGLE_CLIENT_SECRET" {
  type      = string
  sensitive = true
}
variable "RAKUTEN_APPLICATION_ID" {
  type      = string
  sensitive = true
}
variable "AWS_REGION" {
  type      = string
  sensitive = true
}
variable "AWS_ACCESS_KEY_ID" {
  type      = string
  sensitive = true
}
variable "AWS_SECRET_ACCESS_KEY" {
  type      = string
  sensitive = true
}
variable "S3_BUCKET_NAME" {
  type      = string
  sensitive = true
}

