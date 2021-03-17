process_name = pwsync
author_name = kjcodeacct

local:
	GO111MODULE=on CGO_ENALBED=0 go build -v -a

# test at 1 level of subdirectories
test:
	go test -p 1 ./...

binaries: linux_amd64 linux_i386 pi_arm64 pi_arm mac_amd64

linux_amd64:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o binaries/$@/${process_name}
	cp README.md binaries/$@/
	# tar files inside the specified release directory
	tar -czvf binaries/${process_name}.$@.tar.gz -C binaries/$@/ ${process_name} README.md

linux_i386:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -v -a -o binaries/$@/${process_name}
	cp README.md binaries/$@/
	# tar files inside the specified release directory
	tar -czvf binaries/${process_name}.$@.tar.gz -C binaries/$@/ ${process_name} README.md

pi_arm64:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -a -o binaries/$@/${process_name}
	cp README.md binaries/$@/
	# tar files inside the specified release directory
	tar -czvf binaries/${process_name}.$@.tar.gz -C binaries/$@/ ${process_name} README.md

pi_arm:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -v -a -o binaries/$@/${process_name}
	cp README.md binaries/$@/
	# tar files inside the specified release directory
	tar -czvf binaries/${process_name}.$@.tar.gz -C binaries/$@/ ${process_name} README.md

mac_amd64:
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -a -o binaries/$@/${process_name}
	cp README.md binaries/$@/
	# tar files inside the specified release directory
	tar -czvf binaries/${process_name}.$@.tar.gz -C binaries/$@/ ${process_name} README.md

cleanup_binaries:
	rm -rf binaries

docker:
	docker build -t ${author_name}/${process_name} .