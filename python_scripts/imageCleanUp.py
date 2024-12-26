import os
from models.db import get_conn, get_data

# file folder
original_file_folder = "./../image/"
new_file_folder = "./assets/images/original/"

# dconnect to DB
conn = get_conn()

# get data from database
counter = 1
while True:
    data = get_data("select id, origin_path, error from image offset " + str(counter * 10) + " limit 10", conn)
    if len(data) == 0:
        break

    for row in data:
        # check file exist in the original folder
        file_name = row[1]
        if not os.path.exists(original_file_folder + file_name):
            continue

        # copy image to new folder
        os.system("cp " + original_file_folder + file_name + " " + new_file_folder + file_name)

    counter += 1

print("Done")