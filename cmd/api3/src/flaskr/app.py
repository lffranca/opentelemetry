from flask import Flask

from src.flaskr.ext import configuration
from opentelemetry.instrumentation.flask import FlaskInstrumentor


def minimal_app(**config):
    flask_app = Flask(__name__)
    configuration.init_app(flask_app, **config)
    return flask_app


def create_app(**config):
    app = minimal_app(**config)
    configuration.load_extensions(app)
    FlaskInstrumentor().instrument_app(app)
    return app


if __name__ == "__main__":
    create_app().run(port=5001)
