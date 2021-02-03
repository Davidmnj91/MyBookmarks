from ..common.db import Document
from ..common.serializers import ModelSerializerMixin


class Bookmark(Document, ModelSerializerMixin):
    def __init__(self):
        super().__init__('bookmarks')
