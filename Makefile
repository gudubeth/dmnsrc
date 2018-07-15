.PHONY: release clean test dep tools check-ver

release: check-ver clean
	goxc -wlc -pv=$(ver)
	goxc -wlc default publish-github -apikey=$(ghtoken)
	goxc -d bin -bc "linux"

clean:
	rm -Rf bin

test: 
	go test -v ./...

dep:
	dep ensure
		
get-tools: 
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/laher/goxc

check-ver:
	@ if [ "$(ver)" = "" ]; then \
		echo "'ver' is not set"; \
		exit 1; \
	fi

ci-validate:
	circleci config validate -c .circleci/config.yml

ci-build:
	circleci build

try:
	go run main.go check --benchmark ozgur.io \
	    example.com \
		example.org \
		example.net \
		nobodyshouldhavethisdomain-999.name \
		itu.edu.tr \
		alsonobodyshouldhavethisdomain-998.info \
		example.co.uk \
		123123123qweqweqwe123qwe.com \
		tostostostostosos123456.co.uk

try-stdin:
	cat fixtures/domain_names.txt | go run main.go check --benchmark

try-whois:
	go run main.go check --whois ozgur.io example.com example.org example.net itu.edu.tr example.co.uk 123123123qweqweqwe123qwe.com tostostostostosos123456.co.uk	