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
# Copyright 2016 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved. 
#

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME=           cassandra
COMPONENT_VERSION=        3.11.1
COMPONENT_FMRI=           database/$(COMPONENT_NAME)-311
COMPONENT_CLASSIFICATION= System/Databases
COMPONENT_PROJECT_URL=    http://cassandra.apache.org/
COMPONENT_SRC=            apache-$(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE_HASH=   sha256:6feed696759e2f3219e0ebde13f4c9080cac0744a77eb3cb43af136ab527f0ed
COMPONENT_ARCHIVE=        $(COMPONENT_SRC)-bin.tar.gz
COMPONENT_ARCHIVE_URL=    http://apache.mirror.iphh.net/$(COMPONENT_NAME)/$(COMPONENT_VERSION)/$(COMPONENT_ARCHIVE)
COMPONENT_LICENSE=        Apache 2.0
COMPONENT_SUMMARY=        Apache Cassandra is a highly-scalable partitioned row store. 

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

# rm files for windows, lib/sigar-bin contains binaries for all OS supported by cassandra, 
# ise custom Makefile
COMPONENT_POST_UNPACK_ACTION = ( \
   find $(SOURCE_DIR)/              -type f      -name \*.bat      -delete;   \
   find $(SOURCE_DIR)/              -type f      -name \*.ps1      -delete;   \
   find $(SOURCE_DIR)/lib/sigar-bin -type f -not -name '*solaris*' -delete;   \
   cp files/Makefile $(SOURCE_DIR) )

COMPONENT_INSTALL_ARGS += COMPONENT_DIR=$(USRDIR)/$(COMPONENT_NAME)/$(COMPONENT_VERSION)

build:          $(BUILD_64)

install:        $(INSTALL_64)

test:           $(TESTS_64)

REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += system/library
REQUIRED_PACKAGES += developer/java/openjdk8

