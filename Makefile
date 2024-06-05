CC         = go
BUILD_PATH = ./bin
SRC        = ./cmd/terminal-hook/main.go
TARGET     = th
BINS       = $(BUILD_PATH)/$(TARGET)
INST       = /usr/local/bin

.PHONY: all clean build run install

all: run

clean:
	rm -rf $(BUILD_PATH)

build: clean
	mkdir -p $(BUILD_PATH)
	$(CC) build -o $(BINS) $(SRC)

run:
	$(CC) run $(SRC) --about

install: build 
	sudo cp -v $(BINS) $(INST)

uninstall: clean
	sudo rm -f $(INST)/$(TARGET)
