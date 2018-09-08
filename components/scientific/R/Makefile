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
# Copyright 2017 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved.
#

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME=         R
COMPONENT_VERSION=      3.5.1
COMPONENT_SUMMARY=      PETSc - Portable, Extensible Toolkit for Scientific Computation
COMPONENT_PROJECT_URL=  https://www.r-project.org/
COMPONENT_FMRI=         runtime/R
COMPONENT_CLASSIFICATION=Applications/Office
COMPONENT_SRC=          $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE=      $(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_URL=  https://cran.r-project.org/src/base/R-3/$(COMPONENT_ARCHIVE) 
COMPONENT_ARCHIVE_HASH= sha256:0463bff5eea0f3d93fa071f79c18d0993878fd4f2e18ae6cf22c1639d11457ed
COMPONENT_LICENSE=      GPLv2, GPLv3, LGPLv2.1

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/configure.mk
include $(WS_MAKE_RULES)/ips.mk

PATH=$(PATH.gnu)

# iconv.h from gnu iconv is required
CFLAGS += -I/usr/gnu/include
LDFLAGS += -R$(USRLIBDIR64)/R/lib

CONFIGURE_OPTIONS+= --with-x
CONFIGURE_OPTIONS+= --with-blas=no
CONFIGURE_OPTIONS+= --with-lapack=no
CONFIGURE_OPTIONS+= --enable-R-shlib
CONFIGURE_OPTIONS+= --enable-BLAS-shlib
CONFIGURE_OPTIONS+= --enable-memory-profiling
#CONFIGURE_OPTIONS+= --enable-threads=posix

build:		$(BUILD_64)

install:	$(INSTALL_64)

test:		$(TEST_64)

# Auto-generated dependencies
REQUIRED_PACKAGES += $(GFORTRAN_RUNTIME_PKG)
REQUIRED_PACKAGES += compress/bzip2
REQUIRED_PACKAGES += compress/xz
REQUIRED_PACKAGES += image/library/libjpeg8-turbo
REQUIRED_PACKAGES += image/library/libpng16
REQUIRED_PACKAGES += image/library/libtiff
REQUIRED_PACKAGES += library/desktop/cairo
REQUIRED_PACKAGES += library/desktop/pango
REQUIRED_PACKAGES += library/glib2
REQUIRED_PACKAGES += library/gnu-libiconv
REQUIRED_PACKAGES += library/icu
REQUIRED_PACKAGES += library/pcre
REQUIRED_PACKAGES += library/readline
REQUIRED_PACKAGES += library/zlib
REQUIRED_PACKAGES += runtime/tcl-8
REQUIRED_PACKAGES += runtime/tk-8
REQUIRED_PACKAGES += shell/bash
REQUIRED_PACKAGES += system/library
REQUIRED_PACKAGES += $(GCC_RUNTIME_PKG)
REQUIRED_PACKAGES += system/library/math
REQUIRED_PACKAGES += web/curl
REQUIRED_PACKAGES += x11/library/libx11
REQUIRED_PACKAGES += x11/library/libxmu
REQUIRED_PACKAGES += x11/library/toolkit/libxt
