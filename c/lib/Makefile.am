noinst_LTLIBRARIES=	libmlr.la
libmlr_la_SOURCES=	free_flags.h \
			minunit.h \
			mlr_arch.c \
			mlr_arch.h \
			mlr_globals.c \
			mlr_globals.h \
			mlrdatetime.c \
			mlrdatetime.h \
			mlrescape.c \
			mlrescape.h \
			mlrmath.c \
			mlrmath.h \
			mlrstat.c \
			mlrstat.h \
			mlrregex.c \
			mlrregex.h \
			mlrtimezone.h \
			mlrutil.c \
			mlrutil.h \
			mlrval.c \
			mlrval.h \
			mvfuncs.c \
			mvfuncs.h \
			netbsd_strptime.c \
			netbsd_strptime.h \
			nlnet_timegm.c \
			nlnet_timegm.h \
			context.c \
			context.h \
			mtrand.c \
			mtrand.h \
			string_array.c \
			string_array.h \
			string_builder.c \
			string_builder.h \
			mlr_test_util.c \
			mlr_test_util.h \
			utf8.h \
			wcwidth9.h

AM_CPPFLAGS=	-I${srcdir}/../
AM_CFLAGS=	-std=gnu99

wcwidth9.h:
	curl -sL https://github.com/waltarix/localedata/releases/download/13.0.0-r1/wcwidth9.h > $@
