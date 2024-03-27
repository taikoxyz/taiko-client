#!/bin/bash

cp -f .env.tests .env.test
cat ./integration_test/.env >> .env.test