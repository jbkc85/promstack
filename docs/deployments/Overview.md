# Deployments

There are multiple ways to deploy the PromStack including a single monolithic server with all software running, as well as pieces of the stack distributed across a clustered environment.  There is no *ideal or best-practice* that we know of at this time, so I will just provide the scenarios we have deployed in.

### Deploying in Kubernetes

### Deploying in Docker with docker-compose.yml

The docker-compose.yml is the simplest form of running PromStack locally for those who just have Docker installed.  It **does not** store data outside of a container volume, so be aware that any graphs created in Grafana that are not exported, or metrics collected by Prometheus will survive a container volume reboot.

[Read More about the docker-compose.yml deployment example](Docker.md)


### Deploying on Container Linux with cloud-config.yml

COMING SOON.. (may be replaced by Ignition)
