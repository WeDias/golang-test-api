#!/bin/bash

ENV_TYPE=${1:-"dev"}
ENV_PATH="./resources"

if [ $# -eq 0 ]
then
    echo "using default env (dev)"
fi

if test -f "$ENV_PATH/$ENV_TYPE.env"
then
    echo 'setting environment variables'
    echo -n > .env && cat $ENV_PATH/$ENV_TYPE.env >> .env
    echo -n > tests/.env && cat $ENV_PATH/$ENV_TYPE.env >> tests/.env
    echo "env" $ENV_TYPE": successfully configured!"
else
    echo "env" $ENV_TYPE": error configuring environment! check if you chose a valid env."
    echo "available envs: dev, docker."
fi