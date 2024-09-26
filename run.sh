#!/bin/sh

docker run --rm -it -p 8080:8080 -v $(pwd)/static:/static phillezi/tinyhttpfileserver
