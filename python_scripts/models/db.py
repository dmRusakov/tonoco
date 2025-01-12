import os
import psycopg2

class DB:
    def __init__(self):
        db_params = {
            'dbname': os.environ['DATA_STORAGE_DB'],
            'user': os.environ['DATA_STORAGE_USER'],
            'host': os.environ['DATA_STORAGE_HOST'],
            'password': os.environ['DATA_STORAGE_PASSWORD']
        }
        self.conn = psycopg2.connect(**db_params)


    def get(self, query):
        data = []
        cur = self.conn.cursor()
        cur.execute(query)
        rows = cur.fetchall()
        for row in rows:
            data.append(row)
        cur.close()
        return data

    def close(self):
        self.conn.close()
        print("Connection closed")

    def update(self, query):
        cur = self.conn.cursor()
        cur.execute(query)
        self.conn.commit()
        cur.close()
