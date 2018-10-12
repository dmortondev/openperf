ICP_DPDK_TARGET ?= default-linuxicp-clang

PIO_DEPENDS += dpdk

PIO_DRIVER_SOURCES += \
	arg_parser.cpp \
	arg_parser_register.c \
	driver.cpp \
	eal.cpp \
	model/physical_port.cpp \
	model/port_info.cpp

PIO_LIBS += -lcap
