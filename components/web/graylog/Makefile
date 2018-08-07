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
# Copyright 2016 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved. 
#

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME=	          graylog
COMPONENT_VERSION=        2.4.6
COMPONENT_SUMMARY=        central management of logs with analytics engine
COMPONENT_PROJECT_URL=    https://www.graylog.org/
COMPONENT_FMRI=	          web/analytics/graylog-24
COMPONENT_CLASSIFICATION= Web Services/Application and Web Servers
COMPONENT_SRC=      $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE=	$(COMPONENT_SRC).tgz
COMPONENT_ARCHIVE_URL=    https://packages.graylog2.org/releases/$(COMPONENT_NAME)/$(COMPONENT_ARCHIVE)
COMPONENT_ARCHIVE_HASH=	  sha256:fcfaf44c3faea8297f340ddc6ed19e5b1fe8f3de3c1b2a1078119565fe2f751d
COMPONENT_LICENSE=	      GPL 3.0

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

COMPONENT_POST_UNPACK_ACTION = ( \
	$(FIND) $(SOURCE_DIR)/lib/sigar -type f -not -name '*solaris*' -delete; \
  $(MKDIR) $(SOURCE_DIR)/config; \
  $(MV) $(SOURCE_DIR)/graylog.conf.example $(SOURCE_DIR)/config/; \
	$(CP) files/Makefile $(SOURCE_DIR) )

COMPONENT_INSTALL_ARGS += COMPONENT_DIR=$(USRDIR)/$(COMPONENT_NAME)/$(COMPONENT_VERSION)

build:    $(BUILD_64)

install:  $(INSTALL_64)

test:     $(TESTS_64)

REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += runtime/java/openjdk8
REQUIRED_PACKAGES += shell/bash
REQUIRED_PACKAGES += system/library
