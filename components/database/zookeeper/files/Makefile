all-recursive:
    
install:
	mkdir -p $(DESTDIR)${PREFIX}/bin
	mkdir -p $(DESTDIR)${PREFIX}/lib/amd64
	mkdir -p $(DESTDIR)${PREFIX}/sbin
	mkdir -p $(DESTDIR)/etc/zookeeper
	mkdir -p $(DESTDIR)${PREFIX}/share/zookeeper
	cp $(BUILDDIR)/bin/* $(DESTDIR)${PREFIX}/bin/
	cp $(BUILDDIR)/libexec/* $(DESTDIR)${PREFIX}/bin/
	cp $(BUILDDIR)/sbin/* $(DESTDIR)${PREFIX}/bin/
	cp $(BUILDDIR)/conf/* $(DESTDIR)/etc/zookeeper/
	cp $(COMPONENT_DIR)/files/jmx.* $(DESTDIR)/etc/zookeeper/
	cp $(BUILDDIR)/share/zookeeper/templates/conf/zoo.cfg $(DESTDIR)/etc/zookeeper/zoo.cfg
	cp $(BUILDDIR)/share/zookeeper/templates/conf/zookeeper-env.sh $(DESTDIR)/etc/zookeeper/zookeeper-env.sh
	cp -r $(BUILDDIR)/share/zookeeper $(DESTDIR)${PREFIX}/share
	cp -r $(BUILDDIRC)/${PREFIX} $(DESTDIR)
	rm $(DESTDIR)${PREFIX}/bin/*.cmd
	rm $(DESTDIR)${PREFIX}/lib/*.a
	rm $(DESTDIR)${PREFIX}/lib/*.la
	mv $(DESTDIR)${PREFIX}/lib/lib* $(DESTDIR)${PREFIX}/lib/amd64/


