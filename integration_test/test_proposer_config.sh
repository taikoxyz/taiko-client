#!/bin/bash

cp -f .env.tests .env.proposer.test
cat ./integration_test/.env >> .env.proposer.test