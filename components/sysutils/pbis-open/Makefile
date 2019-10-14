#
# CDDL HEADER START
#
# The contents of this file are subject to the terms of the
# Common Development and Distribution License (the "License").
# You may not use this file except in compliance with the License.
#
# You can obtain a copy of the license at usr/src/OPENSOLARIS.LICENSE
# or http://www.opensolaris.org/os/licensing.
# See the License for the specific language governing permissions
# and limitations under the License.
#
# When distributing Covered Code, include this CDDL HEADER in each
# file and include the License file at usr/src/OPENSOLARIS.LICENSE.
# If applicable, add the following below this CDDL HEADER, with the
# fields enclosed by brackets "[]" replaced with your own identifying
# information: Portions Copyright [yyyy] [name of copyright owner]
#
# CDDL HEADER END
#
#
# Copyright 2019 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved.
#


include ../../../make-rules/shared-macros.mk

COMPONENT_NAME= pbis-open
COMPONENT_SUMMARY= PowerBroker Open provides integration in Windows Active Directory
COMPONENT_VERSION= 9.1.0
COMPONENT_FMRI= system/management/$(COMPONENT_NAME)
COMPONENT_CLASSIFICATION=	System/Administration and Configuration
COMPONENT_SRC= $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_PROJECT_URL= https://github.com/BeyondTrust/pbis-open
COMPONENT_ARCHIVE= $(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_HASH= sha256:7539d26ce8a8cc9692e85c85184f892457bc45cae9580f5109e2ee2fdd72fb43
COMPONENT_ARCHIVE_URL= https://github.com/BeyondTrust/$(COMPONENT_NAME)/archive/$(COMPONENT_VERSION).tar.gz
COMPONENT_LICENSE= GPLv2,LGPLv2.1

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/configure.mk
include $(WS_MAKE_RULES)/ips.mk

LDFLAGS =
gcc_OPT =
CC_BITS =
CORES=$(shell psrinfo -v|grep '^Status of'|tail -1|awk '{print $$5}')

CC += -static-libgcc
CFLAGS += -I$(USRDIR)/apr/include
CFLAGS += -I$(USRDIR)/apr-util/include
CXXFLAGS = $(CFLAGS)

ENV += CFLAGS="-I/usr/apr/include -I/usr/apr-util/include"
ENV += CXXFLAGS="-I/usr/apr/include -I/usr/apr-util/include"
ENV += MAKE=$(GMAKE)

ENV += ACCEPT_INFERIOR_RM_PROGRAM=yes
ENV += BUILD_X86_32_LDFLAGS="-L/usr/apr/lib -L/usr/apr-util/lib"
ENV += BUILD_X86_64_LDFLAGS="-L/usr/apr/lib/amd64 -L/usr/apr-util/lib/amd64"
ENV += HOST_X86_32_LDFLAGS="-L/usr/apr/lib -L/usr/apr-util/lib"
ENV += HOST_X86_64_LDFLAGS="-L/usr/apr/lib/amd64 -L/usr/apr-util/lib/amd64"

#CONFIGURE_ENV += "ac_cv_func_realloc_0_nonnull=no"
#COMPONENT_BUILD_ENV += "ac_cv_func_realloc_0_nonnull=no"

CONFIGURE_PREFIX = /opt
CONFIGURE_OPTIONS += --debug=yes
CONFIGURE_OPTIONS += --apache2.4-prefix=$(USRDIR)/apache2/2.4
CONFIGURE_OPTIONS.64 += --build-isas="x86_64"

COMPONENT_BUILD_ARGS += default
COMPONENT_BUILD_ARGS += -j$(CORES)

build:    $(BUILD_32_and_64)

test:     $(NO_TESTS)

install:  $(INSTALL_32_and_64)

REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += system/library
REQUIRED_PACKAGES += $(GXX_RUNTIME_PKG)
REQUIRED_PACKAGES += $(GCC_RUNTIME_PKG)
REQUIRED_PACKAGES += web/server/apache-24
# Auto-generated dependencies
