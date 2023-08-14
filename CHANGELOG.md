# Changelog

## [0.14.0](https://github.com/taikoxyz/taiko-client/compare/v0.13.0...v0.14.0) (2023-08-09)


### Features

* **bindings:** update `TaikoL1BlockMetadataInput` ([#359](https://github.com/taikoxyz/taiko-client/issues/359)) ([1beae59](https://github.com/taikoxyz/taiko-client/commit/1beae59cfbe1345a5bb69714b25ba4397173be45))
* **bindings:** update go contract bindings ([#346](https://github.com/taikoxyz/taiko-client/issues/346)) ([c6454af](https://github.com/taikoxyz/taiko-client/commit/c6454afe28b3a86c8d33c8434cfd345318116076))
* **bindings:** update go contract bindings ([#352](https://github.com/taikoxyz/taiko-client/issues/352)) ([b9da8f6](https://github.com/taikoxyz/taiko-client/commit/b9da8f68e733a51255c1307d016d1ff9e241f3c9))
* **driver:** update `l1Current` check in `ProcessL1Blocks` ([#340](https://github.com/taikoxyz/taiko-client/issues/340)) ([d67f287](https://github.com/taikoxyz/taiko-client/commit/d67f287bd5cce08aa5b7ba9fd33fc00e91ad6190))
* **pkg:** add default timeout for `GetStorageRoot` ([#347](https://github.com/taikoxyz/taiko-client/issues/347)) ([9a4dee0](https://github.com/taikoxyz/taiko-client/commit/9a4dee04f90e521832efef5febeebb1231e22a19))
* **pkg:** improve archive node check ([#334](https://github.com/taikoxyz/taiko-client/issues/334)) ([c6cd1b0](https://github.com/taikoxyz/taiko-client/commit/c6cd1b0492499b3c686ac282d65743793bd162da))
* **pkg:** introduce `EthClient` with a timeout attached ([#337](https://github.com/taikoxyz/taiko-client/issues/337)) ([1608aba](https://github.com/taikoxyz/taiko-client/commit/1608abae268bbbe6671ec9eb89fed2846065852c))
* **pkg:** optimize `CheckL1ReorgFromL1Cursor` ([#329](https://github.com/taikoxyz/taiko-client/issues/329)) ([ed63c1f](https://github.com/taikoxyz/taiko-client/commit/ed63c1f8e4ba6a9fd40b1d1d5f3bba217d470f4b))
* **pkg:** Wait receipt timeout ([#343](https://github.com/taikoxyz/taiko-client/issues/343)) ([cf261d3](https://github.com/taikoxyz/taiko-client/commit/cf261d377f61ea0b0ff049be7e8c8eb75264f386))
* **proposer:** add `--proposeBlockTxGasTipCap` flag ([#349](https://github.com/taikoxyz/taiko-client/issues/349)) ([e40115b](https://github.com/taikoxyz/taiko-client/commit/e40115b97002661def8eed8dfb768ad28c19f0ea))
* **proposer:** update pool content query ([#341](https://github.com/taikoxyz/taiko-client/issues/341)) ([221a3b9](https://github.com/taikoxyz/taiko-client/commit/221a3b92b77f4b3d3e5499eb27fa289ae44b0151))
* **proposer:** use `TaikoConfig.blockMaxGasLimit` as proposed block gasLimit && remove some unused flags ([#344](https://github.com/taikoxyz/taiko-client/issues/344)) ([f0a3da7](https://github.com/taikoxyz/taiko-client/commit/f0a3da7d6bf8af222ae6e780218ccca2c7861137))
* **prover:** add `--proofSubmissionMaxRetry` flag ([#333](https://github.com/taikoxyz/taiko-client/issues/333)) ([8d92b7a](https://github.com/taikoxyz/taiko-client/commit/8d92b7aa96d22ca20de57fd02e52d7f3f6ff9a5f))
* **prover:** changes based on `proofVerifer` protocol updates ([#338](https://github.com/taikoxyz/taiko-client/issues/338)) ([6dcb34a](https://github.com/taikoxyz/taiko-client/commit/6dcb34aab3619731852a19a09b54aadce34de999))
* **prover:** prove block tx gas limit ([#357](https://github.com/taikoxyz/taiko-client/issues/357)) ([8ed4da2](https://github.com/taikoxyz/taiko-client/commit/8ed4da2f0bd0bf5f215767b1bd44106dd878431f))
* **rpc:** check if L1 rpc is an archive node ([#332](https://github.com/taikoxyz/taiko-client/issues/332)) ([b1aa1d3](https://github.com/taikoxyz/taiko-client/commit/b1aa1d388d407f2f5cb14275c006b1a22213b8ff))


### Bug Fixes

* **pkg:** fix returned context error from `WaitL1Origin` ([#331](https://github.com/taikoxyz/taiko-client/issues/331)) ([0ebf121](https://github.com/taikoxyz/taiko-client/commit/0ebf121dcae5e75d359bc7818aa98fa6f7b1bc20))
* **pkg:** set more RPC context timeout ([#356](https://github.com/taikoxyz/taiko-client/issues/356)) ([ffe2f90](https://github.com/taikoxyz/taiko-client/commit/ffe2f906808f99a48f6a848351c9a34ea63f02b7))
* **prover:** default prove unassigned blocks to false ([#354](https://github.com/taikoxyz/taiko-client/issues/354)) ([ed34ef6](https://github.com/taikoxyz/taiko-client/commit/ed34ef670a3deef5f4db88429cd13c5bdb108289))
* **prover:** fix `onBlockProposed` reorg detection ([#348](https://github.com/taikoxyz/taiko-client/issues/348)) ([4877e01](https://github.com/taikoxyz/taiko-client/commit/4877e01f7c35f0cbce329e14948dd78b5de0c911))

## [0.13.0](https://github.com/taikoxyz/taiko-client/compare/v0.12.0...v0.13.0) (2023-07-23)


### Features

* **cmd:** update `proveUnassignedBlocks` flag name ([#315](https://github.com/taikoxyz/taiko-client/issues/315)) ([df640d9](https://github.com/taikoxyz/taiko-client/commit/df640d9d49ceb84268801021ba70fea8e278f39e))
* **driver:** improve `ProcessL1Blocks` for reorg handling ([#325](https://github.com/taikoxyz/taiko-client/issues/325)) ([7272e15](https://github.com/taikoxyz/taiko-client/commit/7272e15650e9ab6aded598e9edcae2659b9d045d))
* **proposer:** add `--txpool.localsOnly` flag ([#326](https://github.com/taikoxyz/taiko-client/issues/326)) ([b292754](https://github.com/taikoxyz/taiko-client/commit/b2927541706e7827dad652140361f4ccf91d1afb))
* **proposer:** handle transaction replacement underpriced error ([#322](https://github.com/taikoxyz/taiko-client/issues/322)) ([2273d10](https://github.com/taikoxyz/taiko-client/commit/2273d105b5dfa6479dc2aa74c16fd0365d06e31a))
* **prover:** add `--oracleProofSubmissionDelay` flag ([#320](https://github.com/taikoxyz/taiko-client/issues/320)) ([85adc04](https://github.com/taikoxyz/taiko-client/commit/85adc04dceabd6218afee72f748e17d69182d81d))
* **prover:** add some prover metrics for Alpha-4 protocol ([#319](https://github.com/taikoxyz/taiko-client/issues/319)) ([d8ff623](https://github.com/taikoxyz/taiko-client/commit/d8ff623a441226c736bd4c52d95df69dd2ce4c86))
* **prover:** flag for proving unassigned proofs or not ([#314](https://github.com/taikoxyz/taiko-client/issues/314)) ([13e6d1d](https://github.com/taikoxyz/taiko-client/commit/13e6d1d87d661c1bdcd9e1537b10b42b33888298))
* **prover:** generate an oracle proof if the incoming proof is incorrect ([#311](https://github.com/taikoxyz/taiko-client/issues/311)) ([003a86b](https://github.com/taikoxyz/taiko-client/commit/003a86bfd3e8f00a4b3c35d048ede6177739a45e))
* **prover:** optimize `skipProofWindowExpiredCheck` check && update `NeedNewProof` check ([#313](https://github.com/taikoxyz/taiko-client/issues/313)) ([b0b4c25](https://github.com/taikoxyz/taiko-client/commit/b0b4c252291ff8d163d2eb71114aa7d63c821c7e))
* **prover:** update `l1Current` cursor to record L1 hash ([#327](https://github.com/taikoxyz/taiko-client/issues/327)) ([4a5adb5](https://github.com/taikoxyz/taiko-client/commit/4a5adb523374008a37831da5febff9a3501a4e81))
* **prover:** update open proving blocks check ([#316](https://github.com/taikoxyz/taiko-client/issues/316)) ([b34930c](https://github.com/taikoxyz/taiko-client/commit/b34930cd4982672bbea962f3706cb83d7e964963))


### Bug Fixes

* **ci:** fix workflow `pnpm install` error ([#321](https://github.com/taikoxyz/taiko-client/issues/321)) ([9eefc8d](https://github.com/taikoxyz/taiko-client/commit/9eefc8d401a35eee1c9b31f5e3c93e18e2754013))
* **prover:** add end height for block filtering if `startHeight` is not nil, and don't block when notifying ([#317](https://github.com/taikoxyz/taiko-client/issues/317)) ([aaec1bb](https://github.com/taikoxyz/taiko-client/commit/aaec1bbdd54df6d60ce39428febbb2747838c31a))
* **prover:** move concurrency guard ([#318](https://github.com/taikoxyz/taiko-client/issues/318)) ([af29c95](https://github.com/taikoxyz/taiko-client/commit/af29c9503def11c373c16555c020307348c5cff6))

## [0.12.0](https://github.com/taikoxyz/taiko-client/compare/v0.11.0...v0.12.0) (2023-07-10)


### Features

* **all:** update bindings && integrate new circuits for L3 ([#290](https://github.com/taikoxyz/taiko-client/issues/290)) ([59469fa](https://github.com/taikoxyz/taiko-client/commit/59469fac2fefe1046d805dc1f19911150e453d87))
* **bindings:** update contract bindings ([#310](https://github.com/taikoxyz/taiko-client/issues/310)) ([021f113](https://github.com/taikoxyz/taiko-client/commit/021f113c2add574843f889b525d55789752b1bd6))
* **prover:** add some prover logs ([#305](https://github.com/taikoxyz/taiko-client/issues/305)) ([e36c76c](https://github.com/taikoxyz/taiko-client/commit/e36c76c7ea6d912477dc8ce61e4639faef00eb5c))
* **prover:** implement staking based tokenomics in client ([#292](https://github.com/taikoxyz/taiko-client/issues/292)) ([7324547](https://github.com/taikoxyz/taiko-client/commit/7324547a80182e93193479089bd334fcce5df7ce))


### Bug Fixes

* **driver:** fix a P2P sync issue ([#298](https://github.com/taikoxyz/taiko-client/issues/298)) ([2ffa052](https://github.com/taikoxyz/taiko-client/commit/2ffa0528110db70f34dd3ef6f48008487caa78a2))
* **prover:** fix a fork choice checking issue ([#309](https://github.com/taikoxyz/taiko-client/issues/309)) ([a393ed8](https://github.com/taikoxyz/taiko-client/commit/a393ed85fed4046039b66bda51bb645ed84d8461))
* **prover:** fix an unlock issue ([#306](https://github.com/taikoxyz/taiko-client/issues/306)) ([392eb78](https://github.com/taikoxyz/taiko-client/commit/392eb78f3721fedea66bd2f361010e2495e385c6))

## [0.11.0](https://github.com/taikoxyz/taiko-client/compare/v0.10.0...v0.11.0) (2023-06-26)


### Features

* **all:** disable no beacon client seen warning  ([#279](https://github.com/taikoxyz/taiko-client/issues/279)) ([cdabcac](https://github.com/taikoxyz/taiko-client/commit/cdabcacb36303667560300775573a4db55fbd5d4))
* **driver:** check the mismatch of last verified block ([#296](https://github.com/taikoxyz/taiko-client/issues/296)) ([79fda87](https://github.com/taikoxyz/taiko-client/commit/79fda8792b29d506b5fa653ed78304d34e892003))
* **driver:** improve error messages ([#289](https://github.com/taikoxyz/taiko-client/issues/289)) ([90e365a](https://github.com/taikoxyz/taiko-client/commit/90e365a79759e0ea701619594b0bf71db4dd3b44))
* **driver:** improve sync progress information ([#288](https://github.com/taikoxyz/taiko-client/issues/288)) ([45d73b9](https://github.com/taikoxyz/taiko-client/commit/45d73b9da34232cf6a3c8636e97aef5854bb86bb))
* **flags:** add retry related flags ([#281](https://github.com/taikoxyz/taiko-client/issues/281)) ([2df4105](https://github.com/taikoxyz/taiko-client/commit/2df4105ab344fb118435b7ef53bcf13ac10f5dc7))
* **metrics:** add `ProverNormalProofRewardGauge` metrics ([#275](https://github.com/taikoxyz/taiko-client/issues/275)) ([cd4e40d](https://github.com/taikoxyz/taiko-client/commit/cd4e40dd477895746843021732a1beba14fa248a))
* **proposer:** add `waitReceiptTimeout` when proposing ([#282](https://github.com/taikoxyz/taiko-client/issues/282)) ([ebf3162](https://github.com/taikoxyz/taiko-client/commit/ebf31623dc491887a25a76da0078559d0b86865c))
* **prover:** improve retry policy for prover ([#280](https://github.com/taikoxyz/taiko-client/issues/280)) ([344bac1](https://github.com/taikoxyz/taiko-client/commit/344bac1435812770c5a1e39efad1545b98d4b106))


### Bug Fixes

* **driver:** fix an issue in `checkLastVerifiedBlockMismatch` ([#297](https://github.com/taikoxyz/taiko-client/issues/297)) ([a68730c](https://github.com/taikoxyz/taiko-client/commit/a68730c0d9cc1b15cdd314ad7939f8971104b362))
* **driver:** fix geth lag to verified block when syncing ([#294](https://github.com/taikoxyz/taiko-client/issues/294)) ([c57f6e8](https://github.com/taikoxyz/taiko-client/commit/c57f6e8ac84ad55c0d51bfae278c88f7694c2265))
* **pkg:** minor fixes for `WaitReceipt` ([#284](https://github.com/taikoxyz/taiko-client/issues/284)) ([feaa2b6](https://github.com/taikoxyz/taiko-client/commit/feaa2b6487e1578c4082ba0b4be087a627512c4b))
* **prover:** ensure L2 reorg finished before generating proofs && add `verificationCheckTicker` ([#277](https://github.com/taikoxyz/taiko-client/issues/277)) ([6fa24ea](https://github.com/taikoxyz/taiko-client/commit/6fa24ea2b4674865dc381098e57a2171c9fce95b))

## [0.10.0](https://github.com/taikoxyz/taiko-client/compare/v0.9.0...v0.10.0) (2023-06-08)


### Features

* **all:** improve proposer && prover logs ([#264](https://github.com/taikoxyz/taiko-client/issues/264)) ([6d0a724](https://github.com/taikoxyz/taiko-client/commit/6d0a7248d78fcd0a73e53a89a21adbeff7f3b61b))
* **driver:** add proof reward metric ([#273](https://github.com/taikoxyz/taiko-client/issues/273)) ([1e00560](https://github.com/taikoxyz/taiko-client/commit/1e00560a1564d61448687ad933fe39a301020bf9))
* **driver:** optimize error handling for `CalldataSyncer` ([#262](https://github.com/taikoxyz/taiko-client/issues/262)) ([580e354](https://github.com/taikoxyz/taiko-client/commit/580e35487b32566761721422bf8d0ca9e5071ed5))
* **pkg:** optimize `WaitL1Origin` ([#267](https://github.com/taikoxyz/taiko-client/issues/267)) ([2d1fda9](https://github.com/taikoxyz/taiko-client/commit/2d1fda90ec54fb25eee789968b9d2177017ace6f))
* **pkg:** update logs when dialing ethclients ([#263](https://github.com/taikoxyz/taiko-client/issues/263)) ([99c980b](https://github.com/taikoxyz/taiko-client/commit/99c980becd0ea2872e6f91b8f422fe66ca8ebfb2))
* **proposer:** add `--maxProposedTxListsPerEpoch` flag ([#258](https://github.com/taikoxyz/taiko-client/issues/258)) ([2cfcf81](https://github.com/taikoxyz/taiko-client/commit/2cfcf814200c2d41d539a427c94fe2a7fefcaf21))
* **prover:** check if a system proof has already been submitted by another system prover ([#274](https://github.com/taikoxyz/taiko-client/issues/274)) ([1fcb244](https://github.com/taikoxyz/taiko-client/commit/1fcb244b29467fcdb7972a724a1ace8b94a67eb8))
* **prover:** improve `onBlockProposed` listener ([#266](https://github.com/taikoxyz/taiko-client/issues/266)) ([5cbdcac](https://github.com/taikoxyz/taiko-client/commit/5cbdcacaa7f902875bb870ea909c7b5ad92220dd))
* **prover:** improve `ZkevmRpcdProducer` logs ([#265](https://github.com/taikoxyz/taiko-client/issues/265)) ([d3fdd94](https://github.com/taikoxyz/taiko-client/commit/d3fdd94f95593567350a86bead5750b12cfd31be))
* **prover:** update proof submission logs ([#261](https://github.com/taikoxyz/taiko-client/issues/261)) ([ea87f7f](https://github.com/taikoxyz/taiko-client/commit/ea87f7f8252073814007d9d54d71cc00171237d7))


### Bug Fixes

* **driver:** fix an issue for P2P sync timeout ([#268](https://github.com/taikoxyz/taiko-client/issues/268)) ([3aee10c](https://github.com/taikoxyz/taiko-client/commit/3aee10c0ba9170eb652e059c51ce029b2af8a3a4))
* **prover:** fix a `targetDelay` calculation issue ([#272](https://github.com/taikoxyz/taiko-client/issues/272)) ([ffcfb53](https://github.com/taikoxyz/taiko-client/commit/ffcfb53e1be7ffe04fdb67ef9a176cc37b7369da))

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
