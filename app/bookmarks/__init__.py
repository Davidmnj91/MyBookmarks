from flask import Blueprint, jsonify, request

from .model import Bookmark

BP_NAME = 'bookmarks'
bp = Blueprint(BP_NAME, '__name__', url_prefix='/' + BP_NAME)

bookmarkSchema = Bookmark()

@bp.route('', methods=['GET'])
def get_bookmarks():
    return jsonify(bookmarkSchema.find())

@bp.route('', methods=['POST'])
def create_bookmark():
    return jsonify(bookmarkSchema.insert(request.json))

@bp.route('/<id>', methods=['GET'])
def get_bookmark(id):
    return jsonify(bookmarkSchema.find_by_id(id))

@bp.route('/<id>', methods=['DELETE'])
def delete_bookmark(id):
    return jsonify(bookmarkSchema.delete(id))
