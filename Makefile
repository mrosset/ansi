include $(GOROOT)/src/Make.inc

TARG=skel
GOFILES=skel.go
GOFMT=gofmt -l -w

include $(GOROOT)/src/Make.cmd

test: format all
	./${TARG}

format:
	${GOFMT} .
