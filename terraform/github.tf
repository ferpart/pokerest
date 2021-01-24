variable "DOCKERHUB_USERNAME" {}
variable "DOCKERHUB_TOKEN" {}

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