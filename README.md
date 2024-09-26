# A tiny file-server

> [!NOTE]  
> This repository hosts a tiny containerized fileserver. It is probably not that secure and is only meant to be used internally in the docker/pod network for utility tasks. Please dont use it publicly.

## Overview

This tiny file server is a lightweight server for serving files within your Docker or pod networks. It allows for easy file access and management, making it ideal for utility tasks.

| Size | Comment |
|------|---------|
| 5.16MB | image on linux/amd64 (uncompressed |
| 2.14 MB | compressed image for linux/amd64 |

## Usage

To run the file server, you can use the following Docker command:
```bash
docker run --rm -it -p 8080:8080 -v ./your-dir:/static ghcr.io/cm1007-fhir-data/tinyfileserver:main
```
