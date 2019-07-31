#!/bin/bash

echo "start gen swagger doc"

if [ ! -d docs ]; then
  mkdir docs
fi

path="./docs/swagger.json"

swagger generate spec -o  $path
echo "finish gen swagger doc"


local_host=$(sed -n 's/"host":\s*"\(.*\)\s*|\s*\(.*\)\s*",/\1/p' $path)
server_host=$(sed -n 's/"host":\s*"\(.*\)\s*|\s*\(.*\)\s*",/\2/p' $path)

if  [[ ! -n "$local_host" ]]; then
    echo "local host: $local_host 配置错误"
    exit 1;
elif  [[ ! -n "$server_host" ]]; then
    echo "server host: $server_host 配置错误"
    exit 1;
fi

host=$local_host

while getopts "h:" opt
do
    case $opt in
        h)
        if [[ $OPTARG = local ]]; then
            host=$local_host
        elif [[ $OPTARG = server ]]; then
            host=$server_host
        else
            echo "参数无效"
            exit 1;
        fi
        ;;
        ?)
        echo "未知参数"
        exit 1;;
    esac
done

# 去除首行空格
host=$(echo $host | sed -n 's/^[ \t]*//p')

echo "API Host: $host"

# 替换Host
sed -i "s?\(\"host\"\:\s*\)\(.*\)?\1\"$host\",?g" $path

if [[ $3 = serve ]]; then
  echo "start swagger serve"
  swagger serve -F=swagger $path
fi