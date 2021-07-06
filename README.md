# Introduction

When starting sql-server in a docker-compose, needed a health check that would connect to sql-server and report
a _ready_ status. Something similar to the `pg_isready` utility. This is a simple command line utility that can
be included your docker setup and used in conjunction with the `healthcheck` directive for arranging service 
dependencies.
