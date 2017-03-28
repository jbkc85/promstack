# PromStack

**THIS SOLUTION IS IN BETA**

PromStack is a collection of software intended to provide a full featured platform for monitoring, alerting and logging of alerts.

> PromStack is NOT associated with the Prometheus group, though I do utilize Prometheus and their awesome software (given the name being 'Prom'Stack!)


**Table of Contents**

* <a href="#description">Description</a>
* <a href="#getting-started">Getting Started</a>
* <a href="#status-and-contributing">Status and Contributing</a>

## <a name="description"></a>Description

As mentioned before, PromStack is a collection of software intended to provide a full featured platform for monitoring, alerting and logging of alerts.  Most of the software in this collection are highly considered in their respective markets, though I (which hopefully turns into a collective 'we') still might change pieces as time moves along to optimize and offer more/better functionality as it comes along.

### A quick look at the Stack

<center>

![](https://cloud.githubusercontent.com/assets/215878/24385277/e3831784-132e-11e7-87ee-bf7448ca591c.png)

</center>

#### Monitoring with Prometheus

If you are looking at this page, [Prometheus](https://prometheus.io) should be pretty well known to you by now.  Prometheus is an open-source monitoring solution that scrapes a configured endpoint for Metrics and allows the user to configure rules and alerts based off of those metrics.

#### Service Discovery with Consul

[Consul](https://www.consul.io/intro/index.html) is a lot of things.  However in PromStack, I am utilizing Consul for its Service Discovery via the Consul Catalog for Nodes and Services as well as the DNS provided from the Catalog functionality.  I am also utilizing the Key Value store for the promstackctl tool.

#### Alerting with AlertManager

[AlertManager](https://prometheus.io/docs/alerting/alertmanager/) is a tool that handles Prometheus alert events and sends the events off to receivers based on labels defined in AlertManager's configuration.

#### Logging with Alerta

An experimental piece of the PromStack, [Alerta](http://alerta.io) Python based application that relies on MongoDB.  Alerta's purpose in this Stack is to provide a logging platform for the Alerts generated by Prometheus for long-term storage.

#### Visualization with Grafana

[Grafana](https://grafana.com) is an open-source graphing tool for time series analytics.  Grafana creates be-u-T-ful graphs from Prometheus Metrics.

#### Tools

**[PromstackCTL](promstackctl/)**: GoLang utility to add pre-defined exporters (configured in Consul KV store) to machines in your cluster.

### More Information

For more information about the layout of the file structure, communication between the software in the Stack, please read more here[*under construction*].

## <a name="getting-started"></a>Getting Started

To get started, we first need to deploy the essential tools of PromStack.  There are multiple ways to deploy the PromStack including a single monolithic server with all software running, as well as pieces of the stack distributed across a clustered environment.  There is no *ideal or best-practice* that I know of at this time, so I will just provide the scenarios I have deployed in.  [Here is an overview of deployment scenarios I have had worked out](docs/deployments/Overview.md).

## <a name="status-and-contributing"></a>Status and Contributing

The current status of PromStack is that it is in **beta**.  Since there are parts that I have deemed 'experimental', I do not wish to say it is 100% stable at this time.  Though, with that being said, running these components in production should provide a stable monitoring and alerting environment, as Alerta (being the logging portion) is the only part under review.

### Contributing

Feedback is greatly appreciated, so please email jbkc85 at gmail.com or open an issue.  Documentation is ongoing, but if you feel like something is lacking please let us know!

### Acknowledgements and other Solutions

I decided to write our deployment of PromStack after reading Percona's [Monitoring and Management](https://www.percona.com/doc/percona-monitoring-and-management/deploy/index.html) write-up.  A lot of the features between the two stacks are similar, and I figured writing the solution up might help others as well as benefit myself in the long run to get feedback on what everyone else out there is doing.

