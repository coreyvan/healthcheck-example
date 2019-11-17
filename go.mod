module github.com/coreyvan/healthcheck-example

go 1.13

require (
	github.com/coreyvan/healthcheck-example/health v0.0.0
	github.com/docker/distribution v2.7.1+incompatible
)

replace github.com/coreyvan/healthcheck-example/health => ./health
