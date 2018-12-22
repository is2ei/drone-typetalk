# COLORS
RED    = $(shell printf "\33[31m")
GREEN  = $(shell printf "\33[32m")
WHITE  = $(shell printf "\33[37m")
YELLOW = $(shell printf "\33[33m")
RESET  = $(shell printf "\33[0m")

build: 
	@echo "${YELLOW}Building...${RESET}"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-typetalk
	@docker build -t is2ei/drone-typetalk .
	@echo "${GREEN}✔ successfully built.${RESET}\n"

push:
	@echo "${YELLOW}Pushing...${RESET}"
	@docker push is2ei/drone-typetalk
	@echo "${GREEN}✔ successfully pushed.${RESET}\n"