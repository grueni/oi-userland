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
# Copyright 2017 Alexander Pyhalov
#

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME= redis
COMPONENT_VERSION= 4.0.2
COMPONENT_SUMMARY= Redis is an advanced key-value store.
COMPONENT_PROJECT_URL= https://redis.io/
COMPONENT_FMRI= database/redis-40
COMPONENT_CLASSIFICATION= System/Databases
COMPONENT_SRC= $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE= $(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_URL= http://download.redis.io/releases/$(COMPONENT_ARCHIVE)
COMPONENT_ARCHIVE_HASH= sha256:b1a0915dbc91b979d06df1977fe594c3fa9b189f1f3d38743a2948c9f7634813
COMPONENT_LICENSE= modified-bsd
COMPONENT_LICENSE_FILE= COPYING

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

COMPONENT_BUILD_ARGS += V=1
COMPONENT_BUILD_ARGS += CC=gcc
COMPONENT_BUILD_ARGS += PREFIX=/usr

COMPONENT_INSTALL_ARGS += $(COMPONENT_BUILD_ARGS)

# when you use make test:
# you need tcl and ruby 
# gem install redis 
# 

COMPONENT_TEST_CMD= ./runtest-cluster && ./runtest-sentinel && ./runtest
COMPONENT_TEST_TARGETS= 

build:    $(BUILD_64)

install:  $(INSTALL_64)

test:     $(TEST_64)

REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += system/library
REQUIRED_PACKAGES += system/library/math
