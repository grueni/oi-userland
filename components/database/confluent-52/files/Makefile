all-recursive:
    
install:
	mkdir -p $(DESTDIR)$(USRSHAREDIR)
	mkdir -p $(DESTDIR)$(ETCDIR)
	mkdir -p $(DESTDIR)$(USRBINDIR)
	ln -s ../..$(USRSHAREDIR)/confluent-52/etc /etc/confluent-52
	cp -r .  $(DESTDIR)$(COMPONENT_DIR)
	rm       $(DESTDIR)$(COMPONENT_DIR)/Makefile
	find     $(DESTDIR)$(COMPONENT_DIR) -maxdepth 1 -name .\* -type f -delete
	mv       $(DESTDIR)$(COMPONENT_DIR)/LICENSE $(DESTDIR)
	rm -rf   $(DESTDIR)$(COMPONENT_DIR)/src
	mv       $(DESTDIR)$(COMPONENT_DIR)/README $(DESTDIR)$(COMPONENT_DIR)/share/doc/confluent-common
	mv       $(DESTDIR)$(COMPONENT_DIR)/etc $(DESTDIR)

