linux-solutions-go: linux-solutions.go
	go build -o $@ -ldflags -d $<

linux-solutions-welcome: linux-solutions.c
	gcc -Os -o $@ -static $<

.PHONY: all
all: linux-solutions-go linux-solutions-welcome
.PHONY: clean
clean:
	rm -f linux-solutions-go linux-solutions-welcome
