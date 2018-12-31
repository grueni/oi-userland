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

COMPONENT_NAME=	          graylog
COMPONENT_MAJOR_VERSION=  2.5
COMPONENT_MINOR_VERSION=  1
COMPONENT_VERSION=        $(COMPONENT_MAJOR_VERSION).$(COMPONENT_MINOR_VERSION)
COMPONENT_SUMMARY=        central management of logs with analytics engine
COMPONENT_PROJECT_URL=    https://www.graylog.org/
COMPONENT_FMRI=	          web/graylog-25
COMPONENT_CLASSIFICATION= Web Services/Application and Web Servers
COMPONENT_SRC=      $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE=	$(COMPONENT_SRC).tgz
COMPONENT_ARCHIVE_URL=    https://packages.graylog2.org/releases/$(COMPONENT_NAME)/$(COMPONENT_ARCHIVE)
COMPONENT_ARCHIVE_HASH=	  sha256:6bda2ceb39ff06c27c865dfc102125dcf4a7fa7ec796ef8e908de86c8634fad8
COMPONENT_LICENSE=	      GPL 3.0

# get SSO plugin
COMPONENT_NAME_1=       graylog-plugin-auth-sso
COMPONENT_VERSION_1=    2.4.2
COMPONENT_SRC_1=        $(COMPONENT_NAME_1)-$(COMPONENT_VERSION_1)
COMPONENT_ARCHIVE_1=    $(COMPONENT_SRC_1).jar
COMPONENT_ARCHIVE_HASH_1= sha256:e8b29f7588a70ab33e4729f8ae9af087020462ba5d7189e4b6fa0adcdae19f5b
COMPONENT_ARCHIVE_URL_1=  https://github.com/Graylog2/$(COMPONENT_NAME_1)/releases/download/$(COMPONENT_VERSION_1)/$(COMPONENT_ARCHIVE_1)

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

PKG_MACROS += COMPONENT_MAJOR_VERSION=$(COMPONENT_MAJOR_VERSION)

COMPONENT_POST_UNPACK_ACTION = ( \
	$(FIND) $(SOURCE_DIR)/lib/sigar -type f -not -name '*solaris*' -delete; \
  $(MKDIR) $(SOURCE_DIR)/config; \
  $(MV) $(SOURCE_DIR)/graylog.conf.example $(SOURCE_DIR)/config/; \
  $(CP) $(USERLAND_ARCHIVES)$(COMPONENT_ARCHIVE_1) $(SOURCE_DIR)/plugin/; \
	$(CP) files/Makefile $(SOURCE_DIR) )

# no need to unpack, it is a jar file
$(SOURCE_DIR_1)/.unpacked:        download Makefile $(PATCHES)
	$(RM) -r $(SOURCE_DIR_1)
	$(MKDIR) $(SOURCE_DIR_1)
	$(TOUCH) $@

COMPONENT_INSTALL_ARGS += COMPONENT_DIR=$(USRDIR)/$(COMPONENT_NAME)/$(COMPONENT_MAJOR_VERSION)

build:    $(BUILD_64)

install:  $(INSTALL_64)

test:     $(TESTS_64)

REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += runtime/java/openjdk8
REQUIRED_PACKAGES += shell/bash
REQUIRED_PACKAGES += system/library
