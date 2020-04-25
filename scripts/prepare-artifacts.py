import sys, os, shutil, subprocess, time, config, sqlite3, oss2


#Config
configUrl = ""
databaseUrl = ""


def buildPath(rootPath, host, build, _configUrl, _databaseUrl):
    config.buildProjectPath(rootPath, host, build)

    global configUrl
    configUrl = _configUrl
    global databaseUrl
    databaseUrl = _databaseUrl


def prepareData():
    #Rename app dir to "app"
    src = config.genAppDir
    des = config.genRootDir + os.sep + config.versionName + os.sep + "app"
    shutil.move(src, des)

    #Download config and test db
    cmd = "curl %s -o %sconfig.json" % ( configUrl, (des + os.sep))
    print(cmd)
    subprocess.call(cmd, shell=True)

    cmd = "curl %s -o %svirustracker-crawler.db" % ( databaseUrl, (des + os.sep))
    print(cmd)
    subprocess.call(cmd, shell=True)

    #Zip folder to app.zip
    cmd = "7z a %s.zip %s" % (des, des)
    subprocess.call(cmd, shell=True)

def main(argv):
    start = time.time()
    print("===========================================================")
    print("                      \033[1;32;40mPREPARE APP DATA\033[0;37;40m")
    print("===========================================================")

    print(str(argv))
    buildPath(argv[0], argv[1], argv[2], argv[3], argv[4])

    prepareData()

    elapsedTime = time.time() - start
    print("Running time: %s s" % str(elapsedTime))

if __name__ == '__main__':
    main(sys.argv[1:])
