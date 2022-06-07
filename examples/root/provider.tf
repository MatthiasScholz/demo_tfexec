provider "aws" {
  region = "us-east-1"
  # FIXME avoid hard coded profile names, inject it via go, or generate this file.
  profile = "tw-beach-push"
}
