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

COMPONENT_NAME= pcs
COMPONENT_VERSION= 0.9.154
COMPONENT_FMRI= application/cluster/$(COMPONENT_NAME)
COMPONENT_CLASSIFICATION= System/Shells
COMPONENT_PROJECT_URL= https://github.com/ClusterLabs/pcs
COMPONENT_SRC= $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE= $(COMPONENT_VERSION).tar.gz
COMPONENT_ARCHIVE_HASH= sha256:f4720b4d64d8647dcd0e63fd883c5c1ad9bc49a3777f5092f8b3e01f193838a4
COMPONENT_ARCHIVE_URL= https://github.com/ClusterLabs/$(COMPONENT_NAME)/archive/$(COMPONENT_ARCHIVE) 
COMPONENT_LICENSE= GPLv2
COMPONENT_SUMMARY= Pacemaker command line interface and GUI.

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

PATH= /usr/gnu/bin:/usr/ruby/1.9/bin/:/usr/bin:/usr/sbin:/sbin

PATCH_LEVEL = 1

PYTHON=python2.7
PYTHON_SITELIB=PYTHON_SITELIB=${CONFIGURE_PREFIX}/lib/python2.7/site-packages

# if not defined DESTDIR is blank
COMPONENT_BUILD_ENV += DESTDIR="$(PROTO_DIR)"

COMPONENT_INSTALL_ENV += DESTDIR="$(PROTO_DIR)"

PKG_PROTO_DIRS += $(COMPONENT_DIR)/files

# copied from components/ruby/puppet/Makefile
# pass Ruby values onto the package manifest
PKG_MACROS +=	RUBY_VERSION=$(RUBY_VERSION)
PKG_MACROS += RUBY_LIB_VERSION=$(RUBY_LIB_VERSION)

# pytthon installation class overwrites the shebang during installation
$(INSTALL_32): 	$(BUILD_32)
								grep -q python2.7 ${PROTO_DIR}/usr/sbin/pcs || \
									gsed -i -- "s~/usr/bin/python~/usr/bin/python2.7~g" ${PROTO_DIR}/usr/sbin/pcs

# beware of: the package has a combined build + install 
build:          ${BUILD_32}

install: 				${INSTALL_32}

test:           $(TESTS_32)


REQUIRED_PACKAGES += SUNWcs
# just for installation, not for runtime
# REQUIRED_PACKAGES += library/python/setuptools-27
REQUIRED_PACKAGES += runtime/python-27
REQUIRED_PACKAGES += runtime/ruby-19
REQUIRED_PACKAGES += system/library
