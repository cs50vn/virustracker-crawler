import sys, os, shutil, subprocess, time, config

# App config
serverUrl = ""
serverUser = ""
serverkey = ""
appUrl = ""

def buildPath(rootPath, host, build, _serverUrl, _serverUser, _appUrl):
    config.buildProjectPath(rootPath, host, build)

    global serverUrl
    serverUrl = _serverUrl
    global serverUser
    serverUser = _serverUser
    #global serverkey
    #serverkey = _serverkey
    global appUrl
    appUrl = _appUrl

def deployApp():

    cmd = '''curl -o %s %s;
        ls -l;
        unzip -q ./%s;
        cd app;
        bash ./uninstall.sh;
        bash ./install.sh;
        rm -f app.zip;
    ''' % ("app.zip", appUrl, "app.zip")
    print(cmd)

    cmd1 = '''
        ssh -i key.pem -o StrictHostKeyChecking=no %s@%s %s
    ''' % (serverUser, serverUrl, cmd)
    subprocess.call(cmd1, shell=True)


def main(argv):
    start = time.time()
    print("===========================================================")
    print("                      \033[1;32;40mDEPLOY APP TO TEST SERVER\033[0;37;40m")
    print("===========================================================")

    print(str(argv))
    buildPath(argv[0], argv[1], argv[2], argv[3], argv[4], argv[5])

    deployApp()

    elapsedTime = time.time() - start
    print("Running time: %s s" % str(elapsedTime))

if __name__ == '__main__':
    main(sys.argv[1:])
