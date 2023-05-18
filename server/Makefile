GOBIN=$(GOPATH)/bin

GENDIR = $(CURDIR)
OUTPKG = mocks
OUTDIR = $(CURDIR)/test/$(OUTPKG)
.PHONY: mocks
mocks:
	$(GOBIN)/mockery --all --dir ${GENDIR} --output ${OUTDIR} --outpkg OUTPKG 
