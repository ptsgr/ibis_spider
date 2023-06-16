
.PHONY: migrate
migrate:
	@[ "${STORAGE_DSN}" ] || ( echo ">> STORAGE_DSN is not set"; exit 1 )
	goose -dir "migrations" postgres "$(STORAGE_DSN)" up

.PHONY: remigrate
remigrate:
	@[ "${STORAGE_DSN}" ] || ( echo ">> STORAGE_DSN is not set"; exit 1 )
	goose -dir "migrations" postgres "$(STORAGE_DSN)" redo

.PHONY: reset
reset:
	@[ "${STORAGE_DSN}" ] || ( echo ">> STORAGE_DSN is not set"; exit 1 )
	goose -dir "migrations" postgres "$(STORAGE_DSN)" reset