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

COMPONENT_NAME= libjemalloc
COMPONENT_VERSION= 4.4.0
COMPONENT_PROJECT_URL= http://jemalloc.net/
COMPONENT_SUMMARY= http://jemalloc.net/
COMPONENT_FMRI= library/libjemalloc
COMPONENT_CLASSIFICATION=	System/Libraries
COMPONENT_SRC= jemalloc-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE= jemalloc-$(COMPONENT_VERSION).tar.bz2
COMPONENT_ARCHIVE_HASH=	sha256:a7aea63e9718d2f1adf81d87e3df3cb1b58deb86fc77bad5d702c4c59687b033
COMPONENT_ARCHIVE_URL= https://github.com/jemalloc/jemalloc/releases/download/$(COMPONENT_VERSION)/$(COMPONENT_ARCHIVE)
COMPONENT_LICENSE= BSD-2

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/configure.mk
include $(WS_MAKE_RULES)/ips.mk

CFLAGS += -std=gnu11 -Wall -Werror=declaration-after-statement -Wsign-compare -pipe -g3 -funroll-loops

COMPONENT_PRE_BUILD_ACTION=(cd $(@D); $(GMAKE) dist )

# common targets
build:    $(BUILD_32_and_64)

install:  $(INSTALL_32_and_64)

test:     $(TEST_32_and_64)
