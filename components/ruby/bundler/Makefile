
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
# Copyright 2016 Grüninger GmbH, Andreas Grüninger (grueni)
#
include ../../../make-rules/shared-macros.mk

COMPONENT_NAME=           bundler
COMPONENT_VERSION=        1.13.6
COMPONENT_FMRI=           library/ruby/$(COMPONENT_NAME)
COMPONENT_SRC=            $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE=        $(COMPONENT_SRC).gem
COMPONENT_ARCHIVE_HASH=   sha256:fafd22dfed658ca0603f321bdd168ed709d7c682e61273b55637716459f2d0f7
COMPONENT_ARCHIVE_URL=    http://rubygems.org/gems/$(COMPONENT_SRC).gem
COMPONENT_PROJECT_URL=    http://puppetlabs.com/bundler
COMPONENT_CLASSIFICATION= Development/Ruby
COMPONENT_LICENSE=        MIT

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/ips.mk
include $(WS_MAKE_RULES)/gem.mk

# symlink support is supported by ruby >= 2.3
# to generate bundler.p5m do the following:
# - add macro $(BUILD_DIR)/%/.built from gem.mk and replace ${CLONEY} .. with
# $(RM) -r $(BUILD_DIR_32) ; $(CP) -r $(SOURCE_DIR) $(@D)
# - create afterwards manifest with gmake sample-manifest
$(BUILD_DIR)/%/.built:  $(SOURCE_DIR)/.prep
	$(RM) -r $(@D) ; $(MKDIR) $(@D)
	$(RM) -r $(BUILD_DIR_32) ; $(CP) -r $(SOURCE_DIR) $(@D)
	$(COMPONENT_PRE_BUILD_ACTION)
	# Build the gem and cause the generation of a new gem spec
	# file in $(COMPONENT_SRC)
	$(GEM_BUILD_ACTION)
	$(COMPONENT_POST_BUILD_ACTION)
	$(TOUCH) $@

# common targets
build:		$(BUILD_32)

install:	$(INSTALL_32)

test:		  $(TESTS_32)


REQUIRED_PACKAGES += runtime/ruby-19
