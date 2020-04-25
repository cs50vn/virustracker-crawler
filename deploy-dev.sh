echo "Process artifacts"
echo "======================="

export BUILD_TYPE=all
export DEPLOY_SERVER_URL=$1
export DEPLOY_SERVER_USER=$2
export DEPLOY_SERVER_KEY=$3
export OSS_TEST_APP_URL=$4

echo "key"
echo "${DEPLOY_SERVER_KEY}" > key.pem
chmod 600 key.pem
ls -l -a

#python3 scripts/deploy-dev.py $PWD linux $BUILD_TYPE $DEPLOY_SERVER_URL $DEPLOY_SERVER_USER "${DEPLOY_SERVER_KEY}" $OSS_TEST_APP_URL
