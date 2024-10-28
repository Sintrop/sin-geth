## Pre requisites
- Docker installed

## Build sin_geth
docker build -t sin_geth .

## Run sin_geth
docker run -it sin_geth

## Start a node
geth --identity Sequoia --datadir ./sequoia_a \
  --sequoia \
  --syncmode "full" \
  --networkid 1500 \
  --cache=1024 \
  --port 30303 \
  -authrpc.addr localhost --authrpc.port 8551 \
  --http.vhosts=* --http.addr "0.0.0.0" --http.port 8545 --http=true \
  --miner.threads=1 \
  --miner.etherbase=0x \
  --nat=none \
  console

  ## Start and stop mining
  miner.start()
  miner.stop()