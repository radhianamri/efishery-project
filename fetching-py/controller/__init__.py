from sanic.response import json, stream


async def unauthorized():
    return json({"success": False, "status": 403,'msg': 'Unauthorized Access'})

async def bad():
    return json({"success": False, "status": 400,'msg': 'Bad Request'})

async def data(data):
    return json({"success": True, "status": 200, 'msg': 'OK', 'data': data})

async def internal_error():
    return json({"success": False, "status": 500, 'msg': 'Internal Server Error'})

async def timeout():
    return json({"success": False, "status": 408, 'msg': 'Timeout server Error'})