#
# This file and its contents are supplied under the terms of the
# Common Development and Distribution License ("CDDL"). You may
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

COMPONENT_NAME=	          elasticsearch
COMPONENT_MAJOR_VERSION=  5.6
COMPONENT_MINOR_VERSION=  13
COMPONENT_VERSION=        $(COMPONENT_MAJOR_VERSION).$(COMPONENT_MINOR_VERSION)
COMPONENT_SUMMARY=        Elasticsearch is a distributed, RESTful search and analytics engine
COMPONENT_PROJECT_URL=    https://www.elastic.co
COMPONENT_FMRI=	          database/elasticsearch-56
COMPONENT_CLASSIFICATION= System/Databases
COMPONENT_SRC=      $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE=	$(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_URL=    https://artifacts.elastic.co/downloads/$(COMPONENT_NAME)/$(COMPONENT_ARCHIVE)
COMPONENT_ARCHIVE_HASH=	  sha256:6800471e65cf18f3580a5d88f1f9dac79c220408aef4bf18cccf295a5211b6b3
COMPONENT_LICENSE=	      Apache 2.0

COMPONENT_ARCHIVE_1= jna-4.4.0.tar.gz
COMPONENT_ARCHIVE_HASH_1= sha256:6743c351df7114f5c3eb8294a5dd6e76aa29b49ce3eb1efdc4130276a39882f2
COMPONENT_ARCHIVE_URL_1= https://github.com/java-native-access/jna/archive/4.4.0.tar.gz

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/ant.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

PKG_MACROS += COMPONENT_MAJOR_VERSION=$(COMPONENT_MAJOR_VERSION)

COMPONENT_POST_UNPACK_ACTION = ( \
	${FIND} $(SOURCE_DIR)/ -type f -name \*.bat -delete; \
	${FIND} $(SOURCE_DIR)/ -type f -name \*.exe -delete; \
	${CP} files/Makefile $(SOURCE_DIR);  \
	$(GTAR) xf $(USERLAND_ARCHIVES)$(COMPONENT_ARCHIVE_1); \
	cd jna-4.4.0; \
	($(ENV) $(COMPONENT_BUILD_ENV) $(ANT) dist test test-platform)); \
	$(CP) $(COMPONENT_DIR)/jna-4.4.0/build/jna.jar $(SOURCE_DIR)/lib/jna-4.4.0-1.jar

COMPONENT_INSTALL_ARGS += COMPONENT_DIR=$(USRDIR)/$(COMPONENT_NAME)/$(COMPONENT_MAJOR_VERSION)

build:    $(BUILD_64)

install:  $(INSTALL_64)

test:     $(NO_TESTS)

# Auto-generated dependencies
REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += runtime/java/openjdk8
REQUIRED_PACKAGES += shell/bash
