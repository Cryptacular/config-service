# Go Configuration Service

[![Build Status](https://travis-ci.org/Cryptacular/config-service.svg?branch=master)](https://travis-ci.org/Cryptacular/config-service)

## Overview

A website with API that allows the user to generate a configuration file based on various input parameters. It will be hosted on AWS.

Possible use case: A/B testing.

## Features

- Go
- AWS
- GraphQL
- Front end with settings
  - Environment (dev, prod, uat)
  - ...
- API
- Database
- Unit tests

## Workflow

The consumer provides the following information:

- market
- language (locale)
- user ID or username
- IP address

If not all are provided it should still work.

The API will then return JSON with the following structure:

```json
{
  "configuration": {},
  "abtests": {}
}
```

Where `configuration` will contain information like country configuration, company configuration etc.

`abtests` will contain information about current A/B tests and which group the user falls into.

## To Do

- [x] Host on AWS Lambda
- [x] Change to JSON
- [x] Implement GraphQL
- [x] Implement CI
- [ ] Setup a DB. Relational or document/nosql, configurable by environment (dev, prod, uat)
