import sys, os, shutil, subprocess, time, config, sqlite3, oss2


#Config
accessId = ""
accessKey = ""
endpoint = ""
bucketName = ""

def buildPath(rootPath, host, build, _accessId, _accessKey, _endpoint, _bucket):
    config.buildProjectPath(rootPath, host, build)

    global accessId
    accessId = _accessId
    global accessKey
    accessKey = _accessKey
    global endpoint
    endpoint = _endpoint
    global bucketName
    bucketName = _bucket


def processData():
    #Upload to OSS
    auth = oss2.Auth(accessId, accessKey)
    bucket = oss2.Bucket(auth, endpoint, bucketName, connect_timeout=30)
    bucket.put_object_from_file("cs50vn/virustracker-crawler/v1/dev/app.zip", config.genRootDir + os.sep + config.versionName + os.sep + "app.zip")


def main(argv):
    start = time.time()
    print("===========================================================")
    print("                      \033[1;32;40mPROCESS APP DATA\033[0;37;40m")
    print("===========================================================")

    print(str(argv))
    buildPath(argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6])

    processData()

    elapsedTime = time.time() - start
    print("Running time: %s s" % str(elapsedTime))

if __name__ == '__main__':
    main(sys.argv[1:])


