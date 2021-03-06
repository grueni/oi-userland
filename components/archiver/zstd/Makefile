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
# Copyright 2019 Michal Nowak
#

PREFERRED_BITS=		64

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME=		zstd
COMPONENT_VERSION=	1.4.0
COMPONENT_REVISION=	1
COMPONENT_SUMMARY=	Zstandard, or zstd as short version, is a fast lossless compression algorithm, targeting real-time compression scenarios
COMPONENT_PROJECT_URL=	https://facebook.github.io/zstd/
COMPONENT_FMRI=		compress/zstd
COMPONENT_CLASSIFICATION=	Applications/System Utilities
COMPONENT_SRC=		$(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE=	$(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_URL=	https://github.com/facebook/zstd/releases/download/v$(COMPONENT_VERSION)/$(COMPONENT_ARCHIVE)
COMPONENT_ARCHIVE_HASH=	sha256:7f323f0e0c18488748f3d9b2d4353f00e904ea2ccd0372ea04d7f77aa1156557
COMPONENT_LICENSE=	BSD,GPLv2.0
COMPONENT_LICENSE_FILE=	$(COMPONENT_NAME).license

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/cmake.mk
include $(WS_MAKE_RULES)/ips.mk

COMPONENT_PRE_CMAKE_ACTION= ($(CLONEY) $(SOURCE_DIR) $(@D))

# CMake files are in build/cmake directory,
# hence mangle with 'configure' target
$(BUILD_DIR)/%/.configured:     $(SOURCE_DIR)/.prep
	($(RM) -rf $(@D) ; $(MKDIR) $(@D))
	$(COMPONENT_PRE_CMAKE_ACTION)
	(cd $(@D) ; $(ENV) $(CMAKE_ENV) \
		$(CMAKE) $(CMAKE_OPTIONS) $(COMPONENT_DIR)/$(COMPONENT_SRC)/build/cmake/)
	$(COMPONENT_POST_CMAKE_ACTION)
	$(TOUCH) $@

build:		$(BUILD_32_and_64)

install:	$(INSTALL_32_and_64)

# Use make targets for tests directly
COMPONENT_TEST_CMD=	gmake
COMPONENT_TEST_ARGS=	-C
COMPONENT_TEST_TARGETS=	tests test-zstd
COMPONENT_TEST_ENV=	$(CMAKE_ENV)

test:		$(TEST_32_and_64)

# Auto-generated dependencies
REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += system/library
