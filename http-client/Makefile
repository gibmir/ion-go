GOBIN=$(GOPATH)/bin

GENDIR = $(CURDIR)/core
OUTPKG = mocks
OUTDIR = $(CURDIR)/$(OUTPKG)
.PHONY: mocks
mocks:
	$(GOBIN)/mockery --all --dir ${GENDIR} --output ${OUTDIR} --outpkg OUTPKG 
