#!/bin/bash
binary_name=poems
goos=$(uname)
version=0.0.2
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  goos=linux64
elif [[ "$OSTYPE" == "darwin"* ]]; then
  goos=mac64
elif [[ "$OSTYPE" == "msys" ]]; then
  goos=win64
else
  echo "Error: The current os is not supported at this time" 1>&2
  exit 1
fi


file_name=poems-${goos}-${version}.tar.gz

url=https://github.com/yanzhoupan/poems/releases/download/${version}/${file_name}
echo "Download url:${url}"

curl "$url" -OL # --progress --retry 2 2>&1

# shellcheck disable=SC2181
if [[ $? -ne 0 ]]; then
  echo "curl failed"
  exit 1
fi

tar -xvf "${file_name}"

if [[ "$OSTYPE" == "msys" ]]; then
  echo "poems download success, run 'poems.exe' in current directory now!"
  del $file_name
  exit 0
fi
sudo cp $binary_name /usr/local/bin/

rm -rf $binary_name
rm -rf $file_name

echo "poems install success, run 'poems' now!"