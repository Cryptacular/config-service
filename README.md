# Go Configuration Service

## Overview

---

A website with API that allows the user to generate a configuration file based on settings, similar to something we'd use for InfoSmart Web. It will be hosted on AWS.

## Features

---

- Go
- AWS
- GraphQL
- Front end with settings
  - Environment (dev, prod, uat)
  - ...
- API
- Database
- Unit tests

## To Do

---

- [ ]  Start with a web API framework that's good at apis.
- [ ]  Create a configuration endpoint that serializes a similar configuration that we might use for say ISW
- [ ]  Setup a DB. Relational or document/nosql, configurable by environment (dev, prod, uat)
- [ ]  Integrate GraphQL to fetch data from database
