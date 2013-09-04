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
# Copyright 2016 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved.
#

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME= resource-agents
COMPONENT_VERSION= 3.9.7
COMPONENT_FMRI= application/cluster/$(COMPONENT_NAME)
COMPONENT_CLASSIFICATION= System/Libraries
COMPONENT_SRC= $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE= v$(COMPONENT_VERSION).tar.gz
COMPONENT_ARCHIVE_HASH= sha256:e5bd62658fbc236acb83b709f64b2cd9fae52aa4a420a44fed5eb667e928b152
COMPONENT_PROJECT_URL= https://github.com/ClusterLabs/resource-agents
COMPONENT_ARCHIVE_URL= https://github.com/ClusterLabs/$(COMPONENT_NAME)/archive/$(COMPONENT_ARCHIVE)
COMPONENT_SUMMARY= Combined repository of OCF agents from the RHCS and Linux-HA projects

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/configure.mk
include $(WS_MAKE_RULES)/ips.mk

CFLAGS += "-lnsl"

CONFIGURE_SCRIPT = $(@D)/configure

CONFIGURE_OPTIONS += --sysconfdir=/etc
CONFIGURE_OPTIONS += --libexecdir=$(CONFIGURE_PREFIX)/lib
CONFIGURE_OPTIONS += --sharedstatedir=/var/cluster
CONFIGURE_OPTIONS += --localstatedir=/var
CONFIGURE_OPTIONS += --enable-fatal-warnings=no
CONFIGURE_OPTIONS += --enable-ansi
CONFIGURE_OPTIONS += --with-ocf-root=$(CONFIGURE_PREFIX)/lib/ocf

COMPONENT_PREP_ACTION=(cd $(SOURCE_DIR); ./autogen.sh) 

COMPONENT_PRE_CONFIGURE_ACTION=($(CLONEY) $(SOURCE_DIR) $(@D))

build:          $(BUILD_32)

install:        $(INSTALL_32)

test:           $(TESTS_32)
