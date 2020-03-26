# SAMPLEAPI

This is a simple API with 5 calls for the purpose of exercising your test writing abilities. 
The API is provided in source code form so you can see what is going behind the scenes. I am also
providing a simple way to build and run the API as a service. The only requirement is Docker.
If you need to install docker, you can find the process here: https://docs.docker.com/install/

Your task is to build an automated suite that can test these API calls. You can choose any language and tools you want. 
We are not testing your familiarity with a lnguage, but rather your ability to design and implement a test given an API.
Ideally the tests will cover functionality, performance and load (capacity).

The calls are described below. Each one of them calculates the area for a different shape.

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
