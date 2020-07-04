# efishery-project
This repository is done in regard to complete eFishery's backend test. Goal of the test was to create two seperate backend services, one for authentication purpose and the other for fetching data.
## Overview project
This project consists of the following:
- Authentication Service (Go)
- Fetching Service (Python)
- MySQL Database
- Swagger Service (Go)
- Envoy Proxy

The below diagram represents the application architecture of the setup:

![Architecture Diagram 1](/static/efishery.png)

As shown in the diagram, the proxy is used to create a more stateless microservice. All traffic is accessed through the proxy so scaling out could be done. Auth service is connected through Envoy to a MySQL Database to store registered users. A Swagger service is also used to see the documentation of each services (auth and fetch).

Each service is deployed using docker-compose and can be run locally.

## Installation
If you have a runnig mysql server locally, you need to stop it beforehand so that the ports wont collide.
```
sudo systemctl stop mysql 
```

This project was testing using the following specifications:
- docker-compose 1.26.0
- docker 19.03.12

Initial Setup:
```
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
sudo curl -L "https://github.com/docker/compose/releases/download/1.26.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```
Test installation:
```
$ docker-compose --version
docker-compose version 1.26.0, build d4451659
```


In order to run the project us the following command:
```
docker-compose up --build -d
``` 
All the images will start to be created and be run as a container. You may check the running containers
```
docker ps -a
```
You may test out the services through the swagger service on URL:
```
http://localhost:9000
```
You may need to access the database after registering a new user, since the password is auto generated (4 charactes)
```
mysql -u root -P 1999 -h 127.0.0.1 -pefishery-DataBase-01

select name, phone, role, password from efishery.users;
```

## Production Architecture

The below diagram is my proposal deploying in production environment

Note: It may be a bit of a disclaimer to say it's too much overkill, but let's assume there is high load coming into the severs.

![Architecture Diagram 2](/static/efishery2.png)

### Database Choice
First of all, in this case i'm using MySQL as my choice of database since the data is still structured and straightforward.

### Master-slave Database
i would like to create a master-slave database to increase redundancy in case a failover is needed. Since we could use envoy as a L7 loadbalancer for mysql, we could increase the amount of replicas without changing any configuration in any services connected to it and with zero downtime.

### Caching
Since there is a need to fetch external data for currency conversion, it is ideal to create a caching mechanism. Here, i would like to use a separate redis instance so the fetch services could become stateless and all read from the same instance so scaling out is easy.

I would like to use cloudwatch with lamda (if aws) or a simple cron service deployed in an instance that will run every 5 minutes to fetch the currency conversion data and update the value to redis.

### Service Deployments
Ideally the services for auth and fetch apis would be deployed using kubernetes with Istio for ease of scaling purpose. A load balancer would also be used as a gateway for incoming traffic.


## Testing Out

Below are a few screenshots of testing out the services.

![Swagger 1](/static/1.png)

![Swagger 2](/static/2.png)

![Swagger 3](/static/3.png)