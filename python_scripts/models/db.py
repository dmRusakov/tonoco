import os
import psycopg2

def get_conn():
    return psycopg2.connect(f"dbname='{os.environ['DATA_STORAGE_DB']}' "
                        f"user='{os.environ['DATA_STORAGE_USER']}' "
                        f"host='{os.environ['DATA_STORAGE_HOST']}' "
                        f"password='{os.environ['DATA_STORAGE_PASSWORD']}'")

def get_data(query, conn):
    data = []
    cur = conn.cursor()
    cur.execute(query)
    rows = cur.fetchall()
    for row in rows:
        data.append(row)
    cur.close()
    return data