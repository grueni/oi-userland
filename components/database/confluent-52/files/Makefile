all-recursive:
    
install:
	mkdir -p $(DESTDIR)$(COMPONENT_DIR)
	cp -r .  $(DESTDIR)$(COMPONENT_DIR)
	rm       $(DESTDIR)$(COMPONENT_DIR)/Makefile
	find     $(DESTDIR)$(COMPONENT_DIR) -maxdepth 1 -name .\* -type f -delete
	mv       $(DESTDIR)$(COMPONENT_DIR)/LICENSE $(DESTDIR)
	mv       $(DESTDIR)$(COMPONENT_DIR)/README $(DESTDIR)$(COMPONENT_DIR)/share/doc/confluent-common

