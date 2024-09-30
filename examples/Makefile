README_PATH=README.md
$(README_PATH): update_readme

.PHONY: update_readme
update_readme:
	USER=takeshi text-template-cli README.md.tmpl > README.md

.PHONY: update-dir-example1
update-dir-example1: update-dir-example1-output update-dir-example1-output-files

.PHONY: update-dir-example1-output
update-dir-example1-output:
	TITLE=Example1 text-template-cli dir-example1/source --output-directory dir-example1/output

.PHONY: update-dir-example1-output-files
update-dir-example1-output-files: update-dir-example1-output-readme update-dir-example1-output-text

.PHONY: update-dir-example1-output-readme
update-dir-example1-output-readme:
	TITLE=Example2 text-template-cli dir-example1/source/dir1/README.md.tmpl --output-directory dir-example1/output

.PHONY: update-dir-example1-output-text
update-dir-example1-output-text:
	text-template-cli dir-example1/source/dir2/simple1.txt --output-directory dir-example1/output
