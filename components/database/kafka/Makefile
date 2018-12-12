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

COMPONENT_NAME=	          kafka
COMPONENT_MAJOR_VERSION=  2.1
COMPONENT_MINOR_VERSION=  0
COMPONENT_VERSION=        $(COMPONENT_MAJOR_VERSION).$(COMPONENT_MINOR_VERSION)
COMPONENT_SUMMARY=        Apache Kafka is a distributed streaming platform
COMPONENT_PROJECT_URL=    https://kafka.apache.org
COMPONENT_FMRI=	          database/kafka-21
COMPONENT_CLASSIFICATION= System/Databases
COMPONENT_SRC=            $(COMPONENT_NAME)_2.12-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE=	      $(COMPONENT_SRC).tgz
COMPONENT_ARCHIVE_URL=    http://apache.mirror.iphh.net/kafka/$(COMPONENT_VERSION)/$(COMPONENT_ARCHIVE)
COMPONENT_ARCHIVE_HASH=	  sha256:a70b44a682f5082f5c69182d4596e7ded2df38629f09b50a629070ddb1ef4f86
COMPONENT_LICENSE=	      Apache-2.0
COMPONENT_LICENSE_FILE=	  LICENSE

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

PKG_MACROS += COMPONENT_MAJOR_VERSION=$(COMPONENT_MAJOR_VERSION)

COMPONENT_POST_UNPACK_ACTION = ( \
	$(RM) -r $(SOURCE_DIR)/site-docs; \
	$(RM) -r $(SOURCE_DIR)/bin/windows/; \
	$(CP)    files/Makefile $(SOURCE_DIR) )

COMPONENT_INSTALL_ARGS += COMPONENT_DIR=$(USRDIR)/$(COMPONENT_NAME)/$(COMPONENT_MAJOR_VERSION)

build:    $(BUILD_64)

install:  $(INSTALL_64)

test:     $(TESTS_64)

# Auto-generated dependencies
REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += runtime/java/openjdk8
REQUIRED_PACKAGES += shell/bash
