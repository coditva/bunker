all: bunker bunkerd

bunker:
	TARGET="bunker" sh scripts/build/binary.sh
bunkerd:
	TARGET="bunkerd" sh scripts/build/binary.sh

test: test_bunker test_bunkerd
test_bunker:
	TARGET="bunker" sh scripts/test/gotest.sh
test_bunkerd:
	TARGET="bunkerd" sh scripts/test/gotest.sh

tags:
	scripts/build/tags.sh

clean:
	rm -rf build/bunker build/bunkerd
