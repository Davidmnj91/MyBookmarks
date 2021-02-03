from flask import Blueprint, jsonify, request

from ..common.validation import schema
from .model import Bookmark

BP_NAME = 'bookmarks'
bp = Blueprint(BP_NAME, '__name__', url_prefix='/' + BP_NAME)

collection = Bookmark()

@bp.route('', methods=['GET'])
def get_bookmarks():
    return jsonify(collection.find())

@bp.route('', methods=['POST'])
@schema('create_bookmark')
def create_bookmark():
    return collection.insert(request.json)

@bp.route('/<id>', methods=['PUT'])
@schema('update_bookmark')
def update_bookmark(id):
    return collection.update(id, request.json)

@bp.route('/<id>', methods=['GET'])
def get_bookmark(id):
    return collection.find_by_id(id)

@bp.route('/<id>', methods=['DELETE'])
def delete_bookmark(id):
    return collection.delete(id)
