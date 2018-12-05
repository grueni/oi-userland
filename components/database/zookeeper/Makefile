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

COMPONENT_NAME=          zookeeper
COMPONENT_MAJOR_VERSION= 3.4
COMPONENT_MINOR_VERSION= 13
COMPONENT_VERSION=	     $(COMPONENT_MAJOR_VERSION).$(COMPONENT_MINOR_VERSION)
COMPONENT_FMRI=          database/$(COMPONENT_NAME)
COMPONENT_SRC=		       $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE=	     $(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_HASH=	 sha256:7ced798e41d2027784b8fd55c908605ad5bd94a742d5dab2506be8f94770594d
COMPONENT_ARCHIVE_URL=	 http://www-eu.apache.org/dist/$(COMPONENT_NAME)/$(COMPONENT_SRC)/$(COMPONENT_ARCHIVE)
COMPONENT_PROJECT_URL=	 https://zookeeper.apache.org
COMPONENT_SUMMARY=       a high-performance coordination service for distributed applications.
COMPONENT_LICENSE=       Apache 2.0

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/ant.mk
include $(WS_MAKE_RULES)/ips.mk

COMPONENT_BUILD_TARGETS= compile compile-native
COMPONENT_INSTALL_TARGETS= bin-package
JAVA_HOME= $(JAVA8_HOME)

ENV += CFLAGS=-m64
ENV += LDFLAGS=-lm

COMPONENT_POST_UNPACK_ACTION = ( \
	$(CP) files/Makefile $(SOURCE_DIR); \
	$(CP) files/zoo.cfg $(SOURCE_DIR)/conf/zoo_sample.cfg; \
	$(CP) files/zookeeper-env.sh $(SOURCE_DIR)/src/packages/templates/conf/zookeeper-env.sh ) 

COMPONENT_PRE_INSTALL_ACTION += ( \
  $(RM) -r $(PROTO_DIR); \
  $(MKDIR) $(PROTO_DIR) ) 

COMPONENT_POST_INSTALL_ACTION= ( \
	cd $(@D); \
	gmake install \
				COMPONENT_DIR=$(COMPONENT_DIR) \
				BUILDDIRC=$(@D)/build/c/build \
				BUILDDIR=$(@D)/build/$(COMPONENT_SRC) \
				DESTDIR=$(PROTO_DIR) \
				PREFIX=$(USRDIR); )

build:    $(BUILD_64)

install:  $(INSTALL_64)

test:     $(TESTS_64)

REQUIRED_PACKAGES += developer/java/openjdk8
REQUIRED_PACKAGES += developer/cppunit
REQUIRED_PACKAGES += library/python-2/setuptools
REQUIRED_PACKAGES += shell/gnu-getopt
REQUIRED_PACKAGES += system/library/math
