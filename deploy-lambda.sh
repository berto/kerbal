set -x
GOOS=linux go build
if [ $? -ne 0 ]; then
    echo ERROR: failed to build go binary
    exit 1
fi

rm lambda.zip
zip lambda.zip ./kerbal
aws  --profile kerbal.me lambda e --function-name kerbal --zip-file fileb://lambda.zip

rm lambda.zip   
rm kerbal