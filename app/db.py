from abc import abstractmethod
from datetime import datetime

from bson import ObjectId
from pymongo.mongo_client import MongoClient

from config import Config

mongo = MongoClient(Config.DB_URI)
db = mongo[Config.DB_NAME]

class Document(object):

    def __init__(self, schema_name) -> None:
        self.document = db[schema_name]

    def insert(self, element):
        element["created"] = datetime.now()
        element["updated"] = datetime.now()
        inserted = self.document.insert_one(element)

        return self.find_by_id(inserted.inserted_id)

    def find(self, criteria={}, projection=None, sort=None, limit=0):
        if "_id" in criteria:
            criteria["_id"] = ObjectId(criteria["_id"])

        found = self.document.find()

        return [self.serialize(data) for data in found]

    def find_by_id(self, id):
        found = self.document.find_one({"_id": ObjectId(id)})
        
        if found is None:
            return not found
        
        return self.serialize(found)

    def update(self, id, element):
        criteria = {"_id": ObjectId(id)}

        element["updated"] = datetime.now()
        set_obj = {"$set": element}  

        updated = self.document.update_one(criteria, set_obj)
        if updated.matched_count == 1:
            return self.serialize(updated.raw_result)

    def delete(self, id):
        deleted = self.document.delete_one({"_id": ObjectId(id)})
        return bool(deleted.deleted_count)

    @abstractmethod
    def serialize(self, data):
        pass
