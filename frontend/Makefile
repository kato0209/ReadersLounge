lint:
	npm run lint

deps:
	npm install

prettier:
	npx prettier --check src

prettier/write:
	npx prettier --write src

codegen: deps
	npx openapi-generator-cli generate -i ../openapi.yml -o ./src/openapi -g typescript-axios -c ./docs/config.yml

typecheck:
	npx tsc --noEmit