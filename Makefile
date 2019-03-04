all: bunker bunkerd

bunker:
	TARGET="bunker" sh scripts/build/binary.sh

bunkerd:
	TARGET="bunkerd" sh scripts/build/binary.sh

clean:
	rm -rf build/bunker build/bunkerd
