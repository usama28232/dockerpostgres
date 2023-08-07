# dockerpostgres example
A minimal application to replicate isolated development environment by throwing Database to a docker image in terms of Schema & Seed


## Problem

Most of the times development is carried out in a separated isolated environment locally
Which makes it difficult to set up each machines by installing base software like database

## Solution

Here is a sample application which throws off the database on a docker image which will contain basic schema and seed data.

Following is the `Dockerfile` which copies inital schema, seed scripts and executes them on startup for you. 

You can publish your own image or expose the container over network from your machine, to allow other machines to access/develop/test

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

Build the docker image:
```
docker build . -t dockerpostgres:1.0
```

Run the docker image:

```
docker run -d -p 5432:5432 --name dockerpostgres dockerpostgres:1.0
```

That's it!!


### Feel free to edit/expand/explore this repository

For feedback and queries, reach me on LinkedIn at [here](https://www.linkedin.com/in/usama28232/?original_referer=)