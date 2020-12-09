# lb-check

Another stupid http server that checks if PostgreSQL replica is delayed and running.

Initially created to be used as a validator for postgres replicas under a AWS Network Load balancer.

Start then open http://localhost:300/check to run all checks. also check http://localhost:300/metrics for prometheus metrics.

TODO:
 - [ ] review and update all metrics
 - [x] expose metrics as prometheus metrics
 - [ ] add some tests