ENTRY := ./server.py
VENV_DIR := venv
PYTHON := python3
PY := $(VENV_DIR)/bin/python
PIP := $(PYTHON) -m pip

.PHONY: install run clean

all: run

$(VENV_DIR):
	$(PYTHON) -m venv $(VENV_DIR)
	$(PY) -m pip install -U pip

install: $(VENV_DIR)
	$(PY) -m pip install -r requirements.txt

run: install
	$(PY) $(ENTRY)

clean:
	rm -rf $(VENV_DIR) __pycache__
