.PHONY: %

%:
	cd sherlock && $(MAKE) -f makefile $@
