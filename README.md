<h1 align="center">docker-cloud-ec2-status</h1>

<p align="center">
  <b>Go</b> service which polls the Docker Cloud <a href="https://docs.docker.com/apidocs/docker-cloud/">REST API</a>, and removes itself from it's AWS auto-scaling group if it fails a Docker Cloud <a href="https://docs.docker.com/apidocs/docker-cloud/#perform-a-health-check-of-a-node">node health check</a>.
</p>

## Usage

In **development**:

```
go build && DC_USER=X DC_API_KEY=X DC_NODE_UUID=X ./docker-cloud-ec2-status
```

Environment variables:

```
- DC_USER      -> Docker Cloud username
- DC_API_KEY   -> Docker Cloud API key
- DC_NODE_UUID -> Any Docker Cloud node UUID
```
