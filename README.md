# dockerpostgres example
A minimal application to replicate isolated development environment by throwing Database to a docker image

## Problem

Most of the times development is carried out in a separated isolated environment locally which makes it difficult to set up each machines by installing database software and managing it

## Solution

Here is a sample application which throws off the database on a docker image which will contain basic schema and seed data.

Following is the `Dockerfile` which copies database schema, seed scripts and executes them on startup for you. 

You can publish your own image or expose the container over network from your machine, to allow other machines to access/develop/test/release

```
FROM postgres:latest

# Copy your migration scripts and seed data to the container
COPY init.sql /docker-entrypoint-initdb.d/
COPY seed_data.sql /docker-entrypoint-initdb.d/seed_data.sql

# Set environment variables
ENV POSTGRES_DB=db
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=user
```

To build the docker image:
```
docker build . -t dockerpostgres:1.0
```

To run the docker image:

```
docker run -d -p 5432:5432 --name dockerpostgres dockerpostgres:1.0
```

Here is the console output of execution:

```
Docker Postgres Example
Connected!
Fetching All Users
1 admin John Connor admin john@example.com 2023-08-07 05:09:54.20298 +0000 +0000 2023-08-07 05:09:54.20298 +0000 +0000
2 jimmy Jimmy Mcgill user jimmy@example.com 2023-08-07 05:09:54.204596 +0000 +0000 2023-08-07 05:09:54.204596 +0000 +0000
3 minh Minhaj Rivers user minh@example.com 2023-08-07 05:09:54.205558 +0000 +0000 2023-08-07 05:09:54.205558 +0000 +0000
4 steel Steel Rodd user steel@example.com 2023-08-07 05:09:54.206525 +0000 +0000 2023-08-07 05:09:54.206525 +0000 +0000
```

Instead of specifying Postgres connection inside docker image, you can also pass environment variables to docker run command as arguments

```
docker run -d -p 5432:5432 -e POSTGRES_DB=db -e POSTGRES_USER=user -e POSTGRES_PASSWORD=user --name dockerpostgres dockerpostgres:1.0
```

This may give you a bit more flexibility over your environments

That's it!!


### Feel free to edit/expand/explore this repository

For feedback and queries, reach me on LinkedIn at [here](https://www.linkedin.com/in/usama28232/?original_referer=)