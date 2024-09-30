README_PATH=README.md
$(README_PATH): update_readme

.PHONY: update_readme
update_readme:
	USER=takeshi text-template-cli README.md.tmpl > README.md

.PHONY: update-dir-example1
update-dir-example1:
	TITLE=Example1 text-template-cli dir-example1/source --output-directory dir-example1/output
