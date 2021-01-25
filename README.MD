# PokeRest

<p align="center">
  <img src="https://raw.githubusercontent.com/PokeAPI/media/master/logo/pokeapi_256.png" />
  <img src="https://golang.org/lib/godoc/images/go-logo-blue.svg" />
</p>

## About

PokeRest is a [GoLang](https://golang.org/) project that calls the [PokéApi](https://pokeapi.co/). It calls the api to make requests and get results in a json.

## Table of Contents

- [Quick start](#quick-start)
- [Documentation](#documentation)
- [Testing](#testing)
- [Deployment](#deployment)
- [Continuous Integration](#continuous-integration)

## Quick Start

### Running locally

1. Using [GoLang](https://golang.org/).

   Follow the installation instructions for [GoLang](https://golang.org/) in the following address clicking for your operating system: [Download and Install Golang](https://golang.org/doc/install)

   After the instalation is done, copy the repository by downloading the zip or running:

   ```bash
   $ git clone https://github.com/ferpart/pokerest.git
   ```

   Go into the directory, and run the following command to run the [GoLang](https://golang.org/) http server:

   ```bash
   $ go run main.go
   ```

   As we are using go modules for this project, all the dependencie will be automatically installed. The http server will be hosted on:

   ```bash
   http://localhost:5000
   ```

2. Using [Docker](https://www.docker.com/)

   Follow the installation instructions for [Docker](https://www.docker.com/) in the following address clicking for your operating system: [Get Docker](https://docs.docker.com/get-docker/)

   After the instalation is done, copy the repository by downloading the zip or running:

   ```bash
   $ git clone https://github.com/ferpart/pokerest.git
   ```

   Go into the directory, and run the following command to build http server with [Docker](https://www.docker.com/):

   ```bash
   $ docker build -t <image-name> .
   ```

   Run the following command to start the http server in the docker container:

   ```bash
   $ docker run -p <desired-port>:5000 --name <container-name> -d <image-name>
   ```

   The http server will be found in the port given in the selected port:

   ```bash
   http://localhost:<desired-port>
   ```

## Documentation

PokeRest is a [GoLang](https://golang.org/) project that calls the [PokéApi](https://pokeapi.co/). It calls the api to make requests and get results in a json.

Pokerest uses the [PokéApi](https://pokeapi.co/) for doing two things depending on the api called.

1. Given two pokemon we return if the first has an advantage over the second using its type as the determinant.
2. Given two or more pokemon, we return the moves they have in common. A limit of how many moves, as well as a language can be provided for this endpoint.

The program has a total of 3 packages each documented bellow:

- [main](https://github.com/ferpart/pokerest/tree/master/docs/main.md)
- [handlers](https://github.com/ferpart/pokerest/tree/master/docs/handlers.md)
- [helpers](https://github.com/ferpart/pokerest/tree/master/docs/helpers.md)

The requests have to be called in the following manner:

- [requests](https://github.com/ferpart/pokerest/tree/master/docs/requests.md)

## Testing

To test if all the packages work correctly, run the following command. (You have to have golang installed, detailed in the [quick start section](#quick-start)).

First enter the root directory, and then run the following code:

```bash
$ go test ./...
```

This will run all the tests saved in the repository.

## Deployment

This api has been dockerized and uploaded as an image to the following docker repository:

- [ferpart/pokemon-rest-go](https://hub.docker.com/repository/docker/ferpart/pokemon-rest-go)

This api is also deployed in a [AWS Elastic Kubernetes Service](https://aws.amazon.com/eks/) using [Kubernetes](https://kubernetes.io/) as the container orchestrator.

The deployment to EKS was made easy by using the [eksctl](https://eksctl.io/). This tool automates the creation of clusters on EKS. The configuration used for this project is stored in the following file:

- [cluster.yaml](https://github.com/ferpart/pokerest/blob/master/k8s/eksctl/cluster.yaml)

And launched using the following command:

```bash
$ eksctl create cluster -f <file>
```

The Kubernates deployment was created using the following file:

- [pokemon-rest.yaml](https://github.com/ferpart/pokerest/blob/master/k8s/pokemon-rest.yaml)

And added to the cluster using the command:

```bash
$ kubectl apply -f <file>
```

## Continuous Integration

The continuous integration is separated in 2 parts:

1. [Initialization](#initialization)
2. [Testing, Building, and Deploying](#testing-building-and-deploying)

### Initialization

[Terraform](https://www.terraform.io/) was used to deploy all the github secrets used for automation with [Github Actions](https://github.com/features/actions). The terraform files are stored in the following directory:

- [terraform](https://github.com/ferpart/pokerest/tree/master/terraform)

First we need to [install terraform](https://www.terraform.io/downloads.html)

Then change the [.example_env](https://github.com/ferpart/pokerest/blob/master/.example_env) file and adding the following values in their respective slots:

- Github Access Token
- DockerHub Username
- DockerHub Token
- AWS Access Key ID
- AWS Access Key Secret
- AWS Region
- Kube Config File (base 64 encoded)

Then exporting the file to the environment variables to be used by Terraform. And finally running the following commands to apply the tokens to github secrets:

```bash
$ terraform init
$ terraform apply --auto-approve
```

### Testing, Building, and Deploying

The testing, building, and deploying is all done with one workflow located in the following file:

- [test-deploy.yml](https://github.com/ferpart/pokerest/blob/master/.github/workflows/test-deploy.yml)

When a change is pushed to the repository, the workflow will be executed, and the following steps will be run sequentially:

1. `test`: This uses the [Golang Action](https://github.com/marketplace/actions/golang-action) action, to run all the package tests and verify that the code works in an isolated environment.
2. `build`: This uses the [Build and push Docker images](https://github.com/marketplace/actions/build-and-push-docker-images) action, to get the directories docker file, and create a new built to upload ro dockerhub. This action uses two of the github secrets created by terrafotm (see [Initialization](#initialization)). The secrets used for this step are the:
   - DOCKERHUB_USERNAME
   - DOCKERHUB_TOKEN
3. `deploy`: This uses the [kubectl-aws-eks](https://github.com/marketplace/actions/kubectl-aws-eks) action, to rune a `kubectl` command to update de image running in our [EKS](https://aws.amazon.com/eks/) cluster. it uses the following command to apply the changes that were built and pushed in the previous step:

   - `kubectl set image deployments/<deployment-name> <app-name>=docker.io/<repository>:<new-tag>`

   This command performs a [rolling update](https://kubernetes.io/docs/tutorials/kubernetes-basics/update/update-intro/) so that there is no downtime on our api. This action uses the remaining 4 secrets in our repository, which are:

   - AWS_ACCESS_KEY_ID
   - AWS_SECRET_KEY
   - AWS_DEFAULT_REGION
   - KUBE_CONFIG_DATA

These 3 steps allow us to automatically test, build, and deploy in less than 5 minutes (given all the actions are successful)
