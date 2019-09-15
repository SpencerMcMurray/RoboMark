from .database import fetch_all, insert, other_op
from flask import jsonify


def args_to_where_suffix(args):
    """(list of tuple of obj, str) -> str, tuple of obj
    Makes a suffix for a select query which finds what we ask for
    """
    q = ""
    valid_args = []
    for arg in args:
        if arg[0] is not None:
            q += (" WHERE" if len(q) == 0 else " AND")
            q += f" `{arg[1]}`=%s"
            valid_args.append(arg[0])
    return (q, tuple(valid_args))


def args_to_insert_query(args):
    """(list of tuple of obj, str) -> str
    Makes a suffix for an insert query which inserts the data desired
    REQ: len(args) > 0
    """
    names = ""
    for arg in args:
        names += f"`{arg[1]}`, "
    return names[:-2]


def get_helper(table, args):
    """(str, list of tuple of obj & str) -> list of dict of str:obj
    Performs a GET request that will be used to fetch data
    """
    suffix, valid_args = args_to_where_suffix(args)
    q = f"""SELECT * FROM `{table}`{suffix}"""
    res = fetch_all(q) if len(
        valid_args) == 0 else fetch_all(q, *valid_args)
    return res


def post_helper(table, args):
    """(str, list of tuple of obj & str) -> json
    Performs a POST request that will be used to insert into a table
    """
    # Error check args
    for arg in args:
        if arg[0] is None:
            return jsonify({"outcome": "error: lacking args"})
    names = args_to_insert_query(args)
    items = "%s, " * len(args)
    params = tuple(arg[0] for arg in args)
    q = f"""INSERT INTO `{table}` ({names}) VALUES ({items[:-2]})"""
    return jsonify({"outcome": "success", "id": insert(q, params)})


def delete_helper(table, args):
    """(str, list of tuple of obj & str) -> json
    Performs a DELETE that will be used to delete from a table
    REQ: args[0] must be the id, it's hardcoded in api resources
    """
    # Error check args
    suffix, valid_args = args_to_where_suffix(args)
    if len(valid_args) == 0:
        return jsonify({"outcome": "error: lacking args"})
    deletions = get_helper(table, args)
    q = f"""DELETE FROM `{table}`{suffix}"""
    other_op(q, *valid_args)
    return jsonify({"outcome": "success", "deletions": deletions})


def patch_helper(table, item_id, new_val, field):
    """(str, int, obj, str) -> json
    Performs a PATCH which will be used to update an item in a table
    """
    if item_id is None or new_val is None or field is None:
        return jsonify({"outcome": "error: lacking args"})
    q = f"""UPDATE {table} SET `{field}`=%s WHERE `id`=%s"""
    before = get_helper(table, [(item_id, "id")])
    other_op(q, new_val, item_id)
    after = get_helper(table, [(item_id, "id")])
    return jsonify({"outcome": "success", "before": before, "after": after})
