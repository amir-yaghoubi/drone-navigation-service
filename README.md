

# Drone Navigation Service
[![Tests](https://github.com/amir-yaghoubi/drone-navigation-service/actions/workflows/tests.yml/badge.svg)](https://github.com/amir-yaghoubi/drone-navigation-service/actions/workflows/tests.yml)

Simple service that will solve imaginary drones nagivation problem.

Currently service provided via **Rest API**, but **gRPC** will be available soon too.


## Development
Useful commands are available in `Makefile`, so check that out.



## Configurations
You can configure this service by providing following enviroment variables:
```bash

PRODUCTION=true
PORT=8080
SECTOR_ID=1

```


## Swagger Documentation
API documentations are available at `http://{BASE_URL}:{PORT}/swagger/index.html`.
For local enviroment you can try [this](http://localhost:8080/swagger/index.html).