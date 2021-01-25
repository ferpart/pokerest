# PokeRest

## About

PokeRest is a [GoLang](https://golang.org/) project that calls the [PokéApi](https://pokeapi.co/). It calls the api to make requests and get results in a json.

## Table of Contents

- [Quick start](#quick-start)
- [Documentation](#documentation)

## Quick Start

### Running locally

1. Using [GoLang](https://golang.org/).

   Follow the installation instructions for [GoLang](https://golang.org/) in the following address clicking for your operating system: [Download and Install Golang](https://golang.org/doc/install)

   After the instalation is done, copy the repository by downloading the zip or running:

   ```bash
   git clone https://github.com/ferpart/pokerest.git
   ```

   Go into the directory, and run the following command to run the [GoLang](https://golang.org/) http server:

   ```bash
   go run main.go
   ```

   As we are using go modules for this project, all the dependencie will be automatically installed. The http server will be hosted on:

   ```bash
   http://localhost:5000
   ```

2. Using [Docker](https://www.docker.com/)

   Follow the installation instructions for [Docker](https://www.docker.com/) in the following address clicking for your operating system: [Get Docker](https://docs.docker.com/get-docker/)

   After the instalation is done, copy the repository by downloading the zip or running:

   ```bash
   git clone https://github.com/ferpart/pokerest.git
   ```

   Go into the directory, and run the following command to build http server with [Docker](https://www.docker.com/):

   ```bash
   docker build -t <image-name> .
   ```

   Run the following command to start the http server in the docker container:

   ```bash
   docker run -p <desired-port>:5000 --name <container-name> -d <image-name>
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

- [main](https://github.com/ferpart/pokerest/tree/master/Docs/main.md)
- [handlers](https://github.com/ferpart/pokerest/tree/master/Docs/handlers.md)
- [helpers](https://github.com/ferpart/pokerest/tree/master/Docs/helpers.md)

The requests have to be called in the following manner:

- [requests](https://github.com/ferpart/pokerest/tree/master/Docs/requests.md)
