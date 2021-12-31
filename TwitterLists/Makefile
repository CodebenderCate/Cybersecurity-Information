.PHONY: run
run:
	go run main.go
	git add README.md
	git commit --amend --no-edit
	git push -fu

.PHONY: watch
watch:
	find -iname '*.go' | entr -c make run
