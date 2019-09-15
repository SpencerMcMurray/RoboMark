from flask_restful import Resource, reqparse
from flask import jsonify, request
from .helpers import get_helper, post_helper, patch_helper

PARAMS = ['user', 'test', 'question', 'answers']

parser = reqparse.RequestParser()
for p in PARAMS:
    parser.add_argument(p)


class Questions(Resource):
    def get(self):
        args = request.args
        args = [(args.get('id'), 'id')] + [(args.get(p), p) for p in PARAMS]
        questions = get_helper("questions", args)
        return jsonify({"questions": questions})

    def post(self):
        args = parser.parse_args()
        args = [(args.get(p), p) for p in PARAMS]
        return post_helper("questions", args)
