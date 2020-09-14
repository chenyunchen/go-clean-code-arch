install-mock:
	brew install vektra/tap/mockery
	brew upgrade mockery

gen-mock:
	cd domain; mockery --all