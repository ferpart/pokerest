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

Coming Soon.
