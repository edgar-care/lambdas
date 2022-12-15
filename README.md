# edgar.care lambdas monorepository
This is the repository where all of the edgar.care api functions are implemented.

## Prerequisities
- [aws_cli](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- [terraform](https://developer.hashicorp.com/terraform/downloads)
- [go 1.19.4 (with gvm)](https://github.com/moovweb/gvm)
- [python 3.9.16 (with pyenv)](https://github.com/pyenv/pyenv)

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
make start t=mylambda
```

## Deployment
Deploy a certain lambda using the following command :
```bash
make deploy t=mylambda
```
