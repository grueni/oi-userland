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
# Copyright 2017 Alexander Pyhalov
# Copyright 2018 Michal Nowak
# Copyright 2018 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved.
#

include ../../../make-rules/shared-macros.mk

COMPONENT_NAME=	mongodb
COMPONENT_MAJOR_VERSION= 4.0
COMPONENT_MINOR_VERSION= 3
COMPONENT_VERSION= $(COMPONENT_MAJOR_VERSION).$(COMPONENT_MINOR_VERSION)
COMPONENT_SUMMARY= MongoDB distributed document-oriented database
COMPONENT_PROJECT_URL=	https://www.mongodb.com/
COMPONENT_FMRI=	service/database/mongodb-40/cluster
COMPONENT_CLASSIFICATION= System/Databases
COMPONENT_SRC=		$(COMPONENT_NAME)-$(COMPONENT_VERSION)
#COMPONENT_SRC=		$(COMPONENT_NAME)-src-r$(COMPONENT_VERSION)
#COMPONENT_ARCHIVE=	$(COMPONENT_SRC).tar.gz
#COMPONENT_ARCHIVE_URL=	https://fastdl.mongodb.org/src/$(COMPONENT_ARCHIVE)
#COMPONENT_ARCHIVE_HASH=	sha256:fbbe840e62376fe850775e98eb10fdf40594a023ecf308abec6dcec44d2bce0c
COMPONENT_LICENSE=	Apache v2.0
COMPONENT_LICENSE_FILE=	LICENSE.txt

include $(WS_MAKE_RULES)/ips.mk

download:

build:
	[ -d $(SOURCE_DIR) ] || $(RM) -r $(SOURCE_DIR)
	$(MKDIR) $(SOURCE_DIR)
	$(CP) $(COMPONENT_DIR)/files/*.txt  $(SOURCE_DIR)/

install:
	[ -d $(PROTO_DIR) ] || $(RM) -r $(PROTO_DIR)
	$(MKDIR) $(PROTO_DIR)/etc/mongodb/4.0/
	$(MKDIR) $(PROTO_DIR)/lib/svc/manifest/application/database
	$(CP) $(COMPONENT_DIR)/files/*.conf $(PROTO_DIR)/etc/mongodb/4.0/
	$(CP) $(COMPONENT_DIR)/files/*.xml  $(PROTO_DIR)/lib/svc/manifest/application/database/

clean:
	$(RM) -r $(BUILD_DIR) $(PROTO_DIR) $(SOURCE_DIR)


