README_PATH=README.md
$(README_PATH): update_readme

.PHONY: update_readme
update_readme:
	text-template-cli README.md.tmpl > README.md
