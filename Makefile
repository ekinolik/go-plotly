.PHONY: build run-scatter run-multi-scatter run-bar run-stacked-bar run-horizontal-bar run-box run-custom-box run-box-statistical run-box-horizontal clean init

build:
	mkdir -p bin
	go build -o bin/scatter cmd/examples/scatter/main.go
	go build -o bin/multi_scatter cmd/examples/multi_scatter/main.go
	go build -o bin/bar cmd/examples/bar/main.go
	go build -o bin/stacked_bar cmd/examples/stacked_bar/main.go
	go build -o bin/horizontal_bar cmd/examples/horizontal_bar/main.go
	go build -o bin/box cmd/examples/box/main.go
	go build -o bin/custom_box cmd/examples/custom_box/main.go
	go build -o bin/box_statistical cmd/examples/box_statistical/main.go
	go build -o bin/box_horizontal cmd/examples/box_horizontal/main.go

run-scatter: build
	./bin/scatter

run-multi-scatter: build
	./bin/multi_scatter

run-bar: build
	./bin/bar

run-stacked-bar: build
	./bin/stacked_bar

run-horizontal-bar: build
	./bin/horizontal_bar

run-box: build
	./bin/box

run-custom-box: build
	./bin/custom_box

run-box-statistical: build
	./bin/box_statistical

run-box-horizontal: build
	./bin/box_horizontal

clean:
	rm -rf bin/
	rm -rf temp_plots/

# Add this to create the temp directory
init:
	mkdir -p temp_plots 