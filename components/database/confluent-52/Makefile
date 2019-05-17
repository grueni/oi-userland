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
# Copyright 2019 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved. 
#

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME=	          confluent
COMPONENT_MAJOR_VERSION=  5.2
COMPONENT_MINOR_VERSION=  1
COMPONENT_VERSION=        $(COMPONENT_MAJOR_VERSION).$(COMPONENT_MINOR_VERSION)
COMPONENT_VERSION_KAFKA=  2.12
COMPONENT_SUMMARY=        Confluent community version of Apache Kafka, Zookeeper, KSQL, schema registry, Rest proxy
COMPONENT_PROJECT_URL=    https://www.confluent.io
COMPONENT_FMRI=	          database/confluent-52
COMPONENT_CLASSIFICATION= System/Databases
COMPONENT_SRC=            $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE=	      confluent-community-$(COMPONENT_MAJOR_VERSION).$(COMPONENT_MINOR_VERSION)-$(COMPONENT_VERSION_KAFKA).tar.gz 
COMPONENT_ARCHIVE_URL=    http://packages.confluent.io/archive/$(COMPONENT_MAJOR_VERSION)/$(COMPONENT_ARCHIVE)
COMPONENT_ARCHIVE_HASH=	  sha256:c5612a324025a84b77b063b5b5698abd36e444e87ca4e01fa4bfbe5aa6396f61
COMPONENT_LICENSE=	      Apache-2.0,Confluent Community License
COMPONENT_LICENSE_FILE=	  LICENSE

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

PKG_MACROS += COMPONENT_MAJOR_VERSION=$(COMPONENT_MAJOR_VERSION)

COMPONENT_POST_UNPACK_ACTION = ( \
	$(RM) -r $(SOURCE_DIR)/bin/windows/; \
	cat $(SOURCE_DIR)/share/doc/kafka/LICENSE > $(SOURCE_DIR)/LICENSE; \
	cat $(SOURCE_DIR)/share/doc/ksql/LICENSE >> $(SOURCE_DIR)/LICENSE; \
	$(CP)    files/Makefile $(SOURCE_DIR) )

COMPONENT_INSTALL_ARGS += COMPONENT_DIR=$(USRSHAREDIR)/confluent-52

build:    $(BUILD_64)

install:  $(INSTALL_64)

test:     $(TESTS_64)

# Auto-generated dependencies
REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += runtime/java/openjdk8
REQUIRED_PACKAGES += shell/bash
