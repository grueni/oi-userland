#
# This file and its contents are supplied under the terms of the
# Common Development and Distribution License ("CDDL)". You may
# only use this file in accordance with the terms of the CDDL.
#
# A full copy of the text of the CDDL should have accompanied this
# source. A copy of the CDDL is also available via the Internet at
# http://www.illumos.org/license/CDDL.
#

#
# Copyright 2017 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved. 
#

include ../../../make-rules/shared-macros.mk

PATH= $(PATH.gnu)

COMPONENT_NAME=           couchdb
COMPONENT_VERSION=        1.6.1
COMPONENT_FMRI=           database/$(COMPONENT_NAME)-16
COMPONENT_CLASSIFICATION= System/Databases
COMPONENT_PROJECT_URL=    http://couchdb.apache.org/
COMPONENT_SRC=            apache-$(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE_HASH=   sha256:5a601b173733ce3ed31b654805c793aa907131cd70b06d03825f169aa48c8627
COMPONENT_ARCHIVE=        $(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_URL=    http://apache.mirror.iphh.net/$(COMPONENT_NAME)/source/$(COMPONENT_VERSION)/$(COMPONENT_ARCHIVE)
COMPONENT_LICENSE=        Apache 2.0
COMPONENT_SUMMARY=        Apache CouchDB is a document-oriented database.

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/configure.mk
include $(WS_MAKE_RULES)/ips.mk

CC= gcc
CXX= g++

CFLAGS += -D__EXTENSIONS__ -D_POSIX_PTHREAD_SEMANTICS
CXXFLAGE += -D__EXTENSIONS__ -D_POSIX_PTHREAD_SEMANTICS

COMPONENT_PRE_CONFIGURE_ACTION=($(CLONEY) $(SOURCE_DIR) $(@D))

#COMPONENT_PREP_ACTION = ( \
#	echo $(PATH); \
#	$(CP) $(COMPONENT_DIR)/files/Makefile.openindiana $(@D); \
#	$(CP) $(COMPONENT_DIR)/files/default.ini $(@D); )

CONFIGURE_SCRIPT= $(@D)/configure

#CONFIGURE_OPTIONS += -u $(COMPONENT_NAME)
CONFIGURE_OPTIONS += CC="$(CC)" CXX="$(CXX)" CFLAGS="$(CFLAGS)" LDFLAGS="$(LDFLAGS)"

#COMPONENT_BUILD_ENV += CC="$(CC)" CXX="$(CXX)" CFLAGS="$(CFLAGS)" LDFLAGS="$(LDFLAGS)" ERLC="$(ERLC)" ERL="$(ERL)" 
COMPONENT_BUILD_ENV += --docdir="$(USRSHAREDOCDIR)"

COMPONENT_INSTALL_ARGS += COMPONENT_DIR=$(USRDIR)/$(COMPONENT_NAME)/$(COMPONENT_VERSION)

#COMPONENT_INSTALL_TARGETS = release

$(xINSTALL_32): COMPONENT_POST_INSTALL_ACTION= \
  ($(MKDIR) $(PROTO_DIR)$(USRLIBDIR32); \
   $(CP) $(BUILD_DIR_32)/LICENSE $(COMPONENT_NAME).license; \
   $(CP) -r $(BUILD_DIR_32)/rel/couchdb $(PROTO_DIR)$(USRLIBDIR32) ); \
	 cd $(BUILD_DIR_32); \
	 $(GMAKE) -f Makefile.openindiana DESTDIR=$(PROTO_DIR) BINDIR=$(USRBINDIR) MANDIR=$(USRSHAREMANDIR) \
			DOCDIR=$(USRSHAREDOCDIR)/couchdb ETCDIR=$(ETCDIR)/couchdb RELDIR=$(USRLIBDIR32)/couchdb

build:          $(BUILD_32)

install:        $(INSTALL_32)

test:           $(TEST_32)

REQUIRED_PACKAGES += runtime/erlang
REQUIRED_PACKAGES += developer/icu
REQUIRED_PACKAGES += system/library/g++-4-runtime
REQUIRED_PACKAGES += runtime/spidermonkey

