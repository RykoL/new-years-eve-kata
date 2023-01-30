EXECUTABLE=new-years-eve
SRC=$(shell find . -name "*.go")

.PHONY=clean start

## Backend

start: $(EXECUTABLE)
	./new-years-eve

clean:
	rm $(EXECUTABLE) || true

$(EXECUTABLE): $(SRC)
	go build .

prepare-for-deployment:
	mkdir -p api
	cp handler/wishes.go api/
	cp go.mod api/
	cp go.sum api/
