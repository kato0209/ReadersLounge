#############################################################
# SSMパラメータストア
#############################################################

#============================================================
# Common
#============================================================


#============================================================
# DB
#============================================================
variable "DB_USERNAME" {}
variable "DB_PASSWORD" {}

resource "aws_ssm_parameter" "db_username" {
  name        = "db-username"
  value       = var.DB_USERNAME
  type        = "SecureString"
  description = "DB_USERNAME"
}
resource "aws_ssm_parameter" "db_password" {
  name        = "db-password"
  value       = var.DB_PASSWORD
  type        = "SecureString"
  description = "DB_PASSWORD"
}

#============================================================
# API
#============================================================
variable "JWT_SECRET" {}
variable "GOOGLE_CLIENT_ID" {}
variable "GOOGLE_CLIENT_SECRET" {}
variable "RAKUTEN_APPLICATION_ID" {}

resource "aws_ssm_parameter" "jwt_secret" {
  name        = "jwt-secret"
  value       = var.JWT_SECRET
  type        = "SecureString"
  description = "JWT_SECRET"
}
resource "aws_ssm_parameter" "google_client_id" {
  name        = "google-client-id"
  value       = var.GOOGLE_CLIENT_ID
  type        = "SecureString"
  description = "GOOGLE_CLIENT_ID"
}
resource "aws_ssm_parameter" "google_client_secret" {
  name        = "google-client-secret"
  value       = var.GOOGLE_CLIENT_SECRET
  type        = "SecureString"
  description = "GOOGLE_CLIENT_SECRET"
}
resource "aws_ssm_parameter" "rakuten_application_id" {
  name        = "rakuten-application-id"
  value       = var.RAKUTEN_APPLICATION_ID
  type        = "SecureString"
  description = "RAKUTEN_APPLICATION_ID"
}

#============================================================
# FRONT
#============================================================
variable "VITE_GOOGLE_OAUTH_CLIENT_ID" {}
variable "VITE_GOOGLE_CLIENT_SECRET" {}

resource "aws_ssm_parameter" "vite_google_oauth_client_id" {
  name        = "vite-google-oauth-client-id"
  value       = var.VITE_GOOGLE_OAUTH_CLIENT_ID
  type        = "SecureString"
  description = "VITE_GOOGLE_OAUTH_CLIENT_ID"
}
resource "aws_ssm_parameter" "vite_google_client_secret" {
  name        = "vite-google-client-secret"
  value       = var.VITE_GOOGLE_CLIENT_SECRET
  type        = "SecureString"
  description = "VITE_GOOGLE_CLIENT_SECRET"
}
