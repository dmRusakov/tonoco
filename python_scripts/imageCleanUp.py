import os
from models.db import DB

forErrOnly = True

# file folder
original_file_folder = "./../image/"
new_file_folder = "./assets/images/original/"

# dconnect to DB
db = DB()

no_wm_file_ends = ["_nowm", "_noWm"]
wm_file_ends = ["_wm", "_Wm", "_wm_1"]
errorSqlWhere = "where error is not null and error != ''" if forErrOnly else ""

# get data from database
counter = 1
while True:
    sql = f"select id, origin_path, error from image {errorSqlWhere} offset {str(counter * 10)} limit 10"
    data = db.get(sql)
    if len(data) == 0:
        break

    for row in data:
        # check file exist in the original folder
        id = row[0]
        file_name = row[1]
        if not os.path.exists(original_file_folder + file_name):
            name, exst = os.path.splitext(file_name)
            possible_files = [
                name + exst,
                name + "_wm" + exst,
                file_name.replace("-1.jpg", "_nowm.jpg"),
                file_name.replace("-1.jpg", "_noWm.jpg"),
                file_name.replace("-2.jpg", "_nowm.jpg"),
                file_name.replace("-2.jpg", "_noWm.jpg"),
                file_name.replace("-3.jpg", "_nowm.jpg"),
                file_name.replace("-3.jpg", "_noWm.jpg"),
                file_name.replace("-4.jpg", "_nowm.jpg"),
                file_name.replace("-4.jpg", "_noWm.jpg"),
                file_name.replace("-5.jpg", "_nowm.jpg"),
                file_name.replace("-5.jpg", "_noWm.jpg"),
                file_name.replace("-6.jpg", "_nowm.jpg"),
                file_name.replace("-6.jpg", "_noWm.jpg"),
                file_name.replace("nown.jpg", "_nowm.jpg"),
                name.replace("-1", "") + exst
            ]
            for possible_file in possible_files:
                if os.path.exists(original_file_folder + possible_file):
                    file_name = possible_file
                    break
            else:
                print("File not found: " + file_name)
                db.update("update image set error = 'File not found' where id = '" + str(id) + "'")
                continue

        # try ti get image without watermark `_wm`
        name, exst = os.path.splitext(file_name)
        # remove '_wm' , '_nowm' from file name
        for end in wm_file_ends:
            name = name.replace(end, "")
        for end in no_wm_file_ends:
            name = name.replace(end, "")

        clear_file_name = name + exst

        no_wm_file = None
        for end in no_wm_file_ends:
            if os.path.exists(original_file_folder + name + end + exst):
                no_wm_file = name + end + exst
                break

        if no_wm_file:
            os.system("cp " + original_file_folder + no_wm_file + " " + new_file_folder + clear_file_name)
        else:
            os.system("cp " + original_file_folder + file_name + " " + new_file_folder + clear_file_name)

        # update file name in the database
        db.update("update image set origin_path = '" + clear_file_name + "' where id = '" + str(id) + "'")

    counter += 1

print("Done")