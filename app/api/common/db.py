from datetime import datetime

from app.config import Config
from bson import ObjectId
from pymongo.mongo_client import MongoClient

from .serializers import ModelSerializerMixin

mongo = MongoClient(Config.DB_URI)
db = mongo[Config.DB_NAME]

class Document(ModelSerializerMixin):

    def __init__(self, schema_name) -> None:
        self.document = db[schema_name]

    def insert(self, element):
        element["created"] = datetime.now()
        element["updated"] = datetime.now()
        inserted = self.document.insert_one(element)

        self = self.find_by_id(inserted.inserted_id)
        
        return self

    def find(self, criteria={}, projection=None, sort=None, limit=0):
        if "_id" in criteria:
            criteria["_id"] = ObjectId(criteria["_id"])

        found = self.document.find()

        collection = [self.to_dict(data) for data in found]
        return collection

    def find_by_id(self, id):
        found = self.document.find_one({"_id": ObjectId(id)})
        
        if found is None:
            return not found
        
        return self.to_dict(found)

    def update(self, id, element):
        criteria = {"_id": ObjectId(id)}

        element["updated"] = datetime.now()
        set_obj = {"$set": element}  

        updated = self.document.update_one(criteria, set_obj)
        if updated.matched_count == 1:
            return self.find_by_id(id)
        else:
            return not updated

    def delete(self, id):
        deleted = self.document.delete_one({"_id": ObjectId(id)})
        return bool(deleted.deleted_count)
