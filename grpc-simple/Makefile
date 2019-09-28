#!/bin/make

INPUT := ./api
OUTPUT := ./gens
SUFFIX := pb.go
PROTOC := protoc
FLAGS := --go_out=plugins=grpc:$(OUTPUT)

SRCS := $(shell find $(INPUT) -name *.proto)
DEPS := $(addsuffix .$(SUFFIX),$(basename $(SRCS)))

%.$(SUFFIX): %.proto
	mkdir -p $(OUTPUT)
	$(PROTOC) $(FLAGS) $*.proto

all: $(DEPS)

clean:
	rm -rd $(OUTPUT)