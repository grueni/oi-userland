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
# Copyright 2018 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved. 
#

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME=           filebeat
COMPONENT_VERSION=        6.4.3
COMPONENT_VERSION_MAJOR=  6.4
COMPONENT_FMRI=           web/$(COMPONENT_NAME)-$(COMPONENT_VERSION_MAJOR)
COMPONENT_CLASSIFICATION= System/Administration and Configuration
COMPONENT_PROJECT_URL=    https://www.elastic.co/de/products/beats/filebeat 
COMPONENT_SRC=            beats
COMPONENT_ARCHIVE_HASH=   sha256:437dfffe2e33e63df3281444a0ed409673957f95e79c9bc834bc108e8b1534ca
COMPONENT_ARCHIVE=        beats-$(COMPONENT_VERSION).tar.gz
COMPONENT_ARCHIVE_URL=    https://github.com/elastic/beats/archive/v$(COMPONENT_VERSION).tar.gz
COMPONENT_SUMMARY=        Filebeat is an open source file harvester, mostly used to fetch logs files and feed them into logstash.  
COMPONENT_LICENSE=        Apache 2.0
COMPONENT_LICENSE_FILE=   LICENSE.txt 

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

COMPONENT_POST_UNPACK_ACTION = ( \
	$(MKDIR) $(SOURCE_DIR)/src/github.com/elastic/$(COMPONENT_SRC) && \
	$(MV)    $(COMPONENT_DIR)/$(COMPONENT_SRC)-$(COMPONENT_VERSION)/* $(SOURCE_DIR)/src/github.com/elastic/$(COMPONENT_SRC)  && \
	$(CP)    $(SOURCE_DIR)/src/github.com/elastic/$(COMPONENT_SRC)/LICENSE.txt $(SOURCE_DIR) && \
	$(RM) -r $(COMPONENT_DIR)/$(COMPONENT_SRC)-$(COMPONENT_VERSION) )

COMPONENT_PREP_ACTION = ( \
  $(CP) $(COMPONENT_DIR)/files/Makefile.openindiana $(@D)/src/github.com/elastic/$(COMPONENT_SRC))

GOPATH = $(@D)

COMPONENT_BUILD_ARGS =-C filebeat

COMPONENT_BUILD_ACTION = \
  cd $(@D)/src/github.com/elastic/$(COMPONENT_SRC); $(ENV) $(COMPONENT_BUILD_ENV) \
  $(GMAKE) $(COMPONENT_BUILD_GMAKE_ARGS) $(COMPONENT_BUILD_ARGS) $(COMPONENT_BUILD_TARGETS)

COMPONENT_INSTALL_ACTION = \
  cd $(@D)/src/github.com/elastic/$(COMPONENT_SRC); $(ENV) $(COMPONENT_INSTALL_ENV) \
	$(GMAKE) -f Makefile.openindiana \
		DESTDIR=$(PROTO_DIR) \
		BINDIR=$(USRBINDIR) \
		ETCDIR=$(ETCDIR)/$(COMPONENT_NAME) \
		USRSHAREDIR=$(USRSHAREDIR)/$(COMPONENT_NAME) 

COMPONENT_PRE_TEST_ACTION = echo "Hallo Welt"

COMPONENT_TEST_TARGETS =-C filebeat testsuite

COMPONENT_TEST_DIR =$(@D)/src/github.com/elastic/$(COMPONENT_SRC)

build:          $(BUILD_64)

install:        $(INSTALL_64)

test:           $(TESTS_64)

# Auto-generated dependencies
REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += system/library
