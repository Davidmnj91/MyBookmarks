from . import backend
from .model import Bookmark


collection = Bookmark()

def create_bookmark(bookmark_data):
    """Create bookmark.
    :param bookmark_data: bookmark information
    :type bookmark_data: dict
    :returns: serialized bookmark object
    :rtype: dict
    """
    bookmark = collection.create_bookmark(bookmark_data)

    return bookmark.to_dict()


def get_bookmark_by_id(bookmark_id):
    """Get bookmark by id.
    :param bookmark_id: id of the bookmark to be retrived
    :type bookmark_id: integer
    :returns: serialized bookmark object
    :rtype: dict
    """
    bookmark = collection.get_bookmark_by_id(bookmark_id)

    return bookmark.to_dict()


def get_all_bookmarks():
    """Get all bookmarks.
    :returns: serialized bookmark objects
    :rtype: list
    """
    bookmarks = collection.get_all_bookmarks()
    return [
        bookmark.to_dict() for bookmark in bookmarks
    ]


def update_bookmark(bookmark_id, bookmark_data):
    """Update bookmark.
    :param bookmark_data: bookmark information
    :type bookmark_data: dict
    :param bookmark_id: id of the bookmark to be updated
    :type bookmark_id: integer
    :returns: serialized bookmark object
    :rtype: dict
    """
    bookmark = backend.update_bookmark(bookmark_id, bookmark_data)

    return bookmark.to_dict()


def delete_bookmark(bookmark_id):
    """Delete bookmark.
    :param bookmark_id: id of the bookmark to be deleted
    :type bookmark_id: integer
    """
    backend.delete_bookmark(bookmark_id)
