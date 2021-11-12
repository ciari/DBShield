#!/bin/bash
pkgname=dbshield
orig_name=DBShield
git_user=ciari

app_dir=/usr/src/app/
env_path=$app_dir/venv
cwd=$(pwd)
deb_dir=${cwd}/debian
tmp_dir=${deb_dir}/${pkgname}-tmp

echo "Set GO Env"
export PATH=$PATH:/usr/local/go/bin
export GOPATH=/tmp/${pkgname}-build/
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
mkdir -p $GOPATH
mkdir -p ${tmp_dir}/bin/
GO111MODULE=on go get -u github.com/${git_user}/${orig_name}
cp $GOPATH/bin/${orig_name} ${tmp_dir}/bin/${pkgname}
echo "Copying binary to: "${tmp_dir}/bin/${pkgname}
rm -rf $GOPATH
echo ${orig_name} builded
