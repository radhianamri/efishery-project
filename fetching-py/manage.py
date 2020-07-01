from config import Init, config
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
    Init(deployment_type)
    print(config)

    app = create_app(deployment_type)
    app.run(host="0.0.0.0", port=6000, debug=True, workers=1)