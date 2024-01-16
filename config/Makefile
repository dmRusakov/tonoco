CONFIG_FOLDER=`basename $(CURDIR)`

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
	sudo ./deleteDbData.sh
	sleep 1 # wait 1 sec for db to be deleted
	@echo "Deleted"

webAssembly: makeWebAssembly minJSFiles
	@echo "WebAssembly builded"

makeWebAssembly:
	# get string with YY.MM.DD:HH:MM:SS for versioning
	$(eval VERSION := $(shell date +"%y.%m.%d:%H:%M:%S"))
	#get hash on the string in the variable VERSION
	$(eval VERSION := $(shell echo -n "adminAppApi$(VERSION)pieceOfLove" | sha256sum | awk '{print $$1}'))
	# get first 8 symbols of hash
	$(eval VERSION := $(shell echo $(VERSION) | cut -c1-8))
	# build wasm
	GOOS=js GOARCH=wasm go build -o ./../bin/adminApp.wasm ./../cmd/adminAppWebAssembly/main.go
	# name file name with version
	$(eval FILENAME := $(VERSION).wasm)
	#clear folder assets/wasm
	rm -rf ./../assets/wasm/*
	# move file to assets/wasm and rename with version
	mv ./../bin/adminApp.wasm './../assets/wasm/./$(FILENAME)'
	# open assets/js/wasm.js, fine variable appFileName and change value to versioned file name
	sed -i "s/appFileName = .*/appFileName = \"$(FILENAME)\";/g" ./../assets/js/wasm.js

minJSFiles:
	# minify all js files in the assets/js folder
	$(foreach file, $(filter-out %.min.js, $(wildcard ./../assets/js/*.js)), \
		uglifyjs $(file) > $(subst .js,.tmp.min.js,$(file)); \
		if cmp -s $(subst .js,.tmp.min.js,$(file)) $(subst .js,.min.js,$(file)); then \
	  		rm $(subst .js,.tmp.min.js,$(file)); \
		else \
			mv $(subst .js,.tmp.min.js,$(file)) $(subst .js,.min.js,$(file)); \
			$(eval FILENAME := $(file)) \
			$(eval FILENAME := $(shell echo $(FILENAME) | sed 's/.*\///')) \
			$(eval FILENAME := $(shell echo $(FILENAME) | sed 's/\.js//')) \
			$(eval ID := "$(FILENAME)-js") \
			FILENAME=$(FILENAME)  make updateCJsVersion; \
		fi; \
	)

updateCJsVersion:
	$(eval ID := "$(FILENAME)-js")
	@echo $(ID)
	$(eval VERSION := $(shell grep '<script.*id=$(ID)' "./../assets/templates/head_file_imports.partial.gohtml" | sed -n 's/.*version="\([^"]*\)".*/\1/p'))
	$(if $(VERSION),,$(eval VERSION := 1.0))
	$(eval VERSION := $(shell echo "$(VERSION) + 0.001" | bc))
	@echo $(VERSION)
	sed -i 's/\(script id=$(ID).*version="\)[^"]*"/\1'$(VERSION)'"/' ./../assets/templates/head_file_imports.partial.gohtml
	sed -i 's/\(script id=$(ID).*version="\)[^"]*"/\1'$(VERSION)'"/' ./../assets/templates/prod_head_file_imports.partial.gohtml
	sed -i 's/\(script id=$(ID).*src="\)[^"]*"/\1'"\/assets\/js\/"$(FILENAME)".min.js?v="$(VERSION)'"/' ./../assets/templates/prod_head_file_imports.partial.gohtml

minCssFiles:
 	# minify all css files in the assets/css folder
	$(foreach file, $(filter-out %.min.css, $(wildcard ./../assets/css/*.css)), \
		uglifycss $(file) > $(subst .css,.tmp.min.css,$(file)); \
		if cmp -s $(subst .css,.tmp.min.css,$(file)) $(subst .css,.min.css,$(file)); then \
	  		rm $(subst .css,.tmp.min.css,$(file)); \
		else \
			mv $(subst .css,.tmp.min.css,$(file)) $(subst .css,.min.css,$(file)); \
			$(eval FILENAME := $(file)) \
			$(eval FILENAME := $(shell echo $(FILENAME) | sed 's/.*\///')) \
			$(eval FILENAME := $(shell echo $(FILENAME) | sed 's/\.css//')) \
			$(eval ID := "$(FILENAME)-css") \
			FILENAME=$(FILENAME)  make updateCssVersion; \
		fi; \
	)

updateCssVersion:
	$(eval ID := "$(FILENAME)-css")
	@echo $(ID)
	$(eval VERSION := $(shell grep '<link.*id=$(ID)' "./../assets/templates/head_file_imports.partial.gohtml" | sed -n 's/.*version="\([^"]*\)".*/\1/p'))
	$(if $(VERSION),,$(eval VERSION := 1.0))
	$(eval VERSION := $(shell echo "$(VERSION) + 0.001" | bc))
	@echo $(VERSION)
	sed -i 's/\(link id=$(ID).*version="\)[^"]*"/\1'$(VERSION)'"/' ./../assets/templates/head_file_imports.partial.gohtml
	sed -i 's/\(link id=$(ID).*version="\)[^"]*"/\1'$(VERSION)'"/' ./../assets/templates/prod_head_file_imports.partial.gohtml
	sed -i 's/\(link id=$(ID).*href="\)[^"]*"/\1'"\/assets\/css\/"$(FILENAME)".min.css?v="$(VERSION)'"/' ./../assets/templates/prod_head_file_imports.partial.gohtml
