
from sanic.response import json, stream
from sanic.request import Request
from sanic import Blueprint
from functools import wraps

bp_fetch = Blueprint('fetch', url_prefix='/fetch', strict_slashes=True)


def authorized():
    def decorator(f):
        @wraps(f)
        async def decorated_function(request, *args, **kwargs):

            response = await f(request, *args, **kwargs)
            return response
        return decorated_function
    return decorator

@bp_fetch.get('/resource')
@authorized()
async def resource(request: Request):
    return json({"success": True, "status": 200,'message': 'pong'})


@bp_fetch.get('/resource/aggregate')
async def resource_aggregate(request: Request):
    return json({"success": True, "status": 200,'message': 'pong'})

@bp_fetch.get('/claims')
async def claims(request: Request):
    return json({"success": True, "status": 200,'message': 'pong'})