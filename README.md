 # GRPC microservice layout
 An example that may be helpful for building your own microservice on golang.

 ## Technology stack used in this API:
 - A version of go micro framework for microservice development https://github.com/unistack-org/micro 
 - Consul for configuration of this API in JSON format (available here http://localhost:8500/ui/dc1/kv/go-micro-layouts/ if this launched within a quickstart)
 - Prometheus metrics for monitoring are available at /metrics endpoint
 - Unique request id is generated for every request that simplifies a search process in logs.
 - Healthchecks /live and /ready for monitoring a state of the service instance.

  ## Try this out (quickstart)
  Launch this with just a one single command (for representation):
  ```
  docker run --rm -it -e consul_host=http://host.docker.internal:8500 -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 -p 9090:9090 $(docker build -f .quickstart/Dockerfile -q .) sh /grpc-layout/.quickstart/quickstart.sh
  ```