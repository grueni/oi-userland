
install:
	/bin/mkdir -p  $(DESTDIR)$(ETCDIR)/webmin
	/bin/mkdir -p  $(DESTDIR)$(VARDIR)/webmin
	/bin/mkdir -p  $(DESTDIR)$(VARDIR)/tmp/webmin
	/bin/mkdir -p  $(DESTDIR)$(WEBMINDIR)
	./setup.sh $(DESTDIR)$(WEBMINDIR)
	mv config $(DESTDIR)$(ETCDIR)/webmin/config
	cd $(DESTDIR)$(WEBMINDIR); tar xzf $(THEME)
	tar zxf acl/Authen-SolarisRBAC-0.1.tar.gz
	cp Makefile.Authen-SolarisRBAC Authen-SolarisRBAC-0.1/Makefile
	cd Authen-SolarisRBAC-0.1; gmake V=1; gmake install $(DESTDIR)
	find $(DESTDIR)$(WEBMINDIR) -type f -iname '*~' -delete 
	find $(DESTDIR)$(WEBMINDIR) -type f -iname '.*' -delete 
	rm $(DESTDIR)$(WEBMINDIR)/Makefile* 
	rm $(DESTDIR)$(ETCDIR)/webmin/webmincron/crons/*
	grep -irl "${DESTDIR}" ${DESTDIR} | xargs sed -i -- "s~${DESTDIR}~~g"
