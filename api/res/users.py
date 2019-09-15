from flask_restful import Resource, reqparse
from flask import jsonify, request
from .helpers import get_helper, post_helper

PARAMS = ['isTeacher', 'name']

parser = reqparse.RequestParser()
for p in PARAMS:
    parser.add_argument(p)


class Users(Resource):
    def get(self):
        args = request.args
        args = [(args.get('id'), 'id')] + [(args.get(p), p) for p in PARAMS]
        users = get_helper("users", args)
        return jsonify({"users": users})

    def post(self):
        args = parser.parse_args()
        args = [(args.get(p), p) for p in PARAMS]
        return post_helper("users", args)
