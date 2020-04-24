import sys, os, shutil, subprocess, time, config


def buildWorker():
    print("===========================================================")
    print("                      \033[1;32;40mBUILD WORKER\033[0;37;40m")
    print("===========================================================")
    
    os.chdir(config.srcDirWorker)

    if config.buildType == "release":
        print("Release build")
    else:
        if config.hostType == "linux":
            os.environ["GOOS"] = "linux"
            cmd = "go build -ldflags \"-s -w\" -o %s main.go" % (config.genRootDir + os.sep + config.outputWorkerFile)
            print(cmd)
            subprocess.call(cmd, shell=True)
            cmd = "chmod 740 %s" % (config.genRootDir + os.sep + config.outputWorkerFile)
            subprocess.call(cmd, shell=True)
        else:
            os.environ["GOOS"] = "windows"
            cmd = "go build -ldflags \"-s -w\" -o %s.exe main.go" % (config.genRootDir + os.sep + config.outputWorkerFile)
            print(cmd)
            subprocess.call(cmd, shell=True)

    os.chdir(config.rootDir)


def buildAPI():
    print("===========================================================")
    print("                      \033[1;32;40mBUILD API\033[0;37;40m")
    print("===========================================================")

    os.chdir(config.srcDirAPI)
    
    if config.buildType == "release":
        print("Release build")
        #cmd = "upx %s" % (config.genRootDir + os.sep + config.outputAPIFile)
        #subprocess.call(cmd, shell=True)

        #cmd = "upx %s.exe" % (config.genRootDir + os.sep + config.outputAPIFile)
        #subprocess.call(cmd, shell=True)
    else:
        if config.hostType == "linux":
            os.environ["GOOS"] = "linux"
            cmd = "go build -ldflags \"-s -w\" -o %s main.go" % (config.genRootDir + os.sep + config.outputAPIFile)
            print(cmd)
            subprocess.call(cmd, shell=True)

            cmd = "chmod 740 %s" % (config.genRootDir + os.sep + config.outputAPIFile)
            subprocess.call(cmd, shell=True)
        else:
            os.environ["GOOS"] = "windows"
            cmd = "go build -ldflags \"-s -w\" -o %s.exe main.go" % (config.genRootDir + os.sep + config.outputAPIFile)
            print(cmd)
            subprocess.call(cmd, shell=True)


    os.chdir(config.rootDir)


def buildGoProgram():
    print("===========================================================")
    print("                      \033[1;32;40mBUILD GO\033[0;37;40m")
    print("===========================================================")
    
    buildWorker()
    buildAPI()

def buildPackage():
    print("===========================================================")
    print("                      \033[1;32;40mBUILD PACKAGE\033[0;37;40m")
    print("===========================================================")

    if config.buildType == "release":
        print("Release build")
        #cmd = "upx %s" % (config.genRootDir + os.sep + config.outputAPIFile)
        #subprocess.call(cmd, shell=True)

        #cmd = "upx %s.exe" % (config.genRootDir + os.sep + config.outputAPIFile)
        #subprocess.call(cmd, shell=True)
    else:
        if config.hostType == "windows":
            src = config.genRootDir + os.sep + config.outputAPIFile + ".exe"
            des = config.genAppDir

            print("\033[1;34;40mFrom:\n\033[0;37;40m" + src)
            print("\033[1;34;40mTo\n\033[0;37;40m" + des)

            if not os.path.exists(des):
                os.makedirs(des, exist_ok=True)
            shutil.copy(src, des)

            src = config.genRootDir + os.sep + config.outputWorkerFile + ".exe"
            shutil.copy(src, des)

            src = config.templateDir + os.sep + "config.json"
            shutil.copy(src, des)

            src = config.dataDir + os.sep + "%s-template.db" % config.appName
            shutil.copyfile(src, des + os.sep + "%s.db" % config.appName)

            cmd = "7z a %s.zip %s" % (des, des)
            subprocess.call(cmd, shell=True)

        else:
            src = config.genRootDir + os.sep + config.outputAPIFile
            des = config.genAppDir

            print("\033[1;34;40mFrom:\n\033[0;37;40m" + src)
            print("\033[1;34;40mTo\n\033[0;37;40m" + des)

            if not os.path.exists(des):
                os.makedirs(des, exist_ok=True)
            shutil.copy(src, des)

            src = config.genRootDir + os.sep + config.outputWorkerFile
            shutil.copy(src, des)

            src = config.templateDir + os.sep + "install.sh"
            shutil.copy(src, des)

            src = config.templateDir + os.sep + "uninstall.sh"
            shutil.copy(src, des)

            src = config.templateDir + os.sep + "config.json"
            shutil.copy(src, des)

            src = config.templateDir + os.sep + "%s.service" % config.workerName
            shutil.copy(src, des)

            src = config.templateDir + os.sep + "%s.service" % config.apiName
            shutil.copy(src, des)

            src = config.dataDir + os.sep + "%s-template.db" % config.appName
            shutil.copyfile(src, des + os.sep + "%s.db" % config.appName)

            cmd = "7z a %s.zip %s" % (des, des)
            subprocess.call(cmd, shell=True)

            print("")
    print("\n")


def main(argv):
    start = time.time()
    print("===========================================================")
    print("                      \033[1;32;40mBUILD APPLICATION\033[0;37;40m")
    print("===========================================================")

    print(str(argv))
    config.buildProjectPath(argv[0], argv[1], argv[2])

    buildGoProgram()

    buildPackage()

    elapsedTime = time.time() - start
    print("Running time: %s s" % str(elapsedTime))


if __name__ == '__main__':
    main(sys.argv[1:])
