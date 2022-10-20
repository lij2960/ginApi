#!/bin/bash

# 定义app名称
appName=ginApi

# shellcheck disable=SC1019
# shellcheck disable=SC1072
# shellcheck disable=SC1073
# shellcheck disable=SC1035
if [ ! -e "./pack/{$appName}" ]; then
  mkdir -p ./pack/$appName
fi

go build -o ./pack/$appName/$appName ./

# shellcheck disable=SC2225
cp -a ./config ./pack/$appName

# 打包项目文件
# shellcheck disable=SC2164
cd ./pack
tar zcf $appName.tar.gz ./$appName
rm -rf ./$appName
