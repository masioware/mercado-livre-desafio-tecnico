from flask import Flask

from app.controllers import distribution_center


def create_app():
    app = Flask(__name__)

    distribution_center.start(app)

    return app
