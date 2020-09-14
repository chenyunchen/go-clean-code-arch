install-mock:
	brew install vektra/tap/mockery
	brew upgrade mockery

gen-mock:
	cd internal/pkg/domain; mockery --all