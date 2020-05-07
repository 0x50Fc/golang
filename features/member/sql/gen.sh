#!/bin/bash

HOME=`pwd`
MS_DIR=$HOME/../../features
GEN_CLI="node $HOME/../../../gen/main.js"

TARGET=$1
FROM=$2
if [ "$FROM" = "" ]; then
    FROM="config"
fi
TS_DIR=$HOME/../ts

$GEN_CLI sql $TS_DIR $HOME/$TARGET $HOME/$FROM.json