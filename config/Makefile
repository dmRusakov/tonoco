APP_PATH = $(shell pwd)/..
CONFIG_FOLDER=`basename $(CURDIR)`
MIN_JS_FILE = uglifyjs
MIN_CSS_FILE = uglifycss
JS_FILE_LOCATION = ./../assets/js/
CSS_FILE_LOCATION = ./../assets/css/

admin:

up: down
	docker-compose --env-file ./dev.env up -d
	@echo "Started"

down:
	docker-compose --env-file ./dev.env down --volumes
	@echo "Stopped"

reset: down delete up down up
	@echo "Reseted"

restart:
	docker-compose --env-file ./dev.env restart
	@echo "Restarted"

delete:
	docker-compose --env-file ./dev.env down --volumes
	sudo chmod -R 777 ./dbData/DataStorage/data/
	sudo chmod -R 777 ./dbData/Cassandra/data/
	sudo chmod -R 777 ./dbData/MqBroker/data/
	rm -rf ./dbData/DataStorage/data/*
	rm -rf ./dbData/CacheStorage/data/*
	rm -rf ./dbData/MqBroker/data/*
	sleep 1 # wait 1 sec for db to be deleted
	@echo "Deleted"

webAssembly: makeTinyGoWebAssemblyCompressed compressWasm makeWebAssemblyJS
	@echo "WebAssembly builded"

makeWebAssembly:
	GOOS=js GOARCH=wasm go build -o $(APP_PATH)/bin/adminApp.wasm $(APP_PATH)/cmd/adminAppWebAssembly/main.go

makeTinyGoWebAssembly:
	tinygo build -target wasm -opt=s -o $(APP_PATH)/bin/adminApp.wasm $(APP_PATH)/cmd/adminAppWebAssembly/main.go

makeTinyGoWebAssemblyCompressed:
	tinygo build -target wasm -no-debug -panic=trap -scheduler=none -gc=leaking -opt=s -o $(APP_PATH)/bin/adminApp.wasm $(APP_PATH)/cmd/adminAppWebAssembly/main.go

compressWasm:
	gzip -9 -k -f $(APP_PATH)/bin/adminApp.wasm

makeWebAssemblyJS:
	if cmp -s $(APP_PATH)/bin/adminApp.wasm $(APP_PATH)/assets/wasm/adminApp.wasm; then \
		echo "same size"; \
		rm ./../bin/adminApp.wasm; \
	else \
		mv ./../bin/adminApp.wasm ./../assets/wasm/adminApp.wasm; \
		$(eval VERSION := $(shell grep 'appFileVersion' ./../assets/js/wasm.js | sed -n 's/.*appFileVersion = "\(.*\)";/\1/p'))\
		$(if $(VERSION),,$(eval VERSION := 1.0))\
		$(eval VERSION := $(shell echo "$(VERSION) + 0.001" | bc))\
		sed -i "s/appFileVersion = .*/appFileVersion = \"$(VERSION)\";/g" ./../assets/js/wasm.js ; \
		make minJSFile file=wasm.js; \
	fi

minJSFile: # make minJSFile file=<filename.js>
	# check if file provided
	$(if $(file),,$(error file is not set))
	$(MIN_JS_FILE) $(JS_FILE_LOCATION)$(file) -o $(JS_FILE_LOCATION)$(subst .js,.min.js,$(file)) || { echo "uglifyjs failed for $(file)"; exit 1; }

 # make compressCssFile file=<filename.css>. For Jetbrains IDEs, need to add to Tools -> File Watchers for automatic minification
compressCssFile:
	# check if file provided
	$(if $(file),,$(error file is not set))
	$(MIN_CSS_FILE) $(CSS_FILE_LOCATION)$(file) > $(CSS_FILE_LOCATION)$(subst .css,.min.css,$(file)) || { echo "uglifycss failed for $(file)"; exit 1; }

## minify js files. For Jetbrains IDEs, need to add to Tools -> File Watchers for automatic minification
#updateJsVersion:
#	# delete all min.js files
#	$(eval Id := "$(FILENAME)-js")
#	@echo $(Id)
#	$(eval VERSION := $(shell grep '<script.*id=$(Id)' "./../assets/templates/element/head_file_imports.partial.gohtml" | sed -n 's/.*version="\([^"]*\)".*/\1/p'))
#	$(if $(VERSION),,$(eval VERSION := 1.0))
#	$(eval VERSION := $(shell echo "$(VERSION) + 0.001" | bc))
#	@echo $(VERSION)
#	sed -i 's/\(script id=$(Id).*version="\)[^"]*"/\1'$(VERSION)'"/' ./../assets/templates/element/head_file_imports.partial.gohtml
#	sed -i 's/\(script id=$(Id).*version="\)[^"]*"/\1'$(VERSION)'"/' ./../assets/templates/element/prod_head_file_imports.partial.gohtml
#	sed -i 's/\(script id=$(Id).*src="\)[^"]*"/\1'"\/assets\/js\/"$(FILENAME)".min.js?v="$(VERSION)'"/' ./../assets/templates/element/prod_head_file_imports.partial.gohtml
#
#minCssFiles:
# 	# minify all css files in the assets/css folder
#	$(foreach file, $(filter-out %.min.css, $(wildcard ./../assets/css/*.css)), \
#		if [ $(file) != $(subst .css,.min.css,$(file)) ]; then \
#			$(MIN_CSS_FILE) $(file) > $(subst .css,.tmp.min.css,$(file)); \
#			if cmp -s $(subst .css,.tmp.min.css,$(file)) $(subst .css,.min.css,$(file)); then \
#		  		rm $(subst .css,.tmp.min.css,$(file)); \
#			else \
#				mv $(subst .css,.tmp.min.css,$(file)) $(subst .css,.min.css,$(file)); \
#				$(eval FILENAME := $(file)) \
#				$(eval FILENAME := $(shell echo $(FILENAME) | sed 's/.*\///')) \
#				$(eval FILENAME := $(shell echo $(FILENAME) | sed 's/\.css//')) \
#				$(eval Id := "$(FILENAME)-css") \
#				FILENAME=$(FILENAME)  make updateCssVersion; \
#			fi; \
#		fi; \
#	)
#	# delete all min.min.css files
#	rm -f ./../assets/css/*.min.min.css

updateCssVersion:
	$(eval Id := "$(FILENAME)-css")
	@echo $(Id)
	$(eval VERSION := $(shell grep '<link.*id=$(Id)' "./../assets/templates/element/head_file_imports.partial.gohtml" | sed -n 's/.*version="\([^"]*\)".*/\1/p'))
	$(if $(VERSION),,$(eval VERSION := 1.0))
	$(eval VERSION := $(shell echo "$(VERSION) + 0.001" | bc))
	@echo $(VERSION)
	sed -i 's/\(link id=$(Id).*version="\)[^"]*"/\1'$(VERSION)'"/' ./../assets/templates/element/head_file_imports.partial.gohtml
	sed -i 's/\(link id=$(Id).*version="\)[^"]*"/\1'$(VERSION)'"/' ./../assets/templates/element/prod_head_file_imports.partial.gohtml
	sed -i 's/\(link id=$(Id).*href="\)[^"]*"/\1'"\/assets\/css\/"$(FILENAME)".min.css?v="$(VERSION)'"/' ./../assets/templates/element/prod_head_file_imports.partial.gohtml
	$(eval VERSION := $(shell grep '<link.*id=$(Id)' "./../assets/templates/element/footer_file_imports.partial.gohtml" | sed -n 's/.*version="\([^"]*\)".*/\1/p'))
	$(if $(VERSION),,$(eval VERSION := 1.0))
	$(eval VERSION := $(shell echo "$(VERSION) + 0.001" | bc))
	@echo $(VERSION)
	sed -i 's/\(link id=$(Id).*version="\)[^"]*"/\1'$(VERSION)'"/' ./../assets/templates/element/footer_file_imports.partial.gohtml
	sed -i 's/\(link id=$(Id).*version="\)[^"]*"/\1'$(VERSION)'"/' ./../assets/templates/element/prod_footer_file_imports.partial.gohtml
	sed -i 's/\(link id=$(Id).*href="\)[^"]*"/\1'"\/assets\/css\/"$(FILENAME)".min.css?v="$(VERSION)'"/' ./../assets/templates/element/prod_footer_file_imports.partial.gohtml

installNode:
	sudo dnf install nodejs
	sudo dnf install npm
	sudo dnf install tinygo
	npm install -g uglifycss
	npm install -g uglify-js
	npm install -g uglifyjs