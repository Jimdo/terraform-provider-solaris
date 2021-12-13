SOLARISBANK_ENDPOINT ?= api.solaris-sandbox.de

export SOLARISBANK_ENDPOINT

default: testacc

guard-%:
	@ if [ "${${*}}" = "" ]; then \
	    echo "Environment variable $* not set"; \
	    exit 1; \
	fi

# Run acceptance tests
.PHONY: testacc
testacc: guard-SOLARISBANK_CLIENT_ID guard-SOLARISBANK_CLIENT_SECRET
	TF_ACC=1 go test ./... -v $(TESTARGS) -cover -timeout 120m

.PHONY: docs
docs:
	env -u SOLARISBANK_ENDPOINT \
	    -u SOLARISBANK_CLIENT_ID \
	    -u SOLARISBANK_CLIENT_SECRET \
	    go generate
