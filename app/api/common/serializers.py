import datetime
from functools import singledispatch

import bson


@singledispatch
def serialize(rv):
    """
    Define a generic serializable function.
    """
    return rv


@serialize.register(datetime.datetime)
def serialize_dt(rv):
    """Register the `datetime.datetime` type
    for the generic serializable function.
    Serialize a `datetime` object to `string`
    according to strict-rfc3339.
    :param rv: object to be serialized
    :type rv: datetetime.datetime
    :returns: string
    """
    return datetime.datetime.strftime(rv, '%Y-%m-%dT%H:%M:%S.%fZ')


@serialize.register(bson.ObjectId)
def serialize_uuid(rv):
    """Register the `ObjectID` type
    for the generic serializable function.
    :param rv: object to be serialized
    :type rv: bson.ObjectID
    :returns: string
    """
    return str(rv)


class ModelSerializerMixin(object):
    """
    Serializable Mixin for the SQLAlchemy objects.
    """
    def to_dict(self, data):
        serialized_model = {}

        for attr in data:
            serialized_model[attr] = self._serialize_attr(data[attr])

        return serialized_model

    def _serialize_attr(self, attr):
        return serialize(attr)
