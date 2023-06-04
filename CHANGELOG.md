# Changelog

## [0.9.0](https://github.com/taikoxyz/taiko-client/compare/v0.8.0...v0.9.0) (2023-06-04)


### Features

* **all:** check L1 reorg before each operation ([#252](https://github.com/taikoxyz/taiko-client/issues/252)) ([e76b03f](https://github.com/taikoxyz/taiko-client/commit/e76b03f4af7ab1d300d206c246f736b0c5cb2241))
* **all:** rename `treasure` to `treasury` ([#233](https://github.com/taikoxyz/taiko-client/issues/233)) ([252959f](https://github.com/taikoxyz/taiko-client/commit/252959f6e80f731da7526c655aeac0eec3b428b2))
* **all:** update protocol bindings and some related changes ([#237](https://github.com/taikoxyz/taiko-client/issues/237)) ([3e12042](https://github.com/taikoxyz/taiko-client/commit/3e12042a9a5b5b9baca7de1b342788b22b2ca17e))
* **bindings:** update bindings with EthDeposit changes ([#255](https://github.com/taikoxyz/taiko-client/issues/255)) ([f91f2dd](https://github.com/taikoxyz/taiko-client/commit/f91f2dd64e1fe25bc55790a8a93ea0ffab54ca3b))
* **bindings:** update go contract bindings ([#243](https://github.com/taikoxyz/taiko-client/issues/243)) ([132500e](https://github.com/taikoxyz/taiko-client/commit/132500e27d135e6e5f89c96716a0bb2d17b6801b))
* **driver:** optimize reorg handling && add more tests ([#256](https://github.com/taikoxyz/taiko-client/issues/256)) ([20c38a1](https://github.com/taikoxyz/taiko-client/commit/20c38a171ef617ddeecbe325d29d64c963792c07))
* **pkg:** do not return error when genesis block not found ([#244](https://github.com/taikoxyz/taiko-client/issues/244)) ([8033e31](https://github.com/taikoxyz/taiko-client/commit/8033e31728c946a80fdd3d07f737241c7e19edf8))
* **proof_producer:** update request parameters based on new circuits changes ([#240](https://github.com/taikoxyz/taiko-client/issues/240)) ([31521ef](https://github.com/taikoxyz/taiko-client/commit/31521ef8b7362dacbf183dc8c7d9a6020d1b0fc4))
* **proposer:** add a `--minimalBlockGasLimit` flag to mitigate the potential gas estimation issue ([#225](https://github.com/taikoxyz/taiko-client/issues/225)) ([ab8305d](https://github.com/taikoxyz/taiko-client/commit/ab8305d39d1ca3375c6477b84d4afe5c729e815f))
* **proposer:** add a new metric to track block fee ([#224](https://github.com/taikoxyz/taiko-client/issues/224)) ([98c17f0](https://github.com/taikoxyz/taiko-client/commit/98c17f00ade4fa20251a59b3aba4cad9e1eb1bd8))
* **proposer:** propose multiple L2 blocks in one L1 block ([#254](https://github.com/taikoxyz/taiko-client/issues/254)) ([36ba5db](https://github.com/taikoxyz/taiko-client/commit/36ba5dbcc2863dc34fda2e59bf8a9d30d3665d04))
* **prover:** add `--expectedReward` flag ([#248](https://github.com/taikoxyz/taiko-client/issues/248)) ([f64a762](https://github.com/taikoxyz/taiko-client/commit/f64a7620726019a2e7f5eada7b92087663b273fd))
* **prover:** improve proof submission delay calculation ([#249](https://github.com/taikoxyz/taiko-client/issues/249)) ([7cc5d54](https://github.com/taikoxyz/taiko-client/commit/7cc5d541bef0eac9078bc93eb5f1d9954b164e9b))
* **prover:** normal prover should wait targetProofTime before submitting proofs ([#232](https://github.com/taikoxyz/taiko-client/issues/232)) ([2128ddc](https://github.com/taikoxyz/taiko-client/commit/2128ddc325aaf8acf538fdd50e299187da8543dd))
* **prover:** remove submission delay when running as a system prover ([#221](https://github.com/taikoxyz/taiko-client/issues/221)) ([49a25dd](https://github.com/taikoxyz/taiko-client/commit/49a25dd72888ee54209ddce51c6a701803728d86))
* **prover:** remove the unnecessary special proof delay ([#226](https://github.com/taikoxyz/taiko-client/issues/226)) ([dcead44](https://github.com/taikoxyz/taiko-client/commit/dcead44a32ec9d064af423af0f2effea8b819fca))
* **prover:** updates based on protocol `proofTimeTarget` changes ([#227](https://github.com/taikoxyz/taiko-client/issues/227)) ([c6ea860](https://github.com/taikoxyz/taiko-client/commit/c6ea860d736828fdb50e16447dee44733371c06f))
* **repo:** enable OpenAI-based review ([#235](https://github.com/taikoxyz/taiko-client/issues/235)) ([88e4dae](https://github.com/taikoxyz/taiko-client/commit/88e4dae2e37c58273438335daade21587f25ec27))


### Bug Fixes

* **driver:** handle reorg ([#216](https://github.com/taikoxyz/taiko-client/issues/216)) ([fc2ec63](https://github.com/taikoxyz/taiko-client/commit/fc2ec637f5509b67572bb4d978f7bc41860e9b43))
* **flag:** add a missing driver flag to configuration ([#246](https://github.com/taikoxyz/taiko-client/issues/246)) ([0b60243](https://github.com/taikoxyz/taiko-client/commit/0b60243fbc03bbfc2aceb8933ae9901d4b385117))
* **prover:** fix an issue in prover event loop ([#257](https://github.com/taikoxyz/taiko-client/issues/257)) ([c550f09](https://github.com/taikoxyz/taiko-client/commit/c550f09d33f638f38461e576684432d90d850ac3))
* **prover:** update bindings && fix a delay calculation issue ([#242](https://github.com/taikoxyz/taiko-client/issues/242)) ([49c3d69](https://github.com/taikoxyz/taiko-client/commit/49c3d6957b296b1312a53fcb5122fcd944b77c2d))
* **repo:** fix openAI review workflow ([#253](https://github.com/taikoxyz/taiko-client/issues/253)) ([f44530b](https://github.com/taikoxyz/taiko-client/commit/f44530b428396b8514f974cf8ec476078d20c9d6))

## [0.8.0](https://github.com/taikoxyz/taiko-client/compare/v0.7.0...v0.8.0) (2023-05-12)


### Features

* **proposer:** check tko balance and fee before proposing ([#205](https://github.com/taikoxyz/taiko-client/issues/205)) ([cc0da63](https://github.com/taikoxyz/taiko-client/commit/cc0da632c825c1379f039f489d7426548527cc80))
* **prover:** add oracle proof submission delay ([#199](https://github.com/taikoxyz/taiko-client/issues/199)) ([7b5ed94](https://github.com/taikoxyz/taiko-client/commit/7b5ed94d12b0982de46e5ed66b38cffcf9c0c0d4))
* **prover:** add special prover (system / oracle) ([#214](https://github.com/taikoxyz/taiko-client/issues/214)) ([1020377](https://github.com/taikoxyz/taiko-client/commit/1020377bec7115efd757a6c2ea78cfe9a97b6430))
* **prover:** cancel proof if it becomes verified ([#207](https://github.com/taikoxyz/taiko-client/issues/207)) ([74d1729](https://github.com/taikoxyz/taiko-client/commit/74d17296c48a323e3ed78424b98aea9a93e081ca))
* **prover:** implementing `--graffiti` flag for prover as input to block evidence ([#209](https://github.com/taikoxyz/taiko-client/issues/209)) ([2340210](https://github.com/taikoxyz/taiko-client/commit/2340210437a14618774265d2ad2f80989296aeae))
* **prover:** improve oracle proof submission delay ([#212](https://github.com/taikoxyz/taiko-client/issues/212)) ([20c1423](https://github.com/taikoxyz/taiko-client/commit/20c14235b087e4624427879aa587a1599690dbbb))
* **prover:** update `ZkevmRpcdProducer` to integrate new circuits ([#217](https://github.com/taikoxyz/taiko-client/issues/217)) ([81cf612](https://github.com/taikoxyz/taiko-client/commit/81cf6120c1610f7a8edaa183eb9a0fbbeb45b5f1))
* **prover:** update canceling proof logic ([#218](https://github.com/taikoxyz/taiko-client/issues/218)) ([21d7e78](https://github.com/taikoxyz/taiko-client/commit/21d7e78d2e83fdd060fbc0303b244dee9777fcc4))
* **prover:** update skip checking for system prover ([#215](https://github.com/taikoxyz/taiko-client/issues/215)) ([79ba210](https://github.com/taikoxyz/taiko-client/commit/79ba2104216dfee0a1b1556c4abc5abc76c5a266))


### Bug Fixes

* **driver:** fix `GetBasefee` parameters ([#210](https://github.com/taikoxyz/taiko-client/issues/210)) ([b5dc5c5](https://github.com/taikoxyz/taiko-client/commit/b5dc5c589d26b8e9e2420ecb38ea5c83b2ae7c2e))
* **prover:** fix some oracle proof submission issues ([#211](https://github.com/taikoxyz/taiko-client/issues/211)) ([e061540](https://github.com/taikoxyz/taiko-client/commit/e06154058127962b90d5ab4a95cfec7c71942de3))
* **prover:** submit L2 signal root with submitting proof ([#220](https://github.com/taikoxyz/taiko-client/issues/220)) ([8b030ed](https://github.com/taikoxyz/taiko-client/commit/8b030ed1a8fcf1a948a2272ff8ae3927c8957d84))
* **prover:** submit L2 signal service root instead of L1 when submitting proof ([#219](https://github.com/taikoxyz/taiko-client/issues/219)) ([74fe156](https://github.com/taikoxyz/taiko-client/commit/74fe1567d0cc43e2d26d3f4af777794bc6c3a9f5))

## [0.7.0](https://github.com/taikoxyz/taiko-client/compare/v0.6.0...v0.7.0) (2023-04-28)


### Features

* **all:** update client softwares based on the new protocol upgrade ([#185](https://github.com/taikoxyz/taiko-client/issues/185)) ([54f7a4c](https://github.com/taikoxyz/taiko-client/commit/54f7a4cb2db72a4ffa9a199e2af1f0d709a1ac27))
* **driver:** changes based on protocol L2 EIP-1559 design ([#188](https://github.com/taikoxyz/taiko-client/issues/188)) ([82e8b97](https://github.com/taikoxyz/taiko-client/commit/82e8b9741782258840696701993b6d009d0260e0))
* **prover:** add oracle prover flag ([#194](https://github.com/taikoxyz/taiko-client/issues/194)) ([ebbc725](https://github.com/taikoxyz/taiko-client/commit/ebbc72559a70c9aefc34286b05b1f4261bae8cd6))
* **prover:** proof skip ([#198](https://github.com/taikoxyz/taiko-client/issues/198)) ([8607af8](https://github.com/taikoxyz/taiko-client/commit/8607af826ed9561a6bdae74074a517f1424e7a69))

## [0.6.0](https://github.com/taikoxyz/taiko-client/compare/v0.5.0...v0.6.0) (2023-03-20)


### Features

* **docs:** remove concept docs and refer to website ([#180](https://github.com/taikoxyz/taiko-client/pull/180)) ([a8dcdac](https://github.com/taikoxyz/taiko-client/commit/a8dcdac77c1a5e3f85e4d7a4b912cfb3d903a3d9))
* **flags:** update txpool.locals flag usage ([#181](https://github.com/taikoxyz/taiko-client/pull/181)) ([dac6102](https://github.com/taikoxyz/taiko-client/commit/dac6102d7508b9bdcb248eab4dcf469022353aa8))
* **proposer:** add `proposeEmptyBlockGasLimit` ([#178](https://github.com/taikoxyz/taiko-client/issues/178)) ([e64d769](https://github.com/taikoxyz/taiko-client/commit/e64d769f45d072b151f429f61c1fe2ab07dec0dc))


## [0.5.0](https://github.com/taikoxyz/taiko-client/compare/v0.4.0...v0.5.0) (2023-03-08)


### Features

* **pkg:** improve `BlockBatchIterator` ([#173](https://github.com/taikoxyz/taiko-client/issues/173)) ([4fab06a](https://github.com/taikoxyz/taiko-client/commit/4fab06a9cba17c5e4da09acbe9b95949d6c4d47f))
* **proposer,prover:** make `context.Context` part of `proposer.waitTillSynced` && `ProofProducer.RequestProof`'s parameters ([#169](https://github.com/taikoxyz/taiko-client/issues/169)) ([4b11e29](https://github.com/taikoxyz/taiko-client/commit/4b11e29689b8fac85023669443c351f428a54fea))
* **proposer:** new flag to propose empty blocks ([#175](https://github.com/taikoxyz/taiko-client/issues/175)) ([6621a5c](https://github.com/taikoxyz/taiko-client/commit/6621a5c89a92e7593f702e4c82e69d1215b2ca59))
* **proposer:** remove `poolContentSplitter` in proposer ([#159](https://github.com/taikoxyz/taiko-client/issues/159)) ([e26c831](https://github.com/taikoxyz/taiko-client/commit/e26c831a42fdf448b32bcf75c1f1f87bd71df481))
* **proposer:** remove an unused flag ([#176](https://github.com/taikoxyz/taiko-client/issues/176)) ([7d2126e](https://github.com/taikoxyz/taiko-client/commit/7d2126efe256bcb698b3d4df7352efdbff957ace))
* **prover:** ensure L2 EE is fully synced when calling `initL1Current` ([#170](https://github.com/taikoxyz/taiko-client/issues/170)) ([6c85058](https://github.com/taikoxyz/taiko-client/commit/6c8505827c035cc7456967bc8aab8bec1861e19b))
* **prover:** new flags for `zkevm-chain` ([#166](https://github.com/taikoxyz/taiko-client/issues/166)) ([1c90a3d](https://github.com/taikoxyz/taiko-client/commit/1c90a3d6b7cada0b116875d88f0952993b54bb5f))
* **prover:** tracking for most recent block ID to ensure (relatively) consecutive proving by notification system ([#174](https://github.com/taikoxyz/taiko-client/issues/174)) ([e500039](https://github.com/taikoxyz/taiko-client/commit/e5000395a3a28bd282df64f54867fd771143a56a))


### Bug Fixes

* **proposer:** remove an unused metric from proposer ([#171](https://github.com/taikoxyz/taiko-client/issues/171)) ([8df5eea](https://github.com/taikoxyz/taiko-client/commit/8df5eea1d9f1482a10b7d395ae19953f5d6ea6ce))

## [0.4.0](https://github.com/taikoxyz/taiko-client/compare/v0.3.0...v0.4.0) (2023-02-22)


### Features

* **all:** update contract bindings && some improvements based on Alex's feedback ([#153](https://github.com/taikoxyz/taiko-client/issues/153)) ([bdaa292](https://github.com/taikoxyz/taiko-client/commit/bdaa2920bcb113d3887409edb17462b5e0d3a2c5))
* **bindings:** parse solidity custom errors ([#163](https://github.com/taikoxyz/taiko-client/issues/163)) ([9a79127](https://github.com/taikoxyz/taiko-client/commit/9a79127a5a3cddf4e95ac899943e6551b02cf432))


### Bug Fixes

* **driver:** fix an issue in sync status checking ([#162](https://github.com/taikoxyz/taiko-client/issues/162)) ([4b21027](https://github.com/taikoxyz/taiko-client/commit/4b2102720e2c1c2fcaef1853ad74b91c6d08aaaa))
* **proposer:** fix a proposer nonce order issue ([#157](https://github.com/taikoxyz/taiko-client/issues/157)) ([80fc0e9](https://github.com/taikoxyz/taiko-client/commit/80fc0e94d819f93ecdeac492eb1f35d5f2bb09ce))

## [0.3.0](https://github.com/taikoxyz/taiko-client/compare/v0.2.4...v0.3.0) (2023-02-15)


### Features

* **prover:** improve the check for whether the current block still needs a new proof ([#145](https://github.com/taikoxyz/taiko-client/issues/145)) ([6c00fc5](https://github.com/taikoxyz/taiko-client/commit/6c00fc544b1ed92a4e38860059ef463282648a42))
* **prover:** update `ZkevmRpcdProducer` to make it connecting to a real proverd service ([#121](https://github.com/taikoxyz/taiko-client/issues/121)) ([8c8ee9c](https://github.com/taikoxyz/taiko-client/commit/8c8ee9c2c3266e25e4233821034b89f50bb08c33))
* **repo:** implement release please ([#148](https://github.com/taikoxyz/taiko-client/issues/148)) ([d8f3ad8](https://github.com/taikoxyz/taiko-client/commit/d8f3ad80d358fe547d356b7f7d7fd6e6ca9279ce))
