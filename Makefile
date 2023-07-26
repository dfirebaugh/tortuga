
.PHONY:docs
docs:
	bash ./scripts/build_docs.sh

serve-docs:
	bash ./scripts/serve_docs.sh

deploy-docs:
	bash ./scripts/deploy_docs.sh


.PHONY: run-example
run-example:
	go run ./examples/$(DIR)

clean:
	rm -rf bin .dist
