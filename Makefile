include $(GOROOT)/src/Make.inc

TARG=fixme
GOFILES=fixme
GOFMT=gofmt -l -w

include $(GOROOT)/src/Make.cmd

test: format clean all
	./${TARG}

format:
	${GOFMT} .
