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
# Copyright 2019 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved.
#

include ../../../make-rules/shared-macros.mk

PATH= $(PATH.gnu)

COMPONENT_NAME= argp-standalone
COMPONENT_VERSION= 1.3
COMPONENT_FMRI= application/cluster/$(COMPONENT_NAME)
COMPONENT_CLASSIFICATION= System/Services
COMPONENT_SRC= $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE= $(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_HASH= sha256:dec79694da1319acd2238ce95df57f3680fea2482096e483323fddf3d818d8be
COMPONENT_PROJECT_URL= https://www.lysator.liu.se/~nisse/
COMPONENT_ARCHIVE_URL= https://www.lysator.liu.se/~nisse/misc/$(COMPONENT_ARCHIVE)
COMPONENT_LICENSE= BSD
COMPONENT_SUMMARY= libargp is a portable library providing the command-line parsing functionality of the argp module of the GNU Portability Library (Gnulib).

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/configure.mk
include $(WS_MAKE_RULES)/ips.mk

#COMPONENT_PREP_ACTION=(cd $(SOURCE_DIR); $(CONFIG_SHELL) build )

build:          $(BUILD_64)

install:        $(INSTALL_64)

test:           $(TESTS_64)
