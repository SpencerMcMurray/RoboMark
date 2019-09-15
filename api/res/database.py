import pymysql


class Database:
    def __init__(self):
        host = "den1.mysql6.gear.host"
        port = 3306
        user = "robomark"
        password = "Gw3CVo~1Z4d~"
        db = "robomark"

        self.con = pymysql.connect(host=host, port=port, user=user, password=password, db=db,
                                   cursorclass=pymysql.cursors.DictCursor, autocommit=True)
        self.cur = self.con.cursor()


def other_op(q, *params):
    """(str, *int/str) -> NoneType
    Runs a DELETE query
    """
    db = Database()
    db.cur.execute(q, params)
    db.con.close()


def insert(q, *params):
    """(str, *int/str) -> int
    Inserts into the db with the given INSERT INTO query, then returns the inserted id
    """
    db = Database()
    db.cur.execute(q, *params)
    ret_id = db.cur.lastrowid
    db.con.close()
    return ret_id


def fetch_all(q, *params):
    """(str, *int/str) -> list of dict
    Performs a SELECT query and returns a list of all results
    """
    db = Database()
    db.cur.execute(q, params)
    ret = db.cur.fetchall()
    db.con.close()
    return ret


def fetch_one(q, *params):
    """(str, *int/str) -> dict
    Performs a SELECT query and returns one result or None if it doesn't exist
    """
    db = Database()
    db.cur.execute(q, params)
    ret = db.cur.fetchone()
    db.con.close()
    return ret
