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

COMPONENT_NAME= glusterfs
COMPONENT_VERSION= 5.3
COMPONENT_FMRI= application/cluster/$(COMPONENT_NAME)
COMPONENT_CLASSIFICATION= System/Services
COMPONENT_SRC= $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE= $(COMPONENT_NAME)-$(COMPONENT_VERSION).tar.gz
COMPONENT_ARCHIVE_HASH= sha256:ae11a89fc2b7980b4913e2b7c196da2961d78ca776d30ddbe9e98c4b8b95021d
COMPONENT_PROJECT_URL= https://www.gluster.org/
COMPONENT_ARCHIVE_URL= https://github.com/gluster/$(COMPONENT_NAME)/archive/v$(COMPONENT_VERSION).tar.gz
COMPONENT_LICENSE= GPLV2, LGPLV3+
COMPONENT_SUMMARY= Gluster Filesystem.

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/configure.mk
include $(WS_MAKE_RULES)/ips.mk

COMPONENT_PREP_ACTION=(cd $(SOURCE_DIR); $(CONFIG_SHELL) autogen.sh )

build:          $(BUILD_64)

install:        $(INSTALL_64)

test:           $(TESTS_64)

