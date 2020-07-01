import toml, ujson

config = toml.load('config/config.toml')
def Init(deployment_type):
    global config
    print(deployment_type)
    print(config["LOCALHOST"])
    print(config)
    config = config["LOCALHOST"]
