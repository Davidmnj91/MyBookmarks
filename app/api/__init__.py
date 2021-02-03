from flask import Flask

from .bookmarks import bp as bookmarks_bp
from .common.middleware import (after_request_middleware,
                                before_request_middleware, response,
                                teardown_appcontext_middleware)


def create_app():
    app = Flask(__name__)

    app.register_blueprint(bookmarks_bp)

    # register before request middleware
    before_request_middleware(app=app)

    # register after request middleware
    after_request_middleware(app=app)

    # register after app context teardown middleware
    teardown_appcontext_middleware(app=app)

    # register custom error handler
    response.json_error_handler(app=app)

    return app
