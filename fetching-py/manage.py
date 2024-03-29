from config import parse_config
import argparse
from server import create_app

if __name__ == '__main__':
    DEPLOYMENT_TYPES = ['LOCALHOST', 'PRODUCTION']
    parser = argparse.ArgumentParser()
    parser.add_argument("--deployment_type", help="Type of deployment")
    deployment_type = parser.parse_args().deployment_type

    if deployment_type == None:
        deployment_type = "LOCALHOST"

    deployment_type = deployment_type.upper() 
    if deployment_type not in DEPLOYMENT_TYPES:
        raise ValueError(
            "'deployment_type' should be one of the followings: {}".format(
                ', '.join(DEPLOYMENT_TYPES)
            )
        )
    

    app = create_app(deployment_type, parse_config(deployment_type))
    app.run(host=app.conf["api_route_host"], port=app.conf["api_route_port"], debug=True, workers=1)