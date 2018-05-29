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
COMPONENT_VERSION=        2.4.6
COMPONENT_SUMMARY=        Elasticsearch is a distributed, RESTful search and analytics engine
COMPONENT_PROJECT_URL=    https://www.elastic.co
COMPONENT_FMRI=	          database/elasticsearch-24
COMPONENT_CLASSIFICATION= System/Databases
COMPONENT_SRC=      $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE=	$(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_URL=    https://download.elastic.co/elasticsearch/release/org/elasticsearch/distribution/tar/$(COMPONENT_NAME)/$(COMPONENT_VERSION)/$(COMPONENT_ARCHIVE)
COMPONENT_ARCHIVE_HASH=	  sha256:5f7e4bb792917bb7ffc2a5f612dfec87416d54563f795d6a70637befef4cfc6f
COMPONENT_LICENSE=	      Apache 2.0

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

COMPONENT_POST_UNPACK_ACTION = ( \
	${FIND} $(SOURCE_DIR)/ -type f -name \*.bat -delete; \
	${FIND} $(SOURCE_DIR)/ -type f -name \*.exe -delete; \
	${CP} files/Makefile $(SOURCE_DIR) )

COMPONENT_INSTALL_ARGS += COMPONENT_DIR=$(USRDIR)/$(COMPONENT_NAME)/$(COMPONENT_VERSION)

build:    $(BUILD_64)

install:  $(INSTALL_64)

test:     $(TESTS_64)

REQUIRED_PACKAGES += developer/java/openjdk8
