echo "Process artifacts"
echo "======================="

export BUILD_TYPE=all
export OSS_ACCESS_KEY_ID=$1
export OSS_ACCESS_KEY_SECRET=$2
export OSS_ENDPOINT=$3
export OSS_BUCKET=$4

python3 scripts/process-artifacts.py $PWD linux $BUILD_TYPE $OSS_ACCESS_KEY_ID $OSS_ACCESS_KEY_SECRET $OSS_ENDPOINT $OSS_BUCKET


