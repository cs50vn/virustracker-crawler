echo "Process artifacts"
echo "======================="

export BUILD_TYPE=all
export DEV_SERVER_URL=$1
export DEV_SERVER_USER=$2
export DEV_SERVER_KEY=$3
export OSS_TEST_APP_URL=$4

echo "key"
echo "${DEV_SERVER_KEY}"


python3 scripts/deploy-dev.py $PWD linux $BUILD_TYPE $DEV_SERVER_URL $DEV_SERVER_USER "${DEV_SERVER_KEY}" $OSS_TEST_APP_URL
