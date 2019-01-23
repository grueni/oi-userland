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
# Copyright 2019 Andreas Grueninger, Grueninger GmbH, (grueni). All rights reserved.
#

include ../../../make-rules/shared-macros.mk

PATH= $(PATH.gnu)

COMPONENT_NAME= glusterfs
COMPONENT_VERSION= 5.3
COMPONENT_FMRI= application/cluster/$(COMPONENT_NAME)
COMPONENT_CLASSIFICATION= System/Services
COMPONENT_SRC= $(COMPONENT_NAME)-$(COMPONENT_VERSION)
COMPONENT_ARCHIVE= $(COMPONENT_NAME)-$(COMPONENT_VERSION).tar.gz
COMPONENT_ARCHIVE_HASH= sha256:ae11a89fc2b7980b4913e2b7c196da2961d78ca776d30ddbe9e98c4b8b95021d
COMPONENT_PROJECT_URL= https://www.gluster.org/
COMPONENT_ARCHIVE_URL= https://github.com/gluster/$(COMPONENT_NAME)/archive/v$(COMPONENT_VERSION).tar.gz
COMPONENT_LICENSE= GPLV2, LGPLV3+
COMPONENT_SUMMARY= Gluster Filesystem.

COMPONENT_NAME_1= argp-standalone
COMPONENT_VERSION_1= 1.3
COMPONENT_SRC_1= $(COMPONENT_NAME_1)-$(COMPONENT_VERSION_1)
COMPONENT_ARCHIVE_1= $(COMPONENT_SRC_1).tar.gz
COMPONENT_ARCHIVE_HASH_1= sha256:dec79694da1319acd2238ce95df57f3680fea2482096e483323fddf3d818d8be
COMPONENT_PROJECT_URL_1= https://www.lysator.liu.se/~nisse/
COMPONENT_ARCHIVE_URL_1= https://www.lysator.liu.se/~nisse/misc/$(COMPONENT_ARCHIVE_1)
COMPONENT_LICENSE_1= GPL

COMPONENT_NAME_2= cmocka
COMPONENT_VERSION_2= 1.1.3
COMPONENT_SRC_2= $(COMPONENT_NAME_2)-$(COMPONENT_VERSION_2)
COMPONENT_ARCHIVE_2= $(COMPONENT_SRC_2).tar.xz
COMPONENT_ARCHIVE_HASH_2= sha256:43eabcf72a9c80e3d03f7c8a1c04e408c18d2db5121eb058a3ef732a9dfabfaf
COMPONENT_PROJECT_URL_2= https://cmocka.org
COMPONENT_ARCHIVE_URL_2= https://cmocka.org/files/1.1/$(COMPONENT_ARCHIVE_2)
COMPONENT_LICENSE_2= GPL


CCFLAGS=$(gcc_PIC) -std=gnu99 -fgnu89-inline -m64 -O2

COMPONENT_POST_UNPACK_ACTION_1 = ( \
  cd $(COMPONENT_SRC); $(LN) -s ../$(COMPONENT_SRC_1) $(COMPONENT_SRC_1); \
  cd $(COMPONENT_SRC_1); \
	patch -p1 -i ../files/1.patch; \
	patch -p1 -i ../files/2.patch; \
	patch -p1 -i ../files/3.patch; \
	patch -p1 -i ../files/4.patch; \
	patch -p1 -i ../files/argp-standalone.patch; \
	CFLAGS="$(CCFLAGS)" ./configure && make )

include $(WS_MAKE_RULES)/prep.mk
include $(WS_MAKE_RULES)/configure.mk
include $(WS_MAKE_RULES)/ips.mk

CORES=$(shell psrinfo -v|grep '^Status of'|tail -1|awk '{print $$5}')

COMPONENT_PREP_ACTION=(cd $(SOURCE_DIR); $(CONFIG_SHELL) autogen.sh )

COMPONENT_PRE_CONFIGURE_ACTION=($(CLONEY) $(SOURCE_DIR) $(@D)) 

##COMPONENT_POST_CONFIGURE_ACTION=(cd $(@D) && make gitclean)

CONFIGURE_SCRIPT= $(@D)/configure

#CFLAGS += $(CPP_XPG7MODE)
#CFLAGS += $(gcc_PIC)
CFLAGS += $(CCFLAGS)  
CFLAGS += -Wall
CFLAGS += -Dasm=__asm__
CFLAGS += -Dtypeof=__typeof__
CFLAGS += -D_GNU_SOURCE
CFLAGS += -DPATH_MAX=1024
CFLAGS += -O2 -g -pipe -D_FORTIFY_SOURCE=2 -D_GLIBCXX_ASSERTIONS -grecord-gcc-switches -mtune=generic -g -rdynamic
CFLAGS += -I$(SOURCE_DIR)/$(COMPONENT_SRC_1)

LDFLAGS +=  -lsocket -lnsl -lresolv -lrpcsvc -lrt -lcrypt -lpthread
LDFLAGS += -L$(SOURCE_DIR)/$(COMPONENT_SRC_1)

CONFIGURE_OPTIONS += --disable-silent-rules
CONFIGURE_OPTIONS += --enable-dependency-tracking
CONFIGURE_OPTIONS += --without-libtirpc
#CONFIGURE_OPTIONS += --enable-cmocka
 
CONFIGURE_ENV += UUID_CFLAGS=-I/usr/include/uuid
CONFIGURE_ENV += UUID_LIBS=-luuid

COMPONENT_BUILD_ARGS += -j$(CORES)

build:          $(BUILD_64)

install:        $(INSTALL_64)

test:           $(TESTS_64)

