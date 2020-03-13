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
# Copyright 2020 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved.
#

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME=			elixir
COMPONENT_MAJOR_VERSION=	1.10
COMPONENT_VERSION=		$(COMPONENT_MAJOR_VERSION).2
COMPONENT_FMRI=			developer/elixir
COMPONENT_CLASSIFICATION=	Development/Other Languages
COMPONENT_SUMMARY=		The dynamic and functional Elixir Programming Language
COMPONENT_SRC=			$(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_PROJECT_URL=		https://elixir-lang.org/
COMPONENT_ARCHIVE=		v$(COMPONENT_VERSION).tar.gz
COMPONENT_ARCHIVE_URL=		https://github.com/elixir-lang/elixir/archive/$(COMPONENT_ARCHIVE)
COMPONENT_ARCHIVE_HASH=		sha256:5adffcf4389aa82fcfbc84324ebbfa095fc657a0e15b2b055fc05184f96b2d50
COMPONENT_LICENSE=		Apache-2.0
COMPONENT_LICENSE_FILE=	LICENSE

COMPONENT_COPY_ACTION = \
	($(CP) -rp $(SOURCE_DIR)/. $(@D))

COMPONENT_INSTALL_ARGS = PREFIX=$(USRDIR)
COMPONENT_TEST_TARGETS = test

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

# common targets
build:		$(BUILD_64)

install:	build $(INSTALL_64)

test:		$(TEST_64)

# not found by REQUIRED_PACKAGES
REQUIRED_PACKAGES += runtime/erlang

# Auto-generated dependencies
REQUIRED_PACKAGES += SUNWcs

