import os


def getPort():
    variables = {}
    fileName = os.path.dirname(os.path.realpath(__file__)) + \
        "/../.env.test"
    with open(fileName) as envFile:
        for line in envFile:
            name, var = line.partition("=")[::2]
            variables[name.strip()] = var
        return variables["FRONTEND_SERVER_PORT"]
