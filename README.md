# Demo for Prometheus histogram visualization

## Requirements
Installed docker engine & docker-compose

## Run
```shell
sh init.sh
```

It starts 4 containers (Prometheus & Grafana, application and load generator).

Generator starts sending requests in loop with small delay.
Requests duration will be visualized by Grafana.