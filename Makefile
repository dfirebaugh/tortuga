
.PHONY:docs
docs:
	bash ./scripts/build_docs.sh

serve-docs:
	bash ./scripts/serve_docs.sh

clean:
	rm -rf bin .dist
