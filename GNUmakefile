SOLARIS_ENDPOINT ?= api.solaris-sandbox.de

export SOLARIS_ENDPOINT

default: testacc

guard-%:
	@ if [ "${${*}}" = "" ]; then \
	    echo "Environment variable $* not set"; \
	    exit 1; \
	fi

# Run acceptance tests
.PHONY: testacc
testacc: guard-SOLARIS_CLIENT_ID guard-SOLARIS_CLIENT_SECRET
	TF_ACC=1 go test ./... -v $(TESTARGS) -cover -timeout 120m

.PHONY: docs
docs:
	env -u SOLARIS_ENDPOINT \
	    -u SOLARIS_CLIENT_ID \
	    -u SOLARIS_CLIENT_SECRET \
	    go generate
