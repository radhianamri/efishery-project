from sanic.response import json, stream


def unauthorized():
    return  json({"success": False, "status": 403,'msg': 'Unauthorized Access'})

def bad():
    return  json({"success": False, "status": 400,'msg': 'Bad Request'})

def data(data):
    return  json({"success": True, "status": 200, 'msg': 'OK', 'data': data})

def internal_error():
    return  json({"success": False, "status": 500, 'msg': 'Internal Server Error'})

def timeout():
    return   json({"success": False, "status": 408, 'msg': 'Timeout server Error'})