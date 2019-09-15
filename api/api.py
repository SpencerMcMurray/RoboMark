import res
from flask import Flask
from flask_cors import CORS
from flask_restful import Resource, Api, reqparse
import os

# Setup flask app
app = Flask(__name__)
app.secret_key = os.urandom(24)
CORS(app, resources={r"/api/*": {"origins": "*"}})
api = Api(app)


api.add_resource(res.Users, "/api/users")

api.add_resource(res.Tests, "/api/tests")

api.add_resource(res.Questions, "/api/questions")

if __name__ == '__main__':
    app.run(debug=True, port=8081, host='localhost')
