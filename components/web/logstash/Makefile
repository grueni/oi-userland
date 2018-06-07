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

COMPONENT_NAME=           logstash
COMPONENT_VERSION=        6.4.3
COMPONENT_VERSION_MAJOR=  6.4
COMPONENT_FMRI=           web/$(COMPONENT_NAME)-$(COMPONENT_VERSION_MAJOR)
COMPONENT_CLASSIFICATION= System/Administration and Configuration
COMPONENT_PROJECT_URL=    https://www.elastic.co/products/logstash
COMPONENT_SRC=            $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE_HASH=   sha256:72dcd1cf21e55cd44503ca7c452aaaa88286564ba1450d5d05941f88e2699cc3
COMPONENT_ARCHIVE=        $(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_URL=    https://artifacts.elastic.co/downloads/$(COMPONENT_NAME)/$(COMPONENT_ARCHIVE)
COMPONENT_SUMMARY=        Logstash is an open source data collection engine with real-time pipelining capabilities. 
COMPONENT_LICENSE=        Apache 2.0
COMPONENT_LICENSE_FILE=   LICENSE.txt 

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

PKG_MACROS += COMPONENT_VERSION_MAJOR=$(COMPONENT_VERSION_MAJOR)

COMPONENT_POST_UNPACK_ACTION = ( \
   find $(SOURCE_DIR)/  -type f -name \*.bat  -delete;   \
   $(MKDIR) $(SOURCE_DIR)/config/conf.d && \
   $(CP) files/Makefile $(SOURCE_DIR)  && \
   $(CP) files/default  $(SOURCE_DIR)/bin/env && \
   $(CP) files/logstash-messages.conf $(SOURCE_DIR)/config/conf.d )

COMPONENT_INSTALL_ARGS += COMPONENT_DIR=$(USRSHAREDIR)/$(COMPONENT_NAME)/$(COMPONENT_VERSION_MAJOR)

build:          $(BUILD_64)

install:        $(INSTALL_64)

test:           $(TESTS_64)

REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += system/library
REQUIRED_PACKAGES += developer/java/openjdk8
