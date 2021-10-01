

# Drone Navigation Service
Simple service that will solve imaginary drones nagivation problem.

Currently service provided via **Rest API**, but **gRPC** will be available soon too.


## Development
Useful commands are available in `Makefile`, so check that out.



## Configurations
You can configure this service by providing following enviroment variables:
```bash

GIN_MODE=release
PORT=8080
SECTION_ID=1

```


## Swagger Documentation
API documentations are available at `http://{BASE_URL}:{PORT}/swagger/index.html`.
For local enviroment you can try [this](http://localhost:8080/swagger/index.html).