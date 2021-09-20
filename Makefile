GO_FILE_DIRS = \
resource/detectors/aws/ec2 \
resource/detectors/aws/ecs \
resource/detectors/aws/eks \
resource/detectors/aws/lambda \
resource/detectors/gcp/cloudfunction \
resource/detectors/gcp/gce \
resource/detectors/gcp/gke 

TOOLS_MOD_DIR := ./tools

GOTEST_MIN = go test 
GOTEST = $(GOTEST_MIN) 
GOTEST_WITH_COVERAGE = $(GOTEST) -coverprofile cover.out ./...

TOOLS_DIR := $(abspath ./tools)

.PHONY: precommit

$(TOOLS_DIR)/gocovmerge: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/gocovmerge github.com/wadey/gocovmerge

.PHONY: test
test:
	set -e; for dir in $(GO_FILE_DIRS); do \
	  (cd "$${dir}" && \
	    $(GOTEST) ); \
	done; 

.PHONY: trial
trial:
	go build -o $(TOOLS_MOD_DIR)/gocovmerge github.com/wadey/gocovmerge
	$(TOOLS_MOD_DIR)/gocovmerge $(shell find . -name cover.out) > coverage.txt
	rm $(shell find . -name cover.out)

.PHONY: cover
cover:
	$(GOTEST_WITH_COVERAGE)
 
.PHONY: test-with-cover
test-with-cover:
	set -e; for dir in $(GO_FILE_DIRS); do \
	  (cd "$${dir}" && \
	    $(GOTEST_WITH_COVERAGE) ); \
	done;
	go build -o $(TOOLS_MOD_DIR)/gocovmerge github.com/wadey/gocovmerge
	$(TOOLS_MOD_DIR)/gocovmerge $(shell find . -name cover.out) > coverage.txt
	rm $(shell find . -name cover.out)
	rm $(TOOLS_DIR)/gocovmerge