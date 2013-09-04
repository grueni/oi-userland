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
# Copyright 2013 Adam Stevko. All rights reserved.
#

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME=		crmsh
COMPONENT_VERSION=	1.2.5
COMPONENT_HG_REV=	ef3f08547688
COMPONENT_SRC=		$(COMPONENT_NAME)-$(COMPONENT_HG_REV)
COMPONENT_ARCHIVE=	$(COMPONENT_HG_REV).tar.gz
COMPONENT_ARCHIVE_HASH= sha256:b0a2feddbec85bb90811bfdfd57ac7fe6731654f7dda36d0b18c5436dc6512f3
COMPONENT_PROJECT_URL=	https://savannah.nongnu.org/projects/crmsh
COMPONENT_ARCHIVE_URL=	http://hg.savannah.gnu.org/hgweb/crmsh/archive/$(COMPONENT_ARCHIVE)

include $(WS_TOP)/make-rules/prep.mk
include $(WS_TOP)/make-rules/configure.mk
include $(WS_TOP)/make-rules/ips.mk

CONFIGURE_OPTIONS += --localstatedir=/var

COMPONENT_PREP_ACTION=(cd $(SOURCE_DIR); ./autogen.sh)

build:          $(BUILD_32)

install:        $(INSTALL_32)

test:           $(TESTS_32)
