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

COMPONENT_NAME= redis
COMPONENT_MAJOR_VERSION=  4.0
COMPONENT_MINOR_VERSION=  11
COMPONENT_VERSION=        $(COMPONENT_MAJOR_VERSION).$(COMPONENT_MINOR_VERSION)
COMPONENT_SUMMARY= Redis is an advanced key-value store.
COMPONENT_PROJECT_URL= https://redis.io/
COMPONENT_FMRI= database/redis-40
COMPONENT_CLASSIFICATION= System/Databases
COMPONENT_SRC= $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE= $(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_URL= http://download.redis.io/releases/$(COMPONENT_ARCHIVE)
COMPONENT_ARCHIVE_HASH= sha256:fc53e73ae7586bcdacb4b63875d1ff04f68c5474c1ddeda78f00e5ae2eed1bbb
COMPONENT_LICENSE= modified-bsd
COMPONENT_LICENSE_FILE= COPYING

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

PKG_MACROS += COMPONENT_MAJOR_VERSION=$(COMPONENT_MAJOR_VERSION)

COMPONENT_BUILD_ARGS += V=1
COMPONENT_BUILD_ARGS += CC=gcc
COMPONENT_BUILD_ARGS += PREFIX=/usr

COMPONENT_INSTALL_ARGS += $(COMPONENT_BUILD_ARGS)

# when you use make test:
# you need tcl and ruby 
# gem install redis 
# 

COMPONENT_TEST_ENV += http_proxy= https_proxy=
COMPONENT_TEST_CMD = ./runtest-cluster && ./runtest-sentinel && ./runtest
COMPONENT_TEST_TRANSFORMER = $(NAWK)
COMPONENT_TEST_TRANSFORMS = /All tests/

build:    $(BUILD_64)

install:  $(INSTALL_64)

test:     $(TEST_64)

# Auto-generated dependencies
REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += system/library
REQUIRED_PACKAGES += system/library/math
