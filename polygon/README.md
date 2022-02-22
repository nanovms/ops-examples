Create your secrets and generate gensis:

```
./polygon-edge secrets init --data-dir test-chain-1
./polygon-edge genesis --consensus ibft --ibft-validators-prefix-path test-chain- --bootnode /ip4/0.0.0.0/tcp/10001/p2p/16Uiu2HAkwWuXvd7VqgoPi2sS9UB5Li2Y2S8sa1AVQqdinaanYRK7
```

Run:

```
eyberg@box:~/p/polygon-edge$ ops run -c config.json polygon-edge -p 10000 -p 10001 -p 10002
booting /home/eyberg/.ops/images/polygon-edge.img ...
en1: assigned 10.0.2.15
2022-02-22T16:09:25.476Z [INFO]  polygon: Data dir: path=./test-chain-1
2022-02-22T16:09:25.602Z [INFO]  polygon.blockchain: Current header: hash=0xcfee7f316796d7850b4c5df1f007442620874924579165fd35aa8e8d97331006 number=2
2022-02-22T16:09:25.606Z [INFO]  polygon.blockchain: genesis: hash=0x0d12ba710d43551729d2294ef9ed9f3b6a0e0fc86d2a97c2f963eeaf10e99e60
2022-02-22T16:09:25.631Z [INFO]  polygon.jsonrpc: http server started: addr=127.0.0.1:10002
2022-02-22T16:09:25.634Z [INFO]  polygon.consensus.ibft: validator key: addr=0x4387F5b9e1a6bdA6C3C9ec476e03437814BBbabd
2022-02-22T16:09:25.637Z [INFO]  polygon: GRPC server running: addr=127.0.0.1:10000
2022-02-22T16:09:25.639Z [INFO]  polygon.network: LibP2P server running: addr=/ip4/127.0.0.1/tcp/10001/p2p/16Uiu2HAkwWuXvd7VqgoPi2sS9UB5Li2Y2S8sa1AVQqdinaanYRK7
2022-02-22T16:09:25.644Z [INFO]  polygon.network: Omitting bootnode with same ID as host: id=16Uiu2HAkwWuXvd7VqgoPi2sS9UB5Li2Y2S8sa1AVQqdinaanYRK7
2022-02-22T16:09:25.660Z [INFO]  polygon.consensus.ibft: state change: new=SyncState
2022-02-22T16:09:25.663Z [INFO]  polygon.consensus.ibft: state change: new=AcceptState
2022-02-22T16:09:25.665Z [INFO]  polygon.consensus.ibft.acceptState: Accept state: sequence=3 round=1
2022-02-22T16:09:25.668Z [INFO]  polygon.consensus.ibft: current snapshot: validators=1 votes=0
2022-02-22T16:09:25.671Z [INFO]  polygon.consensus.ibft.acceptState: we are the proposer: block=3
2022-02-22T16:09:25.674Z [INFO]  polygon.consensus.ibft: executed txns: failed =0 successful=0 remaining in pool=0
2022-02-22T16:09:25.691Z [INFO]  polygon.consensus.ibft: build block: number=3 txns=0
2022-02-22T16:09:25.700Z [INFO]  polygon.consensus.ibft: state change: new=ValidateState
2022-02-22T16:09:25.725Z [INFO]  polygon.consensus.ibft: state change: new=CommitState
2022-02-22T16:09:25.727Z [INFO]  polygon.blockchain: write block: num=3 parent=0xcfee7f316796d7850b4c5df1f007442620874924579165fd35aa8e8d97331006
2022-02-22T16:09:25.735Z [INFO]  polygon.blockchain: new block: number=3 hash=0x167c0321267c7d659c71f35a558fe5f4f67150e2f22297467927c6abdd5945e4 txns=0 generation_time_in_seconds=170
2022-02-22T16:09:25.740Z [INFO]  polygon.consensus.ibft: block committed: sequence=3
```
