
#!/bin/bash

set -e

VERSION=$1
APP_NAME=$2

github-release release --user larse514 \
    --repo serverlessui --tag $VERSION \
    --name "minor release $VERSION" \
    --description "minor release for version $VERSION" \
    --pre-release

for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        github-release upload --user larse514 \
            --repo serverlessui --tag $VERSION \
            --name "$APP_NAME-$GOOS-$GOARCH" \
            --file $APP_NAME-$GOOS-$GOARCH
    done
done
