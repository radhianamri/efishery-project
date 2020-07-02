import toml, ujson

def parse_config(deployment_type):
    config = toml.load('config/config.toml')
    return config[deployment_type]
