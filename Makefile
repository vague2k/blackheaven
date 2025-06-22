templ:
	@templ generate --watch --proxy="localhost:3000" --open-browser=false

air:
	@air \
    --build.cmd "go build -o tmp/bin/main ./main.go" \
    --build.bin "tmp/bin/main" \
    --build.delay "100" \
    --build.exclude_dir "node_modules" \
    --build.include_ext "go" \
    --build.stop_on_error "false" \
    --misc.clean_on_exit true

tailwind:
	@tailwindcss -i ./views/assets/css/input.css -o ./views/assets/css/output.css --watch

dev:
	@make -j3 tailwind templ air

components:
	@templui add selectbox separator form input label textarea toast table dropdown drawer icon

vet: lint format test
