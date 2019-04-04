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

PATH = $(PATH.gnu)

COMPONENT_NAME= webmin
COMPONENT_VERSION= 1.900
COMPONENT_FMRI= system/management/$(COMPONENT_NAME)
COMPONENT_CLASSIFICATION= Applications/System Utilities
COMPONENT_PROJECT_URL= http://www.webmin.com/
COMPONENT_SRC= $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE= $(COMPONENT_NAME)-$(COMPONENT_VERSION).tar.gz
COMPONENT_ARCHIVE_HASH= sha256:44269655a736d14e41f1a95663b1c7d6f908bd5248ae5f4c4caad0a26998fe1b
COMPONENT_ARCHIVE_URL= https://github.com/$(COMPONENT_NAME)/$(COMPONENT_NAME)/archive/$(COMPONENT_VERSION).tar.gz
#COMPONENT_ARCHIVE_URL= https://github.com/grueni/$(COMPONENT_NAME)/archive/v$(COMPONENT_VERSION)-solaris.tar.gz
#	https://github.com/grueni/webmin/archive/v1.820-solaris.tar.gz
COMPONENT_LICENSE= BSD-3-Clause 
COMPONENT_SUMMARY= Webmin is a web-based interface for system administration for Unix. 

COMPONENT_NAME_1=       authentic-theme
COMPONENT_VERSION_1=    19.31
COMPONENT_SRC_1=        $(COMPONENT_NAME_1)
COMPONENT_ARCHIVE_1=    $(COMPONENT_SRC_1)-$(COMPONENT_VERSION_1).wbt.gz
COMPONENT_ARCHIVE_URL_1=https://github.com/$(COMPONENT_NAME_1)/$(COMPONENT_NAME_1)/releases/download/$(COMPONENT_VERSION_1)/$(COMPONENT_ARCHIVE_1)
COMPONENT_ARCHIVE_HASH_1= sha256:0d57475b7cdb7c62984cd7b861b56cabd395415701bfa3a7c1ea03ac4f29cd52
COMPONENT_LICENSE_1= MIT

# http://github.com/authentic-theme/authentic-theme/releases/download/19.31/authentic-theme-19.31.wbt.gz
# cd /usr/webmin;tar xf 

# files/Makefile.PL: how to create files/Makefile

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/justmake.mk
include $(WS_MAKE_RULES)/ips.mk

# automatic installation webmin
ENV += login=admin
ENV += password=admin
ENV += tempdir=${PROTO_DIR}$(VARDIR)/tmp/webmin
ENV += config_dir=${PROTO_DIR}$(ETCDIR)/webmin
ENV += var_dir=${PROTO_DIR}$(VARDIR)/webmin
ENV += perl=$(PERL)
ENV += port=10000
ENV += ssl=1
ENV += atboot=0
ENV += nostart=1
ENV += tar=${GTAR}

WEBMINDIR=/usr/webmin
GREP= /usr/gnu/bin/grep

$(BUILD_DIR)/%/.built:  $(SOURCE_DIR)/.prep
	$(RM) -r $(@D) ; $(MKDIR) $(@D)
	${FIND} $(SOURCE_DIR) -type f -print0|xargs -0 file|${GREP} ELF|cut -d: -f1|xargs ${RM}
	${ENV} $(SOURCE_DIR)/makepackage.pl "$(SOURCE_DIR)" "$(@D)" 
ifeq   ($(strip $(PARFAIT_BUILD)),yes)
	-$(PARFAIT) build
endif
	$(CP)    $(COMPONENT_DIR)/files/Makefile* $(@D) 
	$(CP)    $(COMPONENT_DIR)/files/config $(@D) 
	$(TOUCH) $@

COMPONENT_INSTALL_ARGS += DESTDIR=$(PROTO_DIR)
COMPONENT_INSTALL_ARGS += WEBMINDIR=$(WEBMINDIR)
COMPONENT_INSTALL_ARGS += RELDIR=$(@D)
COMPONENT_INSTALL_ARGS += ETCDIR=$(ETCDIR)
COMPONENT_INSTALL_ARGS += VARDIR=$(VARDIR)
COMPONENT_INSTALL_ARGS += THEME=$(USERLAND_ARCHIVES)$(COMPONENT_ARCHIVE_1)

# make install repeatable  
COMPONENT_PRE_INSTALL_ACTION = \
	$(RM) -r $(PROTO_DIR); \
	${MKDIR} $(PROTO_DIR)  

COMPONENT_POST_INSTALL_ACTION = \
	${CP}    $(COMPONENT_DIR)/files/defaultacl.system-status $(PROTO_DIR)${WEBMINDIR}/system-status/defaultacl  

# common targets
build:          $(BUILD_32)

install:        $(INSTALL_32)

test:           $(TEST_32)

REQUIRED_PACKAGES += shell/bash
REQUIRED_PACKAGES += SUNWcs
REQUIRED_PACKAGES += system/library
REQUIRED_PACKAGES += runtime/perl-522
REQUIRED_PACKAGES += library/perl-5/authen-pam-522
