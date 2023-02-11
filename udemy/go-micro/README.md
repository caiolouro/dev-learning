# Working with Microservices in Go (Golang)

Files to code along the [Working with Microservices in Go (Golang)](https://www.udemy.com/course/working-with-microservices-in-go/) course on Udemy.

# Running locally

## Prerequisites

1. Install Docker and Go version 1.18 [directly](https://go.dev/doc/install) or by using a version manager like [asdf](https://asdf-vm.com/) (the version config file is under source control);
2. Create a `db-data` folder and, inside it, a `postgres` folder.

## Starting the services

Run `make fresh-start-back` and then, on another tab, run `make start-front`.

## Finally

Open the http://localhost URL in a web browser.