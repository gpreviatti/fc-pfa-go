# FullCycle Program Acceleration Go

Simple application that was designed in the full cycle golang acceleration week. 
The objective of this program are learn and study simples concepts aboout golang, goroutines and pub-sub architecture.

## Technologies

- golang 1.19
- mongo
- rabbitmq
- docker
- kubernates
- prometheus
- grafana
- github actions

## How to run

If you would like to run, test and debug the producer, consumer or api you can use docker-compose file inside `infra/docker` with the command `docker-compose up -d` and attach in `goapp` service. However if you would like to run the entire project with prometheus and grafana you should apply the kubernates manifest files inside `infra/k8s`. ps: it is important to apply first the namespace manifest to create the namescape for the project.

## Diagram

![Untitled Diagram drawio](https://user-images.githubusercontent.com/24505963/194855059-c339ced6-3127-4a15-87a0-a1c7d99e985e.png)
