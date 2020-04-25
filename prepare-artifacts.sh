echo "Prepare artifacts"
echo "======================="

export BUILD_TYPE=release
export OSS_TEST_CONFIG_URL=$1
export OSS_TEST_DATABASE_URL=$2

python3 scripts/prepare-artifacts.py $PWD linux $BUILD_TYPE $OSS_TEST_CONFIG_URL $OSS_TEST_DATABASE_URL


