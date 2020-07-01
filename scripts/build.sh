# only supposed to be run from Make file
NAME="$1"
$(go env GOPATH)/bin/qtdeploy build linux "./internal/cmd/$NAME"
mkdir -p bin
cp "./internal/cmd/$NAME/deploy/linux/$NAME" ./bin/
rm -rf "./internal/cmd/$NAME/linux" "./internal/cmd/$NAME/deploy"
