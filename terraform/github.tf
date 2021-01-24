variable "DOCKERHUB_USERNAME" {}
variable "DOCKERHUB_TOKEN" {}
variable "AWS_ACCESS_KEY_ID" {}
variable "AWS_SECRET_KEY" {}
variable "AWS_DEFAULT_REGION" {}
variable "KUBE_CONFIG_DATA" {}

resource "github_actions_secret" "DOCKERHUB_USERNAME" {
    repository= "pokerest"
    secret_name = "DOCKERHUB_USERNAME"
    plaintext_value = var.DOCKERHUB_USERNAME
}

resource "github_actions_secret" "DOCKERHUB_TOKEN" {
    repository= "pokerest"
    secret_name = "DOCKERHUB_TOKEN"
    plaintext_value = var.DOCKERHUB_TOKEN
}

resource "github_actions_secret" "AWS_ACCESS_KEY_ID" {
    repository= "pokerest"
    secret_name = "AWS_ACCESS_KEY_ID"
    plaintext_value = var.AWS_ACCESS_KEY_ID
}

resource "github_actions_secret" "AWS_SECRET_KEY" {
    repository= "pokerest"
    secret_name = "AWS_SECRET_KEY"
    plaintext_value = var.AWS_SECRET_KEY
}

resource "github_actions_secret" "AWS_DEFAULT_REGION" {
    repository= "pokerest"
    secret_name = "AWS_DEFAULT_REGION"
    plaintext_value = var.AWS_DEFAULT_REGION
}

resource "github_actions_secret" "KUBE_CONFIG_DATA" {
    repository= "pokerest"
    secret_name = "KUBE_CONFIG_DATA"
    plaintext_value = var.KUBE_CONFIG_DATA
}