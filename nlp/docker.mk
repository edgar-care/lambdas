pip_config_file ?= requirements-local.txt
python_bin ?= python3

install: install-upgrade-pip install-deps
install-deps:
	@$(python_bin) -m pip --disable-pip-version-check install -r $(pip_config_file)
install-upgrade-pip:
	@$(python_bin) -m pip --disable-pip-version-check install --upgrade pip

start:
	@TF_CPP_MIN_LOG_LEVEL="2" uvicorn --host 0.0.0.0 --port 5999 app:app --reload

.DEFAULT_GOAL := install

.PHONY: install install-deps install-upgrade-pip