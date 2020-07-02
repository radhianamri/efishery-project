
from sanic.response import json, stream
from sanic.request import Request
from functools import wraps
import aiohttp, ujson, jwt
from datetime import datetime as dt 
from . import *


from expiringdict import ExpiringDict
cache = ExpiringDict(max_len=1, max_age_seconds=300)

from sanic import Blueprint
bp_fetch = Blueprint('fetch', url_prefix='/fetch', strict_slashes=True)

auth_scheme = "Bearer"

def authorized():
    def decorator(f):
        @wraps(f)
        async def decorated_function(request, *args, **kwargs):
            auth = request.headers.get("Authorization")
            l  = len(auth_scheme)
            if len(auth) > l+1 and auth[:l] == auth_scheme:
                userToken = auth[l+1:]
            else:
                return unauthorized()

            if userToken == "":
                return unauthorized()

            res = jwt.decode(userToken, request.app.conf["jwt_secret"], algorithms=['HS256'])
            request["claims"] = res
            response = await f(request, *args, **kwargs)

            return response
        return decorated_function
    return decorator

@bp_fetch.get('/resource')
@authorized()
async def resource(request: Request):
    resource = []
    conversion_rate = 0
    try:
        async with aiohttp.ClientSession(timeout=aiohttp.ClientTimeout(total=5)) as session:
            async with session.get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list") as resp:
                content = await resp.read()
                if resp.status == 200:
                    resource = ujson.loads(content)
                else:
                    print(resp)
                    return internal_error()
            if cache.get("IDR_USD") == None:
                async with session.get("https://free.currconv.com/api/v7/convert?apiKey=f1ed7b91c18de876f08d&q=IDR_USD&compact=ultra") as resp:
                    res = await resp.read()
                    if resp.status == 200:
                        conversion_rate = ujson.loads(res)["IDR_USD"]
                        cache["IDR_USD"] = conversion_rate
                    else:
                        print(resp)
                        return internal_error()
            else:
                conversion_rate = cache["IDR_USD"]
    except Exception as e:
        print(e)     
        return timeout()

   
    res = []
    for row in (resource):
        row["price_usd"] = (float(row["price"]) if row["price"] != None else 0) * conversion_rate
        res.append(row)

    return data(res)


@bp_fetch.get('/resource/aggregate')
@authorized()
async def resource_aggregate(request: Request):
    resource = []
    try:
        async with aiohttp.ClientSession(timeout=aiohttp.ClientTimeout(total=5)) as session:
            async with session.get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list") as resp:
                content = await resp.read()
                if resp.status == 200:
                    resource = ujson.loads(content)
                else:
                    print(resp)
                    return internal_error()
    except Exception as e:
        print(e) 
        return timeout()
   
    res = []
    dates = []
    for row in (resource):
        if row.timestamp:
            if row.get("tgl_parsed") not in dates:
                dates.append(row.get("tgl_parsed"))

    print(dates)
    return data(res)

@bp_fetch.get('/claims')
@authorized()
async def claims(request: Request):
    return data(request["claims"])
