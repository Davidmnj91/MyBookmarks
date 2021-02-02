from datetime import datetime

from bson import ObjectId

from ..db import Document


class Bookmark(Document):
    def __init__(self):
        super().__init__('bookmarks')

    def serialize(self, data):
        return {
            '_id': str(data['_id']),        
            'url': data['url'],        
            'tags': data['tags'],        
            'note': data['note'],        
            'created': data['created'],        
            'updated': data['updated'],        
        }
