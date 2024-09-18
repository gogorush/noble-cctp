cd proto
buf generate
cd ..

cp -r github.com/gogorush/noble-cctp/* ./
rm -rf github.com
