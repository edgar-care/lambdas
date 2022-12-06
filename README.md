# edgar.care lambdas monorepository
This is the repository where all of the edgar.care api functions are implemented.

## Prerequisities
- aws_cli
- terraform
- go
- nodejs

## Installation
To install the dependencies use the following command :
```bash
make
```
To install the dependencies for a certain lambda use the following command :
```bash
make install t=mylambda
```

## Run
To run a particular lambda use the following command :
```bash
make run t=mylambda
```

## Deployment
Create the terraform environment using the following command :
```bash
make terraform
```

Deploy your modifications to AWS using the following command :
```bash
make deploy
```
