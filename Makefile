include $(GOROOT)/src/Make.inc

TARG=skel
GOFILES=skel.go
GOFMT=gofmt -l -w

include $(GOROOT)/src/Make.cmd

test: format clean all
	./${TARG}

format:
	${GOFMT} .
