
.PHONY: proto
proto:
	docker build -t proto:latest .
	docker run --rm --name=proto -v `pwd`:/build proto:latest