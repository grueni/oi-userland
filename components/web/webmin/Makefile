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
# Copyright 2016 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved. 
#

# transform *.pl +x
# transform *.cgi +x

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME= webmin
COMPONENT_VERSION= 1.820
COMPONENT_FMRI= web/$(COMPONENT_NAME)
COMPONENT_CLASSIFICATION= Applications/System Utilities
COMPONENT_PROJECT_URL= http://www.webmin.com/
COMPONENT_SRC= $(COMPONENT_NAME)-$(COMPONENT_VERSION)-solaris
COMPONENT_ARCHIVE= $(COMPONENT_NAME)-$(COMPONENT_VERSION).tar.gz
COMPONENT_ARCHIVE_HASH= sha256:130088ed78b3800d0aed2912ab37c12936347b32ece83e7a692caa8dfb1175cb
#COMPONENT_ARCHIVE_URL= https://github.com/$(COMPONENT_NAME)/$(COMPONENT_NAME)/archive/$(COMPONENT_VERSION).tar.gz
COMPONENT_ARCHIVE_URL= https://github.com/grueni/$(COMPONENT_NAME)/archive/v$(COMPONENT_VERSION)-solaris.tar.gz
#	https://github.com/grueni/webmin/archive/v1.820-solaris.tar.gz
COMPONENT_LICENSE= BSD-3-Clause 
COMPONENT_SUMMARY= Webmin is a web-based interface for system administration for Unix. 

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

ENV += login=admin
ENV += password=admin
ENV += tempdir=${PROTO_DIR}/var/tmp/webmin
ENV += config_dir=${PROTO_DIR}/etc/webmin
ENV += var_dir=${PROTO_DIR}/var/webmin
ENV += perl=/usr/bin/perl
ENV += port=10000
ENV += ssl=1
ENV += atboot=0
ENV += nostart=1
ENV += tar=${GTAR}

WEBMINDIR=/opt/webmin
GREP= /usr/gnu/bin/grep

WP2D += .built .unpacked .patched .prep setup_sh.patch

$(BUILD_DIR)/%/.built:  $(SOURCE_DIR)/.prep
	$(RM) -r $(@D) ; $(MKDIR) $(@D)
	${FIND} $(SOURCE_DIR) -type f -print0|xargs -0 file|${GREP} ELF|cut -d: -f1|xargs ${RM}
	${ENV} $(SOURCE_DIR)/makepackage.pl "$(SOURCE_DIR)" "$(@D)" 
ifeq   ($(strip $(PARFAIT_BUILD)),yes)
	-$(PARFAIT) build
endif
	$(TOUCH) $@

COMPONENT_INSTALL_ACTION = \
	cd $(@D) ; \
	${MKDIR} $(PROTO_DIR)/var/tmp/webmin; \
	${MKDIR} $(PROTO_DIR)/etc/webmin; \
	${MKDIR} $(PROTO_DIR)/var/webmin; \
	${MKDIR} $(PROTO_DIR)${WEBMINDIR}; \
	${ENV} ./setup.sh ${PROTO_DIR}${WEBMINDIR} ; \
	(cd $(PROTO_DIR)${WEBMINDIR} ; for WORD in `echo "${WP2D}"` ; do ${RM} -r $${WORD} ; done ;); \
	${GREP} -irl "${PROTO_DIR}" ${PROTO_DIR} | xargs $(GSED) -i -- "s~${PROTO_DIR}~~g"

# common targets
build:          $(BUILD_32)

install:        $(INSTALL_32)

test:           $(TEST_32)

REQUIRED_PACKAGES += shell/bash
REQUIRED_PACKAGES += shell/ksh93
REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += system/library
REQUIRED_PACKAGES += runtime/perl-512
REQUIRED_PACKAGES += library/perl-5/authen-pam-512
REQUIRED_PACKAGES += system/storage/smp_utils
REQUIRED_PACKAGES += system/storage/smartmontools
