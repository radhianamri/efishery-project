

from sanic import Sanic
# from . import models
import configparser, os

def create_app(deployment_type):

    app = Sanic(__name__, strict_slashes=True)
    
    from controller.fetch import bp_fetch
    app.blueprint(bp_fetch)

    from sanic_cors import CORS, cross_origin


    cors_kwargs = {
            'automatic_options': True,
            'supports_credentials': True,
            'allow_headers': '*',
            "origins": "*"
        }
    CORS(app, **cors_kwargs)

    from sanic.request import Request
    from sanic import response
    @app.get('/ping')
    async def ping(request: Request):
        return response.json({"success": True, "status": 200,'message': 'pong'})


    return app