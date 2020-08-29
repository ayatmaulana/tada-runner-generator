#!/usr/bin/env bash

package_name='tada-runner-generator'
platforms=("linux/amd64" "linux/386" "darwin/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name='./bin/'$package_name'-'$GOOS'-'$GOARCH

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    else
        echo 'Success building binary for' $GOOS '-' $GOARCH
    fi
done