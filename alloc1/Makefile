GO=go
TARGETS=main.txt analysis1.txt analysis2.txt analysis3.txt

all: $(TARGETS)

.SUFFIXES:
.SUFFIXES: .txt .go
.go.txt:
	$(GO) build -tags $* -gcflags "-m -v" -ldflags "-v" 2>&1 | cat > $@

clean:
	$(GO) clean
	rm -f $(TARGETS)
