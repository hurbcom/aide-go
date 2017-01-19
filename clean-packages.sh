#!/bin/bash
array=($(grep -oP '(\w[\w\.\-]+\/([\w\.\-]+|v\d{2,2})(\/[\w\-]+)*)' Godeps/Godeps.json))
array_len=${#array[@]}
echo 'Encontrados '$array_len' pacotes'
for ((i=1; i<${array_len}; i++)); do
    IFS='/' read -ra ADDR <<< ${array[$i]}
    repo=${ADDR[0]}'/'${ADDR[1]}
    if [ ${ADDR[2]} ]; then
        repo=$repo'/'${ADDR[2]}
    fi
    echo 'Removing '$i' '$repo
    echo $(go clean -i $repo)
    echo $(trash ${GOPATH}/src/$repo)
done
