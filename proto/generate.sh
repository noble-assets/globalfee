cd proto
buf generate --template buf.gen.gogo.yaml
buf generate --template buf.gen.pulsar.yaml
cd ..

cp -r github.com/noble-assets/globalfee/* ./
cp -r api/noble/globalfee/* api/

rm -rf github.com
rm -rf api/noble
rm -rf noble
