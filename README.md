# SAMPLEAPI

A go api server to create sample tests for:

## TRIANGLE
URL: https://localhost:8000/triangle/{base}/{height}/area (GET)
Response: the area of the triangle as a json structure

## RECTANGLE
URL: https://localhost:8000/rectangle/{base}/{height}/area (GET)
Response: the area of the rectangle as a json structure

## CIRCLE 
URL: https://localhost:8000/circle/{radius}/area (GET)
Response: the area of the circle as a json structure

## SQUARE 
URL: https://localhost:8000/square/{base}/area (GET)
Response: the area of the square as a json structure

## ELLIPSE
URL: https://localhost:8000/ellipse/{width}/{height}/area (GET)
Response: the area of the ellipse as a json structure

## Trying it out

Pre-requisite:

docker: https://docs.docker.com/install/

3. make build

The executable is deposited into the `dist` folder.

## Building the docker image

Pre-requisite:
docker

`make image` will create an image tagged `sampleapi:latest`.

1. git clone https://github.com/agrapa/sampleapi.git
2. cd sampleapi
3. make image
4. make dockerrun
5. curl http://localhost:8000/triangle/10/5/area

## Other Makefile targets

`make` will display a menu for all the targets
