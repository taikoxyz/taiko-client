# How to debug test cases?
* start docker compose
```
./docker/start.sh
```

* deploy L1 contracts
```
# TAIKO_MONO_DIR variable is taiko-mono repo path.
TAIKO_MONO_DIR=$taiko-mono ./integration_test/deploy_l1_contract.sh
```

* show environment variables
```
# TAIKO_MONO_DIR variable is taiko-mono repo path.
TAIKO_MONO_DIR=$taiko-mono ./integration_test/test_env.sh
```

* copy that the result of previous step and release them in `Debug configurations`
> after debugged don't forget stop docker compose
```
./docker/stop.sh
```