all-recursive:
    
install:
	mkdir -p $(DESTDIR)$(ETCDIR)
	mkdir -p $(DESTDIR)$(USRBINDIR)
	mkdir -p $(DESTDIR)$(USRSHAREDIR)
	cp -r ./etc/*    $(DESTDIR)$(ETCDIR)
	cp -r ./bin/*    $(DESTDIR)$(USRBINDIR)
	cp -r ./share/*  $(DESTDIR)$(USRSHAREDIR)
	cp -r ./src      $(DESTDIR)$(USRSHAREDIR)/confluent-common/
	find     $(DESTDIR)$(COMPONENT_DIR) -maxdepth 1 -name .\* -type f -delete
	cp       LICENSE $(DESTDIR)
	cp       README $(DESTDIR)$(USRSHAREDIR)/doc/confluent-common

