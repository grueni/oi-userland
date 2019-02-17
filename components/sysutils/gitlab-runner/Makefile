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

# Copyright 2019 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved.
#

include ../../../make-rules/shared-macros.mk

PATH=$(PATH.gnu)

COMPONENT_NAME=			gitlab-runner
COMPONENT_MAJOR_VERSION=	11.7
COMPONENT_VERSION=		$(COMPONENT_MAJOR_VERSION).0
COMPONENT_FMRI=			sysutils/gitlab-runner
COMPONENT_CLASSIFICATION=	Development/Other Languages
COMPONENT_SUMMARY=		Gitlab runner
COMPONENT_SRC=			$(COMPONENT_NAME)-v$(COMPONENT_VERSION)
COMPONENT_PROJECT_URL=		https://gitlab.com/
COMPONENT_ARCHIVE=		$(COMPONENT_SRC).tar.gz
COMPONENT_ARCHIVE_URL=		https://gitlab.com/gitlab-org/$(COMPONENT_NAME)/-/archive/v$(COMPONENT_VERSION)/$(COMPONENT_ARCHIVE)
COMPONENT_ARCHIVE_HASH=		sha256:ba5726feef52b302b3bf6f38a1a9eeb9d957cf54ec7b45a9774518ecc5995523
COMPONENT_LICENSE=	MIT, CC BY-SA 4.0	
COMPONENT_LICENSE_FILE=		LICENSE

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

ENV += GOOS=solaris
ENV += GOARCH=amd64
GOPATH=$(COMPONENT_DIR)

# if github.com/golang/crypto/ssh/terminal is updated to master, part after cp of Makefile can be removed
COMPONENT_PREP_ACTION += ( \
  $(CP) $(COMPONENT_DIR)/files/Makefile.openindiana $(@D); \
	$(CP) $(COMPONENT_DIR)/files/service_solaris.go $(@D)/vendor/github.com/ayufan/golang-kardianos-service; \
	$(CP) $(COMPONENT_DIR)/files/system_logger_solaris.go $(@D)/log/; \
	export GOPATH=$(GOPATH); \
	go get -u github.com/kardianos/govendor; \
	$(MKDIR) $(GOPATH)/src; \
	cd $(GOPATH)/src; \
	$(GOPATH)/bin/govendor init; \
	$(GOPATH)/bin/govendor fetch github.com/golang/crypto/ssh/terminal; \
	$(RM) $(SOURCE_DIR)/vendor/golang.org/x/crypto/ssh/terminal/terminal_test.go; \
	$(CP) $(GOPATH)/src/vendor/github.com/golang/crypto/ssh/terminal/* $(SOURCE_DIR)/vendor/golang.org/x/crypto/ssh/terminal/; \
	cd $(GOPATH); \
	$(RM) -r $(GOPATH)/bin $(GOPATH)/.cache $(GOPATH)/src; )

# The project doesn't have an install section supporting DESTDIR but let the user copy the folder packaging/root.
# Here we prepare manually the image in the prototype folder with a makefile.
COMPONENT_POST_INSTALL_ACTION += ( \
   cd $(@D); \
   $(RM) -r $(PROTO_DIR); \
   $(GMAKE) -f Makefile.openindiana \
      DESTDIR=$(PROTO_DIR) \
			USRBINDIR=$(USRBINDIR) \
			USRSHAREDIR=$(USRSHAREDIR) \
			BINDIR=$(BUILD_DIR_$(BITS))/.gopath/bin \
      ROOTDIR=$(BUILD_DIR_$(BITS))/packaging/root )

# common targets
build:		$(BUILD_64)

install:	$(INSTALL_64)

test:		$(TEST_64)

REQUIRED_PACKAGES += developer/golang-111
# Auto-generated dependencies
REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += system/library
